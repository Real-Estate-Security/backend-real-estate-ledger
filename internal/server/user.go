package server

import (
	"backend_real_estate/internal/database"
	"backend_real_estate/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)



type createUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Dob	  time.Time `json:"dob" binding:"required" validate:"datetime=2006-01-02"`
	Role	  string `json:"role" binding:"required" validate:"oneof=user agent"`
}
type createUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Dob	  time.Time `json:"dob"`
	Role	  string `json:"role"`
}

func newUserResponse(user database.Users) createUserResponse {
	return createUserResponse{
		Username: user.Username,
		Email:    user.Email,
		Dob:	  user.Dob,
		Role:	  string(user.Role),
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

	resp := newUserResponse(user)

	c.JSON(http.StatusOK, resp)
}

