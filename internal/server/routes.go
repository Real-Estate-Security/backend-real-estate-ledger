package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) RegisterRoutes() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	// general health check routes
	router.GET("/hello-world", server.HelloWorldHandler)
	router.GET("/health", server.healthHandler)

	// user routes unprotected
	router.POST("/user/signup", server.CreateUserHandler)
	router.POST("/user/login", server.LoginUserHandler)

	// user routes protected
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/user/me", server.UserMeHandler)

	server.router = router
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.dbService.Health())
}
