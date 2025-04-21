package server

import (
	"backend_real_estate/internal/database"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type getListingByIDRequest struct {
	PropertyID int64  `json:"PropertyID" binding:"required"`
	Username   string `json:"Username" binding:"required,alphanum"`
}

type updateAcceptedBidIdByListingIdRequest struct {
	ListId        int64 `json:"ListId" binding:"required"`
	AcceptedBidId int64 `json:"AcceptedBidId" binding:"required"`
}

type listingResponse struct {
	ID            int64     `json:"ID" binding:"required"`
	PropertyID    int64     `json:"PropertyID" binding:"required"`
	AgentID       int64     `json:"AgentID" binding:"required"`
	Price         string    `json:"Price" binding:"required"`
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
		Price:         currentListing.Price,
		ListingStatus: currentListing.ListingStatus,
		ListingDate:   currentListing.ListingDate,
		Description: func() string {
			if currentListing.Description.Valid {
				return currentListing.Description.String
			}
			return "" // or another default value
		}(),
		AcceptedBidID: func() int64 {
			if currentListing.AcceptedBidID.Valid {
				return currentListing.AcceptedBidID.Int64
			}
			return 0 // or another default value
		}(),
	}
}

// getListingByPropertyIDHandler handles the request for get listing by property id
//
// @Summary get listing by property id
// @Description get listing by property id
// @Tags listing
// @Accept json
// @Produce json
// @Param getListingByIDRequest body getListingByIDRequest true "get listig by property id"
// @Success 200 {object} listingResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /listing/getListingByPropertyID [post]
func (s *Server) getListingByPropertyIDHandler(c *gin.Context) {
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

// updateAcceptedBidIdByListingIdHandler handles the request to update accepted bid id for a given listing id
//
// @Summary update accepted bid id for a given listing id
// @Description update accepted bid id for a given listing id
// @Tags listing
// @Accept json
// @Produce json
// @Param updateAcceptedBidIdByListingIdRequest body updateAcceptedBidIdByListingIdRequest true "update accepted bid id for a given listing id"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /listing/updateAcceptedBidIdByListingId [post]
// @Security BearerAuth
func (s *Server) updateAcceptedBidIdByListingIdHandler(c *gin.Context) {
	var req updateAcceptedBidIdByListingIdRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	updateAcceptedBidIdByListingIdParameters := database.UpdateAcceptedBidIdByListingIdParams{
		AcceptedBidID: sql.NullInt64{Int64: req.AcceptedBidId, Valid: true},
		ID:            req.ListId,
	}

	err := s.dbService.UpdateAcceptedBidIdByListingId(c, updateAcceptedBidIdByListingIdParameters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, req.ListId)
}
