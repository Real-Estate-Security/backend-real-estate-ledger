package server

import (
	"backend_real_estate/internal/database"
	"backend_real_estate/internal/token"
	"backend_real_estate/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username  string    `json:"username" binding:"required,alphanum"`
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	Password  string    `json:"password" binding:"required,min=6"`
	Email     string    `json:"email" binding:"required,email"`
	Dob       time.Time `json:"dob" binding:"required" validate:"datetime=2006-01-02"`
	Role      string    `json:"role" binding:"required" validate:"oneof=user agent"`
}
type userResponse struct {
	Username  string    `json:"username" binding:"required,alphanum"`
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Dob       time.Time `json:"dob" binding:"required" validate:"datetime=2006-01-02"`
	Role      string    `json:"role" binding:"required" validate:"oneof=user agent"`
}

func createUserResponse(user database.Users) userResponse {
	return userResponse{
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Dob:       user.Dob,
		Role:      string(user.Role),
	}
}

// CreateUserHandler handles the creation of a new user.
//
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags users
// @Accept json
// @Produce json
// @Param createUserRequest body createUserRequest true "Create User Request"
// @Success 200 {object} userResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /user/signup [post]
func (s *Server) CreateUserHandler(c *gin.Context) {
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := database.CreateUserParams{
		Username:       req.Username,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		HashedPassword: hashedPassword,
		Email:          req.Email,
		Dob:            req.Dob,
		Role:           database.UserRole(req.Role),
	}

	user, err := s.dbService.CreateUser(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := createUserResponse(user)

	c.JSON(http.StatusOK, resp)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

// LoginUserHandler handles the user login process.
//
// @Summary User login
// @Description Authenticates a user and returns an access token along with user details.
// @Tags users
// @Accept json
// @Produce json
// @Param loginUserRequest body loginUserRequest true "Login request"
// @Success 200 {object} loginUserResponse "Successful login"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal server error"
// @Router /user/login [post]
func (s *Server) LoginUserHandler(c *gin.Context) {
	var req loginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var user database.Users
	var err error

	// Check if the input is an email or username
	if util.IsEmail(req.Username) {
		user, err = s.dbService.GetUserByEmail(c, req.Username)
	} else {
		user, err = s.dbService.GetUserByUsername(c, req.Username)
	}

	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if err := util.CheckPassword(req.Password, user.HashedPassword); err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, _, err := s.tokenMaker.CreateToken(user.Username, string(user.Role), s.config.AccessTokenDuration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := loginUserResponse{
		AccessToken: accessToken,
		User:        createUserResponse(user),
	}

	c.JSON(http.StatusOK, resp)
}

type userMeResponse struct {
	User userResponse `json:"user"`
}

// UserMeHandler handles the request to retrieve the authenticated user's information.
//
// @Summary Get authenticated user information
// @Description Retrieves the information of the authenticated user based on the authorization token provided.
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} userMeResponse "Successfully retrieved user information"
// @Failure 401 {object} string "Unauthorized"
// @Security BearerAuth
// @Router /user/me [get]
func (s *Server) UserMeHandler(c *gin.Context) {
	// get authorization payload
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := s.dbService.GetUserByUsername(c, authPayload.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	resp := userMeResponse{
		User: createUserResponse(user),
	}

	c.JSON(http.StatusOK, resp)
}
