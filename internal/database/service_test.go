package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
	"time"
	_ "github.com/lib/pq"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var testQueries *Queries
var testDB *sql.DB

// runMigration executes `make migrate-up` in the test environment
func runMigration() error {
	fmt.Println("Running database migrations...")
	// @migrate -path internal/database/migration \
	// -database "postgresql://${LOCAL_DB_USERNAME}:${LOCAL_DB_PASSWORD}@${LOCAL_DB_HOST}:${LOCAL_DB_PORT}/${LOCAL_DB_DATABASE}?sslmode=disable" \
	// -verbose up
	// print the variables 
	fmt.Println("username_test: ", username)
	fmt.Println("password_test: ", password)
	fmt.Println("host_test: ", host)
	fmt.Println("port_test: ", port)
	fmt.Println("database_test: ", database)
	// run the migration
	cmd := exec.Command("migrate", "-path", "./migration",
		"-database", fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			username, password, host, port, database),
		"-verbose", "up")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// mustStartPostgresContainer starts a new postgres container for integration testing
// and returns a function to terminate it.
// It also runs the migration tool to apply the migrations.
func mustStartPostgresContainer() (func(context.Context, ...testcontainers.TerminateOption) error, error) {
	var (
		dbName = "database"
		dbPwd  = "password"
		dbUser = "user"
	)

	dbContainer, err := postgres.Run(
		context.Background(),
		"postgres:latest",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPwd),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second)), // Increased timeout
	)
	if err != nil {
		return nil, err
	}

	database = dbName
	password = dbPwd
	username = dbUser

	dbHost, err := dbContainer.Host(context.Background())
	if err != nil {
		return dbContainer.Terminate, err
	}

	dbPort, err := dbContainer.MappedPort(context.Background(), "5432/tcp")
	if err != nil {
		return dbContainer.Terminate, err
	}

	host = dbHost
	port = dbPort.Port()

	// Set environment variables for migration tool
	os.Setenv("LOCAL_DB_HOST", host)
	os.Setenv("LOCAL_DB_PORT", port)
	os.Setenv("LOCAL_DB_DATABASE", database)
	os.Setenv("LOCAL_DB_USERNAME", username)
	os.Setenv("LOCAL_DB_PASSWORD", password)

	// Run migration after database is up
	if err := runMigration(); err != nil {
		return dbContainer.Terminate, err
	}

	dbDriver := "postgres"
	dbSource := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPwd, dbHost, dbPort.Port(), dbName)

	// Connect to the database
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		return dbContainer.Terminate, err
	}

	testQueries = New(testDB)

	return dbContainer.Terminate, nil
}

func TestMain(m *testing.M) {
	teardown, err := mustStartPostgresContainer()
	if err != nil {
		log.Fatalf("could not start postgres container: %v", err)
	}

	m.Run()

	if teardown != nil && teardown(context.Background()) != nil {
		log.Fatalf("could not teardown postgres container: %v", err)
	}
}

func TestNew(t *testing.T) {
	srv := NewService()
	if srv == nil {
		t.Fatal("New() returned nil")
	}
}

func TestHealth(t *testing.T) {
	srv := NewService()

	stats := srv.Health()

	if stats["status"] != "up" {
		t.Fatalf("expected status to be up, got %s", stats["status"])
	}

	if _, ok := stats["error"]; ok {
		t.Fatalf("expected error not to be present")
	}

	if stats["message"] != "It's healthy" {
		t.Fatalf("expected message to be 'It's healthy', got %s", stats["message"])
	}
}

func TestClose(t *testing.T) {
	srv := NewService()

	if srv.Close() != nil {
		t.Fatalf("expected Close() to return nil")
	}
}
