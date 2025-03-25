package server

import (
	"backend_real_estate/internal/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type getListingByIDRequest struct {
	PropertyID int64  `json:"PropertyID" binding:"required"`
	Username   string `json:"Username" binding:"required,alphanum"`
}

type listingResponse struct {
	ID            int64     `json:"ID" binding:"required"`
	PropertyID    int64     `json:"PropertyID" binding:"required"`
	AgentID       int64     `json:"AgentID" binding:"required"`
	Price         float64   `json:"Price" binding:"required"`
	ListingStatus string    `json:"ListingStatus" binding:"required"`
	ListingDate   time.Time `json:"ListingDate" binding:"required"`
	Description   string    `json:"Description" binding:"required"`
	AcceptedBidID int64     `json:"AcceptedBidID" binding:"required"`
}

func getListingResponse(currentListing database.Listings) listingResponse {
	return listingResponse{
		ID:            currentListing.ID,
		PropertyID:    currentListing.PropertyID,
		AgentID:       currentListing.AgentID,
		Price:         float64(currentListing.Price),
		ListingStatus: currentListing.ListingStatus,
		ListingDate:   currentListing.ListingDate,
		Description:   currentListing.Description,
		AcceptedBidID: currentListing.AcceptedBidID,
	}
}

// GetPropertyByIDHandler handles the request for get property by id
//
// @Summary get property by id
// @Description get property by id
// @Tags property
// @Accept json
// @Produce json
// @Param getPropertyByIDRequest body getPropertyByIDRequest true "get property by id"
// @Success 200 {object} propertyResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /property/getPriorityByID [post]
func (s *Server) GetListingByPropertyIDHandler(c *gin.Context) {
	var req getListingByIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	listing, err := s.dbService.GetListingByPropertyID(c, req.PropertyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := getListingResponse(listing)

	c.JSON(http.StatusOK, resp)
}
