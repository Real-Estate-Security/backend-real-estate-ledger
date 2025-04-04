package server

import (
	"backend_real_estate/internal/database"
	"net/http"
	"github.com/gin-gonic/gin"
)

type createBidRequest struct {
	ListingID     int64         `json:"listing_id" binding:"required"`
	BuyerID       int64         `json:"buyer_id" binding:"required"`
	AgentID       int64         `json:"agent_id" binding:"required"`
	Amount        string        `json:"amount" binding:"required"`
	PreviousBidID int64 `json:"previous_bid_id binding:"required""`
}

type bidResponse struct {
	ListingID     int64         `json:"listing_id" binding:"required"`
	BuyerID       int64         `json:"buyer_id" binding:"required"`
	AgentID       int64         `json:"agent_id" binding:"required"`
	Amount        string        `json:"amount" binding:"required"`
	PreviousBidID int64 `json:"previous_bid_id binding:"required""`
}

func createBidResponse(currentBid database.Bids) bidResponse {
	return bidResponse{
		ListingID:  currentBid.ListingID,
		BuyerID: currentBid.BuyerID,
		AgentID: currentBid.AgentID,
		Amount: currentBid.Amount,
		PreviousBidID: int(currentBid.PreviousBidID),
		
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
	dbParam.PreviousBidID = req.PreviousBidID

	property, err := s.dbService.CreateBid(c, dbParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := getPropertyResponse(property)

	c.JSON(http.StatusOK, resp)
}
