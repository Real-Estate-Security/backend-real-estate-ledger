package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/hash"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"backend_real_estate/internal/database"
	"backend_real_estate/internal/gateway"
	"backend_real_estate/internal/server"
	"backend_real_estate/util"
)

func gracefulShutdown(httpServer *http.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Info().Msg("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Info().Msg("Server shutdown complete.")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

func runDBMigrations(migrationUrl string, dbSource string) {
	m, err := migrate.New(migrationUrl, dbSource)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create migration")
	}

	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Info().Msg("No migration needed")
		} else {
			log.Fatal().Err(err).Msg("Failed to run migration")
		}
	}

	log.Info().Msg("Migration complete")
}

func main() {

	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	if config.AppEnv == "local" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// connect to gateway

	clientConnection := gateway.NewGrpcConnection()
	defer clientConnection.Close()

	id := gateway.NewIdentity()
	sign := gateway.NewSign()

	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithHash(hash.SHA256),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	defer gw.Close()

	// connect to the database

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", config.DbUsername, config.DbPassword, config.DbHost, config.DbPort, config.DbDatabase, config.DbSchema)

	dbConn, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	runDBMigrations(config.MigrationUrl, connStr)

	httpServer, err := server.NewHTTPServer(config, database.NewService(dbConn), gw)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create server")
	}

	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(httpServer, done)

	err = httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	// Wait for the graceful shutdown to complete
	<-done
	log.Info().Msg("Graceful shutdown complete.")
}
