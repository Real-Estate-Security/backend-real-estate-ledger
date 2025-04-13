package server

import (
	"backend_real_estate/internal/database"
	"backend_real_estate/internal/token"
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

type listBidsRequest struct {
	BuyerID  int64  `json:"BuyerID" binding:"required"`
	Username string `json:"Username" binding:"required,alphanum"`
}

type listBidsOnListingRequest struct {
	ListingID int64 `json:"ListingID" binding:"required"`
}

type rejectBidRequest struct {
	ID int64 `json:"ID" binding:"required"`
}

type bidResponse struct {
	ID            int64  `json:"ID" binding:"required"`
	ListingID     int64  `json:"ListingID" binding:"required"`
	BuyerID       int64  `json:"BuyerID" binding:"required"`
	AgentID       int64  `json:"AgentID" binding:"required"`
	Amount        string `json:"Amount" binding:"required"`
	PreviousBidID int64  `json:"PreviousBidID" binding:"required"`
}

type listBidResponse struct {
	ID            int64  `json:"ID" binding:"required"`
	ListingID     int64  `json:"ListingID" binding:"required"`
	BuyerID       int64  `json:"BuyerID" binding:"required"`
	AgentID       int64  `json:"AgentID" binding:"required"`
	Amount        string `json:"Amount" binding:"required"`
	PreviousBidID int64  `json:"PreviousBidID" binding:"required"`
	Status        string `json:"Status" binding:"required"`
}

func createBidResponse(currentBid database.Bids) bidResponse {
	return bidResponse{
		ID:        currentBid.ID,
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

// listBidsHandler handles the request for listing all bids belonging to a given buyer
//
// @Summary given user, list all bid with them as buyer
// @Description listing all bids belonging to a given buyer
// @Tags bidding
// @Accept json
// @Produce json
// @Param buyerID body listBidsRequest true "listing all bids belonging to a given buyer"
// @Success 200 {array} listBidResponse "list of bids"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /bidding/listBids [post]
// @Security BearerAuth
func (s *Server) listBidsHandler(c *gin.Context) {
	var req listBidsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Check if the client exists
	client, err := s.dbService.GetUserByUsername(c, req.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	// Get the agent's username from the authorization payload
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)

	if authPayload== nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization payload"})
		return
	}

	var bidList []database.Bids
	bidList, err = s.dbService.ListBids(c, client.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Convert database.Representations to RepresentationsWithNullableTime
	var response []listBidResponse

	for _, r := range bidList {
		response = append(response, listBidResponse{
			ID:            r.ID,
			ListingID:     r.ListingID,
			AgentID:       r.AgentID,
			BuyerID:       r.BuyerID,
			Amount:        r.Amount,
			PreviousBidID: r.PreviousBidID.Int64,
			Status:        string(r.Status),
		})
	}

	c.JSON(http.StatusOK, response)
}

// listBidsOnListingHandler handles the request for listing all bids belonging to a given buyer
//
// @Summary given listing, list all bids with that as the listing
// @Description listing all bids with a given listing
// @Tags bidding
// @Accept json
// @Produce json
// @Param listingID body listBidsOnListingRequest true "listing all bids that have a specific listing"
// @Success 200 {array} listBidResponse "list of bids"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /bidding/listBidsOnListing [post]
func (s *Server) listBidsOnListingHandler(c *gin.Context) {
	var req listBidsOnListingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var bidList []database.Bids

	// Fetch representations based on the user's role

	bidList, err := s.dbService.ListBidsOnListing(c, req.ListingID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Convert database.Representations to RepresentationsWithNullableTime
	var response []listBidResponse

	for _, r := range bidList {
		response = append(response, listBidResponse{
			ID:            r.ID,
			ListingID:     r.ListingID,
			AgentID:       r.AgentID,
			BuyerID:       r.BuyerID,
			Amount:        r.Amount,
			PreviousBidID: r.PreviousBidID.Int64,
			Status:        string(r.Status),
		})
	}

	c.JSON(http.StatusOK, response)
}

// rejectBidHandler handles the request to reject a bid
//
// @Summary reject a bid
// @Description reject a bid
// @Tags bidding
// @Accept json
// @Produce json
// @Param rejectBidRequest body rejectBidRequest true "reject a bid"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /bidding/rejectBid [put]
func (s *Server) rejectBidHandler(c *gin.Context) {
	var req rejectBidRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var bidID = req.ID

	err := s.dbService.RejectBid(c, bidID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, bidID)
}

// acceptBidHandler handles the request to accept a bid
//
// @Summary accept a bid
// @Description accept a bid
// @Tags bidding
// @Accept json
// @Produce json
// @Param rejectBidRequest body rejectBidRequest true "accept a bid"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /bidding/acceptBid [put]
func (s *Server) acceptBidHandler(c *gin.Context) {
	var req rejectBidRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var bidID = req.ID

	err := s.dbService.AcceptBid(c, bidID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, bidID)
}
