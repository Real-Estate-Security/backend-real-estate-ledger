package server

import (
	"backend_real_estate/internal/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type listingDisplayResponse struct {
	Owner     int64  `json:"Owner" binding:required"`
	Address   string `json: Address" binding:required`
	City      string `json:"City" binding:required`
	State     string `json:"State" binding:required`
	Zipcode   int32  `json:"Zipcode" binding:required`
	Bedrooms  int32  `json:"Bedrooms" binding:required`
	Bathrooms int32  `json:"Bathrooms" binding:required`
}

func getListingDisplayResponse(display database.Properties) listingResponse {
	return listingDisplayResponse{
		Owner: 
	}
}
