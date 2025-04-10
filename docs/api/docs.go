// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/agent/accept-representation/{id}": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Allows an agent to accept a representation request.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "agent"
                ],
                "summary": "Accept representation request",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Representation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Representation request accepted successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Representation not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/agent/decline-representation/{id}": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Allows an agent to decline a representation request.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "agent"
                ],
                "summary": "Decline representation request",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Representation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Representation request declined successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Representation not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/agent/representation": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Fetches all representations for the authenticated user, whether they are an agent or a regular user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "representations"
                ],
                "summary": "List representations",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit (default: 10)",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset (default: 0)",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of representations",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/server.RepresentationData"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/agent/request-representation": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Allows an agent to request representation for a user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "agent"
                ],
                "summary": "Request representation",
                "parameters": [
                    {
                        "description": "Request Representation Request",
                        "name": "requestAgentRepresentationRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.requestAgentRepresentationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Representation request submitted successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Client not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/bidding/acceptBid": {
            "put": {
                "description": "accept a bid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bidding"
                ],
                "summary": "accept a bid",
                "parameters": [
                    {
                        "description": "accept a bid",
                        "name": "rejectBidRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.rejectBidRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/bidding/createBid": {
            "post": {
                "description": "create a bid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bidding"
                ],
                "summary": "create a bid",
                "parameters": [
                    {
                        "description": "create a bid",
                        "name": "createBidRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.createBidRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.bidResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/bidding/listBids": {
            "post": {
                "description": "listing all bids belonging to a given buyer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bidding"
                ],
                "summary": "given user, list all bid with them as buyer",
                "parameters": [
                    {
                        "description": "listing all bids belonging to a given buyer",
                        "name": "buyerID",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.listBidsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "list of bids",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/server.listBidResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/bidding/listBidsOnListing": {
            "post": {
                "description": "listing all bids with a given listing",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bidding"
                ],
                "summary": "given listing, list all bids with that as the listing",
                "parameters": [
                    {
                        "description": "listing all bids that have a specific listing",
                        "name": "listingID",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.listBidsOnListingRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "list of bids",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/server.listBidResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/bidding/rejectBid": {
            "put": {
                "description": "reject a bid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bidding"
                ],
                "summary": "reject a bid",
                "parameters": [
                    {
                        "description": "reject a bid",
                        "name": "rejectBidRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.rejectBidRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Returns the health status of the server",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/hello-world": {
            "get": {
                "description": "HelloWorld example",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "HelloWorld example",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/listing/getListingByPropertyID": {
            "post": {
                "description": "get listing by property id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "listing"
                ],
                "summary": "get listing by property id",
                "parameters": [
                    {
                        "description": "get listig by property id",
                        "name": "getListingByIDRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.getListingByIDRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.listingResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/property/getPropertyByID": {
            "post": {
                "description": "get property by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "property"
                ],
                "summary": "get property by id",
                "parameters": [
                    {
                        "description": "get property by id",
                        "name": "getPropertyByIDRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.getPropertyByIDRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.propertyResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Authenticates a user and returns an access token along with user details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Login request",
                        "name": "loginUserRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.loginUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful login",
                        "schema": {
                            "$ref": "#/definitions/server.loginUserResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/me": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieves the information of the authenticated user based on the authorization token provided.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get authenticated user information",
                "responses": {
                    "200": {
                        "description": "Successfully retrieved user information",
                        "schema": {
                            "$ref": "#/definitions/server.userMeResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/signup": {
            "post": {
                "description": "Create a new user with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "Create User Request",
                        "name": "createUserRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.createUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.userResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "server.NullableTime": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "type": "boolean"
                }
            }
        },
        "server.RepresentationData": {
            "type": "object",
            "properties": {
                "agent_first_name": {
                    "type": "string"
                },
                "agent_id": {
                    "type": "integer"
                },
                "agent_last_name": {
                    "type": "string"
                },
                "agent_username": {
                    "type": "string"
                },
                "client_first_name": {
                    "type": "string"
                },
                "client_id": {
                    "type": "integer"
                },
                "client_last_name": {
                    "type": "string"
                },
                "client_username": {
                    "type": "string"
                },
                "end_date": {
                    "$ref": "#/definitions/server.NullableTime"
                },
                "id": {
                    "type": "integer"
                },
                "is_active": {
                    "type": "boolean"
                },
                "requested_at": {
                    "type": "string"
                },
                "signed_at": {
                    "$ref": "#/definitions/server.NullableTime"
                },
                "start_date": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "server.bidResponse": {
            "type": "object",
            "required": [
                "AgentID",
                "Amount",
                "BuyerID",
                "ID",
                "ListingID",
                "PreviousBidID"
            ],
            "properties": {
                "AgentID": {
                    "type": "integer"
                },
                "Amount": {
                    "type": "string"
                },
                "BuyerID": {
                    "type": "integer"
                },
                "ID": {
                    "type": "integer"
                },
                "ListingID": {
                    "type": "integer"
                },
                "PreviousBidID": {
                    "type": "integer"
                }
            }
        },
        "server.createBidRequest": {
            "type": "object",
            "required": [
                "AgentID",
                "Amount",
                "BuyerID",
                "ListingID"
            ],
            "properties": {
                "AgentID": {
                    "type": "integer"
                },
                "Amount": {
                    "type": "string"
                },
                "BuyerID": {
                    "type": "integer"
                },
                "ListingID": {
                    "type": "integer"
                },
                "PreviousBidID": {
                    "type": "integer"
                }
            }
        },
        "server.createUserRequest": {
            "type": "object",
            "required": [
                "dob",
                "email",
                "first_name",
                "last_name",
                "password",
                "role",
                "username"
            ],
            "properties": {
                "dob": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "role": {
                    "type": "string",
                    "enum": [
                        "user",
                        "agent"
                    ]
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "server.getListingByIDRequest": {
            "type": "object",
            "required": [
                "PropertyID",
                "Username"
            ],
            "properties": {
                "PropertyID": {
                    "type": "integer"
                },
                "Username": {
                    "type": "string"
                }
            }
        },
        "server.getPropertyByIDRequest": {
            "type": "object",
            "required": [
                "PropertyID",
                "Username"
            ],
            "properties": {
                "PropertyID": {
                    "type": "integer"
                },
                "Username": {
                    "type": "string"
                }
            }
        },
        "server.listBidResponse": {
            "type": "object",
            "required": [
                "AgentID",
                "Amount",
                "BuyerID",
                "ID",
                "ListingID",
                "PreviousBidID",
                "Status"
            ],
            "properties": {
                "AgentID": {
                    "type": "integer"
                },
                "Amount": {
                    "type": "string"
                },
                "BuyerID": {
                    "type": "integer"
                },
                "ID": {
                    "type": "integer"
                },
                "ListingID": {
                    "type": "integer"
                },
                "PreviousBidID": {
                    "type": "integer"
                },
                "Status": {
                    "type": "string"
                }
            }
        },
        "server.listBidsOnListingRequest": {
            "type": "object",
            "required": [
                "ListingID"
            ],
            "properties": {
                "ListingID": {
                    "type": "integer"
                }
            }
        },
        "server.listBidsRequest": {
            "type": "object",
            "required": [
                "BuyerID",
                "Username"
            ],
            "properties": {
                "BuyerID": {
                    "type": "integer"
                },
                "Username": {
                    "type": "string"
                }
            }
        },
        "server.listingResponse": {
            "type": "object",
            "required": [
                "AcceptedBidID",
                "AgentID",
                "Description",
                "ID",
                "ListingDate",
                "ListingStatus",
                "Price",
                "PropertyID"
            ],
            "properties": {
                "AcceptedBidID": {
                    "type": "integer"
                },
                "AgentID": {
                    "type": "integer"
                },
                "Description": {
                    "type": "string"
                },
                "ID": {
                    "type": "integer"
                },
                "ListingDate": {
                    "type": "string"
                },
                "ListingStatus": {
                    "type": "string"
                },
                "Price": {
                    "type": "string"
                },
                "PropertyID": {
                    "type": "integer"
                }
            }
        },
        "server.loginUserRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "server.loginUserResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/server.userResponse"
                }
            }
        },
        "server.propertyResponse": {
            "type": "object",
            "required": [
                "Address",
                "City",
                "ID",
                "NumOfBathrooms",
                "NumOfBedrooms",
                "Owner",
                "State",
                "ZipCode"
            ],
            "properties": {
                "Address": {
                    "type": "string"
                },
                "City": {
                    "type": "string"
                },
                "ID": {
                    "type": "integer"
                },
                "NumOfBathrooms": {
                    "type": "integer"
                },
                "NumOfBedrooms": {
                    "type": "integer"
                },
                "Owner": {
                    "type": "integer"
                },
                "State": {
                    "type": "string"
                },
                "ZipCode": {
                    "type": "integer"
                }
            }
        },
        "server.rejectBidRequest": {
            "type": "object",
            "required": [
                "ID"
            ],
            "properties": {
                "ID": {
                    "type": "integer"
                }
            }
        },
        "server.requestAgentRepresentationRequest": {
            "type": "object",
            "required": [
                "client_username",
                "end_date",
                "start_date"
            ],
            "properties": {
                "client_username": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                }
            }
        },
        "server.userMeResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/server.userResponse"
                }
            }
        },
        "server.userResponse": {
            "type": "object",
            "required": [
                "dob",
                "email",
                "first_name",
                "last_name",
                "role",
                "username"
            ],
            "properties": {
                "dob": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "role": {
                    "type": "string",
                    "enum": [
                        "user",
                        "agent"
                    ]
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
