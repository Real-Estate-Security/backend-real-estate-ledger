package server

import (
	"backend_real_estate/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getPropertyByIDRequest struct {
	PropertyID int64  `json:"PropertyID" binding:"required"`
	Username   string `json:"Username" binding:"required,alphanum"`
}

type propertyResponse struct {
	ID             int64  `json:"ID" binding:"required"`
	Owner          int64  `json:"Owner" binding:"required"`
	Address        string `json:"Address" binding:"required"`
	City           string `json:"City" binding:"required"`
	State          string `json:"State" binding:"required"`
	ZipCode        int    `json:"ZipCode" binding:"required"`
	NumOfBedrooms  int    `json:"NumOfBedrooms" binding:"required"`
	NumOfBathrooms int    `json:"NumOfBathrooms" binding:"required"`
}

func getPropertyResponse(currentProperty database.Properties) propertyResponse {
	return propertyResponse{
		ID:             currentProperty.ID,
		Owner:          currentProperty.Owner,
		Address:        currentProperty.Address,
		City:           currentProperty.City,
		State:          currentProperty.State,
		ZipCode:        int(currentProperty.Zipcode),
		NumOfBedrooms:  int(currentProperty.Bedrooms),
		NumOfBathrooms: int(currentProperty.Bathrooms),
	}
}

// getPropertyByIDHandler handles the request for get property by id
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
// @Router /property/getPropertyByID [post]
func (s *Server) getPropertyByIDHandler(c *gin.Context) {
	var req getPropertyByIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	property, err := s.dbService.GetPropertyByID(c, req.PropertyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := getPropertyResponse(property)

	c.JSON(http.StatusOK, resp)
}
