package server

import (
	"backend_real_estate/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"time"
)


type listingDisplayResponse struct {
	Price         string         `json:"price"`
	ListingStatus string         `json:"listing_status"`
	ListingDate   time.Time      `json:"listing_date"`
	Description   string 		 `json:"description"`
	FirstName     string         `json:"first_name"`
	LastName      string         `json:"last_name"`
	Email         string         `json:"email"`
	Address       string         `json:"address"`
	City          string         `json:"city"`
	State         string         `json:"state"`
	Zipcode       int32          `json:"zipcode"`
	Bedrooms      int32          `json:"bedrooms"`
	Bathrooms     int32          `json:"bathrooms"`
}

func getListingDisplayResponse(display []database.GetListingsRow) []listingDisplayResponse {

	var listings []listingDisplayResponse

	for i := 0; i < len(display); i++ {
		listing := listingDisplayResponse{
			Price:	   		display[i].Price,
			ListingStatus:  display[i].ListingStatus,
			ListingDate:    display[i].ListingDate,
			Description:    display[i].Description.String,
			FirstName: 		display[i].FirstName,
			LastName: 		display[i].LastName,
			Email:			display[i].Email,
			Address:   		display[i].Address,
			City:      		display[i].City,
			State:     		display[i].State,
			Zipcode:   		display[i].Zipcode,
			Bedrooms:  		display[i].Bedrooms,
			Bathrooms: 		display[i].Bathrooms,
		}

		listings = append(listings, listing)
	}

	return listings
}

// getListingDisplayHandler handles the response for getting property information
//
// @Summary Display properties
// @Description Get listings with optional pagination
// @Tags listing
// @Accept json
// @Produce json
// @Success 200 {array} listingDisplayResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /listing [get]
func (s *Server) GetListingDisplayHandler(c *gin.Context) {
	listings, err := s.dbService.GetListings(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	returnedListings := getListingDisplayResponse(listings)
	if len(returnedListings) == 0 {
		returnedListings = []listingDisplayResponse{}
	}

	c.JSON(http.StatusOK, returnedListings)

}