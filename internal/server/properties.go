package server

import(
	db "backend_real_estate/internal/database"
	"net/http"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type createListingRequest struct {
	OwnerFirstName string  `json:"OwnerFirstName" binding:"required,alphanum"`
	OwnerLastName  string  `json:"OwnerLastName" binding:"required,alphanum"`
	OwnerEmail     string  `json:"OwnerEmail" binding:"required,email"`
	AgentFirstName string  `json:"AgentFirstName"`
	AgentLastName  string  `json:"AgentLastName"`
	AgentEmail     string  `json:"AgentEmail" binding:"required,email"`
	Price          string `json:"Price" binding:"required"`
	Address        string  `json:"Address" binding:"required"`
	City           string  `json:"City" binding:"required"`
	State          string  `json:"State" binding:"required"`
	Zipcode        int32  `json:"Zipcode" binding:"required"`
	Bedrooms       int     `json:"Bedrooms" binding:"required"`
	Bathrooms      int     `json:"Bathrooms" binding:"required"`
	Description    string  `json:"Description"`
}

// CreatePropertyAndListingHandler handles the request for creating a listing for a property
//
// @Summary given listing, create property if doesn't exist and then create listing for that property
// @Description creating a listing for a property
// @Tags properties
// @Accept json
// @Produce json
// @Param createListingRequest body createListingRequest true "creating a listing for a property"
// @Success 200 {object} string 
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /properties [post]
func (s *Server) CreatePropertyAndListingHandler(c *gin.Context) {
	var req createListingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	listing, err := s.CreateOrUsePropertyAndThenCreateListing(
		c.Request.Context(),
		req,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, listing)
}

func (s *Server) CreateOrUsePropertyAndThenCreateListing(ctx context.Context, req createListingRequest) (db.Listings, error) {
	//Try getting existing property (see if property already exists)
//var emptyResponse db.Listings
	property, err := s.dbService.GetPropertyByAddress(ctx, req.Address)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return db.Listings{}, fmt.Errorf("check property: %w", err)
	}

	// Create property if it doesn't exist
	if err == sql.ErrNoRows {
		ownerID, er := s.dbService.GetUserIDByEmail(ctx, req.OwnerEmail)
		if er != nil {
			return db.Listings{}, fmt.Errorf("get owner ID: %w", er)
		}
		property, err = s.dbService.CreateProperty(ctx, db.CreatePropertyParams{
			Owner:   ownerID,
			Address:   req.Address,
			City:      req.City,
			State:     req.State,
			Zipcode:   req.Zipcode,
			Bedrooms:  int32(req.Bedrooms),
			Bathrooms: int32(req.Bathrooms),
		})
		if err != nil {
			return db.Listings{}, fmt.Errorf("create property: %w", err)
		}
	}

	agentID, err := s.dbService.GetUserIDByEmail(ctx, req.AgentEmail)
	if err != nil {
		return db.Listings{}, fmt.Errorf("get agent ID: %w", err)
	}
	// Create listing
	listing, err := s.dbService.CreateListing(ctx, db.CreateListingParams{
		PropertyID: property.ID,
		AgentID:  agentID,
		Price:      req.Price,
		Description: sql.NullString{
			String: req.Description,
			Valid:  req.Description != "",
		},
	})
	if err != nil {
		return listing, fmt.Errorf("create listing: %w", err)
	}

	return listing, nil
}

