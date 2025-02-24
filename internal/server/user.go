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
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Dob      time.Time `json:"dob" binding:"required" validate:"datetime=2006-01-02"`
	Role     string    `json:"role" binding:"required" validate:"oneof=user agent"`
}
type userResponse struct {
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Dob      time.Time `json:"dob"`
	Role     string    `json:"role"`
}

func createUserResponse(user database.Users) userResponse {
	return userResponse{
		Username: user.Username,
		Email:    user.Email,
		Dob:      user.Dob,
		Role:     string(user.Role),
	}
}

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
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

func (s *Server) LoginUserHandler(c *gin.Context) {
	var req loginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := s.dbService.GetUserByUsername(c, req.Username)
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
