package server

import (
	"backend_real_estate/internal/database"
	"backend_real_estate/internal/token"
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type requestRepresentationRequest struct {
	ClientUsername string `json:"client_username" binding:"required,alphanum"`
	StartDate      time.Time
	EndDate        time.Time
}

// RequestRepresentationHandler handles an agent's request to represent a user.
//
// @Summary Request representation
// @Description Allows an agent to request representation for a user.
// @Tags agent
// @Accept json
// @Produce json
// @Param requestRepresentationRequest body requestRepresentationRequest true "Request Representation Request"
// @Success 200 {object} string "Representation request submitted successfully"
// @Failure 400 {object} string "Invalid request"
// @Failure 404 {object} string "Client not found"
// @Failure 500 {object} string "Internal server error"
// @Router /agent/request-representation [post]
func (s *Server) RequestRepresentationHandler(c *gin.Context) {
	var req requestRepresentationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Check if the client exists
	client, err := s.dbService.GetUserByUsername(c, req.ClientUsername)
	if err != nil {
		c.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	// Get the agent's username from the authorization payload
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	agent, err := s.dbService.GetUserByUsername(c, authPayload.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	// Ensure the requester is an agent
	if agent.Role != database.UserRoleAgent {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Only agents can request representation"})
		return
	}

	// Submit the representation request to the database
	arg := database.CreateRepresentationParams{
		UserID:    client.ID,
		AgentID:   agent.ID,
		StartDate: time.Now(),
		IsActive:  false, // Pending status
	}

	_, err = s.dbService.CreateRepresentation(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Representation request submitted successfully"})
}

// AcceptRepresentationHandler handles accepting a representation request.
//
// @Summary Accept representation request
// @Description Allows an agent to accept a representation request.
// @Tags agent
// @Accept json
// @Produce json
// @Param id path int true "Representation ID"
// @Success 200 {object} string "Representation request accepted successfully"
// @Failure 400 {object} string "Invalid request"
// @Failure 404 {object} string "Representation not found"
// @Failure 500 {object} string "Internal server error"
// @Router /agent/accept-representation/{id} [post]
func (s *Server) AcceptRepresentationHandler(c *gin.Context) {
	id, err := parseIDParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := database.AcceptRepresentationParams{
		SignedDate: sql.NullTime{Time: time.Now(), Valid: true},
		ID:         id,
	}

	representation, err := s.dbService.AcceptRepresentation(c, arg)
	if err != nil {
		c.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Representation request accepted successfully", "representation": representation})
}

// DeclineRepresentationHandler handles declining a representation request.
//
// @Summary Decline representation request
// @Description Allows an agent to decline a representation request.
// @Tags agent
// @Accept json
// @Produce json
// @Param id path int true "Representation ID"
// @Success 200 {object} string "Representation request declined successfully"
// @Failure 400 {object} string "Invalid request"
// @Failure 404 {object} string "Representation not found"
// @Failure 500 {object} string "Internal server error"
// @Router /agent/decline-representation/{id} [post]
func (s *Server) DeclineRepresentationHandler(c *gin.Context) {
	id, err := parseIDParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	representation, err := s.dbService.RejectRepresentation(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Representation request declined successfully", "representation": representation})
}

// ListRepresentationsHandler handles fetching all representations for the authenticated user.
//
// @Summary List representations
// @Description Fetches all representations for the authenticated user, whether they are an agent or a regular user.
// @Tags representations
// @Accept json
// @Produce json
// @Success 200 {array} database.Representations "List of representations"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal server error"
// @Router /agent/representations [get]
func (s *Server) ListRepresentationsHandler(c *gin.Context) {
	// Get the authenticated user's information
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := s.dbService.GetUserByUsername(c, authPayload.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	var representations []database.Representations

	// Fetch representations based on the user's role
	if user.Role == database.UserRoleAgent {
		representations, err = s.dbService.ListRepresentationsByAgentID(c, database.ListRepresentationsByAgentIDParams{
			AgentID: user.ID,
			Limit:   100, // Default limit
			Offset:  0,   // Default offset
		})
	} else {
		representations, err = s.dbService.ListRepresentationsByUserID(c, database.ListRepresentationsByUserIDParams{
			UserID: user.ID,
			Limit:  100, // Default limit
			Offset: 0,   // Default offset
		})
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, representations)
}

// parseIDParam parses the ID parameter from the URL.
func parseIDParam(c *gin.Context, param string) (int64, error) {
	idParam := c.Param(param)
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}
