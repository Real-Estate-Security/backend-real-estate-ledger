package server

import (
	"backend_real_estate/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type listingDisplayResponse struct {
	Address   string `json:"Address" binding:required`
	City      string `json:"City" binding:required`
	State     string `json:"State" binding:required`
	Zipcode   int32  `json:"Zipcode" binding:required`
	Bedrooms  int32  `json:"Bedrooms" binding:required`
	Bathrooms int32  `json:"Bathrooms" binding:required`
}

func getListingDisplayResponse(display database.Properties) listingDisplayResponse {
	return listingDisplayResponse{
		Address:   display.Address,
		City: 	   display.City,
		State: 	   display.State,
		Zipcode:   display.Zipcode,
		Bedrooms:  display.Bedrooms,
		Bathrooms: display.Bathrooms,
	}
}

// getListingDisplayHandler handles the reponse for get property information
//
// @Summary display properties
// @Description get listing by property id
// @Tags listing
// @Accept json
// @Produce json
// @Param limit query int false "Limit (default: 10)"
// @Success 200 {object} listingDisplayResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /listing/getListings [get]
func (s *Server) GetListingDisplayHandler(c *gin.Context) {
	listings, err := s.dbService.ListProperties(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	var responses []listingDisplayResponse

	for _, listing := range (listings) {
		display := getListingDisplayResponse(listing)
		
		responses = append(responses, display)
	}

	c.JSON(http.StatusOK, responses)
}
