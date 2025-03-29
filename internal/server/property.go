package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Property struct {
	ID      string         `json:"ID"`
	Address string         `json:"Address"`
	Owner   string         `json:"Owner"`
	Agent   string         `json:"Agent"`
	State   string         `json:"State"`
	Bids    map[string]Bid `json:"Bids"`
}

type Bid struct {
	ID              string `json:"ID"`
	Amount          int    `json:"Amount"`
	Bidder          string `json:"Bidder"`
	Agent           string `json:"Agent"`
	BuyerCountered  bool   `json:"BuyerCountered"`
	SellerCountered bool   `json:"SellerCountered"`
}

func (s *Server) GetAllProperties(c *gin.Context) {
	network := s.gwService.GetNetwork("mychannel")
	contract := network.GetContract("realestatesec")

	evaluateResult, err := contract.EvaluateTransaction("ViewProperties")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Could not get properties")
		return
	}

	c.Data(http.StatusOK, "application/json", evaluateResult)
}

func (s *Server) RegisterProperty(c *gin.Context) {

	network := s.gwService.GetNetwork("mychannel")
	contract := network.GetContract("realestatesec")

	var property Property
	c.ShouldBindBodyWithJSON(&property)

	//Maybe change chaincode params to take in property object as a whole
	_, err := contract.SubmitTransaction("RegisterProperty", property.ID, property.Address, property.Owner, property.Agent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Could not register property")
		return
	}

	//maybe send back the data of the registered property
	c.Data(http.StatusOK, "plain/text", []byte("Property registered"))
}

func (s *Server) ListProperty(c *gin.Context) {
	network := s.gwService.GetNetwork("mychannel")
	contract := network.GetContract("realestatesec")
	propertyID := c.Param("id")

	_, err := contract.SubmitTransaction("ListProperty", propertyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Could not list property")
		return
	}

	c.Data(http.StatusOK, "plain/text", []byte("Property listed"))
}

func (s *Server) PlaceBid(c *gin.Context) {
	network := s.gwService.GetNetwork("mychannel")
	contract := network.GetContract("realestatesec")
	propertyID := c.Param("id")

	var bid Bid
	c.ShouldBindBodyWithJSON(&bid)

	_, err := contract.SubmitTransaction("PlaceBid", propertyID, bid.ID, strconv.Itoa(bid.Amount), bid.Bidder, bid.Agent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Could not place bid")
		return
	}

	c.Data(http.StatusOK, "plain/text", []byte("Bid placed"))
}

func (s *Server) RejectBid(c *gin.Context) {
	network := s.gwService.GetNetwork("mychannel")
	contract := network.GetContract("realestatesec")
	propertyID := c.Param("propertyID")
	bidID := c.Param("bidID")

	_, err := contract.SubmitTransaction("RejectBid", propertyID, bidID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Could not reject bid")
		return
	}

	c.Data(http.StatusOK, "plain/text", []byte("Bid rejected"))
}
