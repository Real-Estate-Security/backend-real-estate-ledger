package server

import (
	"net/http"

	docs "backend_real_estate/docs/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// @title Secure Real Estate Ledger Backend API
// @version 1.0
// @description This is the backend API for the Secure Real Estate Ledger application.
// @termsOfService http://swagger.io/terms/

// @contact.name malik{lastname}5@gmail.com
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your Bearer token in the format: Bearer <token>

func (server *Server) RegisterRoutes() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	docs.SwaggerInfo.BasePath = "/"
	// set base path for swagger

	// general health check routes
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/hello-world", server.HelloWorldHandler)
	router.GET("/health", server.healthHandler)

	//ledger routes
	router.GET("/properties", server.GetAllProperties)
	router.POST("/properties-ledger", server.RegisterProperty)
	router.POST("/properties/list/:id", server.ListProperty)
	router.POST("/properties/bid/:id", server.PlaceBid)
	router.DELETE("/properties/bid/:propertyID/:bidID", server.RejectBid)

	// user routes unprotected
	router.POST("/user/signup", server.CreateUserHandler)
	router.POST("/user/login", server.LoginUserHandler)
	router.GET("/listing", server.GetListingDisplayHandler) //CHANGE MIGHT BE NEEDED

	router.POST("/property/getPropertyByID", server.getPropertyByIDHandler)
	router.POST("/listing/getListingByPropertyID", server.getListingByPropertyIDHandler)
	router.POST("/bidding/createBid", server.createBidHandler)
	router.POST("/bidding/listBids", server.listBidsHandler)
	router.PUT("/bidding/rejectBid", server.rejectBidHandler)
	router.PUT("/bidding/acceptBid", server.acceptBidHandler)
	router.POST("bidding/listBidsOnListing", server.listBidsOnListingHandler)
	// user routes protected
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/user/me", server.UserMeHandler)

	// properties/listings routes 
	router.POST("/properties", server.CreatePropertyAndListingHandler)
	// agent routes protected
	authRoutes.POST("/agent/request-representation", server.RequestRepresentationHandler)
	authRoutes.POST("/agent/accept-representation/:id", server.AcceptRepresentationHandler)
	authRoutes.POST("/agent/decline-representation/:id", server.DeclineRepresentationHandler)

	// agent and user routes protected
	authRoutes.GET("/agent/representation", server.ListRepresentationsHandler)

	server.router = router
}

// HelloWorld godoc
// @Summary HelloWorld example
// @Schemes
// @Description HelloWorld example
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /hello-world [get]
func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

// healthHandler godoc
// @Summary Health Check
// @Description Returns the health status of the server
// @Tags health
// @Produce json
// @Success 200 {object} string
// @Router /health [get]
func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.dbService.Health())
}
