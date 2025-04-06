package server

import (
	"backend_real_estate/internal/database"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createBidRequest struct {
	ListingID     int64  `json:"ListingID" binding:"required"`
	BuyerID       int64  `json:"BuyerID" binding:"required"`
	AgentID       int64  `json:"AgentID" binding:"required"`
	Amount        string `json:"Amount" binding:"required"`
	PreviousBidID int64  `json:"PreviousBidID,omitempty"`
}

type bidResponse struct {
	ID            int64  `json:"ID" binding:"required"`
	ListingID     int64  `json:"ListingID" binding:"required"`
	BuyerID       int64  `json:"BuyerID" binding:"required"`
	AgentID       int64  `json:"AgentID" binding:"required"`
	Amount        string `json:"Amount" binding:"required"`
	PreviousBidID int64  `json:"PreviousBidID" binding:"required"`
}

func createBidResponse(currentBid database.Bids) bidResponse {
	return bidResponse{
		ID: currentBid.ID,
		ListingID: currentBid.ListingID,
		BuyerID:   currentBid.BuyerID,
		AgentID:   currentBid.AgentID,
		Amount:    currentBid.Amount,
		PreviousBidID: func() int64 {
			if currentBid.PreviousBidID.Valid {
				return currentBid.PreviousBidID.Int64
			}
			return 0 // or another default value
		}(),
	}
}

// createBidHandler handles the request for create a bid
//
// @Summary create a bid
// @Description create a bid
// @Tags bidding
// @Accept json
// @Produce json
// @Param createBidRequest body createBidRequest true "create a bid"
// @Success 200 {object} bidResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /bidding/createBid [post]
func (s *Server) createBidHandler(c *gin.Context) {
	var req createBidRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var dbParam database.CreateBidParams
	dbParam.AgentID = req.AgentID
	dbParam.Amount = req.Amount
	dbParam.BuyerID = req.BuyerID
	dbParam.ListingID = req.ListingID
	dbParam.PreviousBidID = sql.NullInt64{
		Int64: req.PreviousBidID,
		Valid: req.PreviousBidID != 0, // Adjust logic if 0 is a valid value
	}

	bid, err := s.dbService.CreateBid(c, dbParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := createBidResponse(bid)

	c.JSON(http.StatusOK, resp)
}
