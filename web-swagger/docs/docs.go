// Code generated by swaggo/swag. DO NOT EDIT.

package docs

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
        "/cars": {
            "get": {
                "description": "Get details of all cars",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "Get details",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Car"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new car details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "Create a new car details",
                "parameters": [
                    {
                        "description": "Car details",
                        "name": "models.Car",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Car"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Car"
                        }
                    }
                }
            }
        },
        "/cars/{Id}": {
            "get": {
                "description": "Get details of car corresponding to the input Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "Get details for a given Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the car",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Car"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a car details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "Delete a car details",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the car",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            },
            "patch": {
                "description": "Update a car details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "Update a car details",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the car",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Car"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Car": {
            "type": "object",
            "properties": {
                "harga": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "merk": {
                    "type": "string"
                },
                "pemilik": {
                    "type": "string"
                },
                "typecars": {
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
