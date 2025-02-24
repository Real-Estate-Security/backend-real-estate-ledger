package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"backend_real_estate/internal/database"
	"backend_real_estate/internal/token"
	"backend_real_estate/util"
)

type Server struct {
	config     util.Config
	tokenMaker token.Maker
	dbService  database.Service
	router     *gin.Engine
}

func NewHTTPServer(config util.Config, dbService database.Service) (*http.Server, error) {

	// Create a new server
	NewServer, err := NewGinServer(config, dbService)

	if err != nil {
		return nil, err
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", NewServer.config.Port),
		Handler:      NewServer.router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server, nil
}

func NewGinServer(config util.Config, dbService database.Service) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	ginServer := &Server{
		config:     config,
		tokenMaker: tokenMaker,
		dbService:  dbService,
	}

	ginServer.RegisterRoutes()

	return ginServer, nil
}
