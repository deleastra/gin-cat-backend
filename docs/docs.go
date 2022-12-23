// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "email": "support@example.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cats": {
            "get": {
                "description": "Retrieves a list of all cats stored in the database.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cats"
                ],
                "summary": "Retrieves a list of all cats",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Cats"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new cat and stores it in the database.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cats"
                ],
                "summary": "Creates a new cat",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the cat",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Image of the cat",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Cats"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/cats/{id}": {
            "get": {
                "description": "Retrieves a single cat by ID from the database.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cats"
                ],
                "summary": "Retrieves a single cat by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the cat",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Cats"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Updates a cat by ID and stores the changes in the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cats"
                ],
                "summary": "Updates a cat by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the cat",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated cat information",
                        "name": "cat",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Cats"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Cats"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a cat from the database by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cats"
                ],
                "summary": "Deletes a cat by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the cat",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "models.Cats": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "description": "ID is the primary key and auto-incrementing field for the cats table in the database.",
                    "type": "integer"
                },
                "image": {
                    "description": "Image is the file path or URL of an image for the cat.",
                    "type": "string"
                },
                "name": {
                    "description": "Name is the name of the cat.",
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Cat API",
	Description:      "A simple API for managing cats.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
