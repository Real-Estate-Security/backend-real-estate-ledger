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
        "/example/helloworld": {
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
        "/property/getPriorityByID": {
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
