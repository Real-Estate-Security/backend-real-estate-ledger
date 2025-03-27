package server

import (
	"backend_real_estate/internal/database"
	"backend_real_estate/internal/token"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type requestAgentRepresentationRequest struct {
	ClientUsername string    `json:"client_username" binding:"required,alphanum"`
	StartDate      time.Time `json:"start_date" binding:"required" validate:"datetime=2006-01-02"`
	EndDate        time.Time `json:"end_date" binding:"required" validate:"datetime=2006-01-02"`
}

// RequestRepresentationHandler handles an agent's request to represent a user.
//
// @Summary Request representation
// @Description Allows an agent to request representation for a user.
// @Tags agent
// @Accept json
// @Produce json
// @Param requestAgentRepresentationRequest body requestAgentRepresentationRequest true "Request Representation Request"
// @Success 200 {object} string "Representation request submitted successfully"
// @Failure 400 {object} string "Invalid request"
// @Failure 404 {object} string "Client not found"
// @Failure 500 {object} string "Internal server error"
// @Router /agent/request-representation [post]
// @Security BearerAuth
func (s *Server) RequestRepresentationHandler(c *gin.Context) {
	var req requestAgentRepresentationRequest
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

	// check if client is same as agent

	// Get the agent's username from the authorization payload
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)

	if client.Username == authPayload.Username {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Client cannot be the same as the agent"})
		return
	}

	agent, err := s.dbService.GetUserByUsername(c, authPayload.Username)

	// chekc if user is already represented by this agent

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
// @Security BearerAuth
func (s *Server) AcceptRepresentationHandler(c *gin.Context) {
	id, err := parseIDParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// check if the representation exists
	representation, err := s.dbService.GetRepresentationByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	// check if the representation is already accepted or declined
	if representation.Status != "pending" {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Representation request already %s", representation.Status)})
		return
	}

	arg := database.AcceptRepresentationParams{
		SignedDate: sql.NullTime{Time: time.Now(), Valid: true},
		ID:         id,
	}

	representation, err = s.dbService.AcceptRepresentation(c, arg)
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
// @Security BearerAuth
func (s *Server) DeclineRepresentationHandler(c *gin.Context) {
	id, err := parseIDParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// check if the representation exists
	representation, err := s.dbService.GetRepresentationByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	// check if the representation is already accepted
	if representation.Status == "accepted" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Representation request already accepted"})
		return
	}
	// check if the representation is already declined
	if representation.Status == "rejected" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Representation request already declined"})
		return
	}

	representation, err = s.dbService.RejectRepresentation(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Representation request declined successfully", "representation": representation})
}

type listRepresentationsRequest struct {
	Limit  int32 `json:"limit" binding:"min=1,max=100"`
	Offset int32 `json:"offset" binding:"min=0"`
}

// NullableTime is a custom type to represent sql.NullTime in Swagger documentation.
type NullableTime struct {
	Time  time.Time `json:"time,omitempty"`
	Valid bool      `json:"valid"`
}

// MarshalJSON customizes the JSON representation of NullableTime.
func (nt NullableTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nt.Time)
}

// RepresentationsWithNullableTime is a struct to replace database.Representations for Swagger.
type RepresentationsWithNullableTime struct {
	ID         int64        `json:"id"`
	UserID     int64        `json:"user_id"`
	AgentID    int64        `json:"agent_id"`
	StartDate  time.Time    `json:"start_date"`
	EndDate    NullableTime `json:"end_date"`
	Status     string       `json:"status"`
	SignedDate NullableTime `json:"signed_date"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	IsActive   bool         `json:"is_active"`
}

// ListRepresentationsHandler handles fetching all representations for the authenticated user.
//
// @Summary List representations
// @Description Fetches all representations for the authenticated user, whether they are an agent or a regular user.
// @Tags representations
// @Accept json
// @Produce json
// @Param limit query int false "Limit (default: 10)"
// @Param offset query int false "Offset (default: 0)"
// @Success 200 {array} RepresentationsWithNullableTime "List of representations"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal server error"
// @Router /agent/representations [get]
// @Security BearerAuth
func (s *Server) ListRepresentationsHandler(c *gin.Context) {
	// Set default values
	req := listRepresentationsRequest{
		Limit:  10, // Default limit
		Offset: 0,  // Default offset
	}

	// Bind query parameters, overriding defaults if provided
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

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
			Limit:   req.Limit,
			Offset:  req.Offset,
		})
	} else {
		representations, err = s.dbService.ListRepresentationsByUserID(c, database.ListRepresentationsByUserIDParams{
			UserID: user.ID,
			Limit:  req.Limit,
			Offset: req.Offset,
		})
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Convert database.Representations to RepresentationsWithNullableTime
	var response []RepresentationsWithNullableTime

	for _, r := range representations {
		response = append(response, RepresentationsWithNullableTime{
			ID:         r.ID,
			UserID:     r.UserID,
			AgentID:    r.AgentID,
			StartDate:  r.StartDate,
			EndDate:    NullableTime{Time: r.EndDate.Time, Valid: r.EndDate.Valid},
			Status:     string(r.Status),
			SignedDate: NullableTime{Time: r.SignedDate.Time, Valid: r.SignedDate.Valid},
			CreatedAt:  r.CreatedAt,
			UpdatedAt:  r.UpdatedAt,
			IsActive:   r.IsActive,
		})
	}

	c.JSON(http.StatusOK, response)
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
