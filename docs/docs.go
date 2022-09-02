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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Petrov Nikita",
            "url": "https://github.com/MyMarvel",
            "email": "nikita.petrov.programmer@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/likes/doLike": {
            "post": {
                "description": "Increments likes counter for a person in case the user didn't like them yet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Likes"
                ],
                "summary": "Adds a like to an account",
                "parameters": [
                    {
                        "description": "Params:\u003cul\u003e\u003cli\u003e\u003ci\u003eTargetAccount\u003c/i\u003e: The account name we like.\u003c/li\u003e\u003cli\u003e\u003ci\u003eCurrentAccount\u003c/i\u003e: Our account (the person who likes)\u003c/li\u003e\u003c/ul\u003e",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routes.doLikeParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Example: \u003ccode\u003e{'likes': likesCount}\u003c/code\u003e",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "208": {
                        "description": "Example: \u003ccode\u003e{'warning': 'Already liked'}\u003c/code\u003e",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Example: \u003ccode\u003e{'error': 'Error description'}\u003c/code\u003e",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/likes/getLikeCount": {
            "post": {
                "description": "Returns amount of likes collected by a certain account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Likes"
                ],
                "summary": "Get likes for an account",
                "parameters": [
                    {
                        "description": "Params:\u003cul\u003e\u003cli\u003e\u003ci\u003eTargetAccount\u003c/i\u003e: The account name we like\u003c/li\u003e\u003cul\u003e",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routes.getLikeCountParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Example: \u003ccode\u003e{'likes': likesCount}\u003c/code\u003e",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Example: \u003ccode\u003e{'error': 'Error description'}\u003c/code\u003e",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "routes.doLikeParams": {
            "type": "object",
            "required": [
                "currentAccount",
                "targetAccount"
            ],
            "properties": {
                "currentAccount": {
                    "type": "string"
                },
                "targetAccount": {
                    "type": "string"
                }
            }
        },
        "routes.getLikeCountParams": {
            "type": "object",
            "required": [
                "targetAccount"
            ],
            "properties": {
                "targetAccount": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1.0",
	Schemes:          []string{},
	Title:            "Likes API for Accounts",
	Description:      "Allows users to add likes to other users.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
