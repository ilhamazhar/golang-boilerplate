// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/posts": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get post profile data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Posts"
                ],
                "summary": "Get Posts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create a new post",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Posts"
                ],
                "summary": "Create Post",
                "parameters": [
                    {
                        "type": "string",
                        "description": " ",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": " ",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/posts/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get details of a post by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Posts"
                ],
                "summary": "Get post by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Post"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update post data by ID",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Posts"
                ],
                "summary": "Update post by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": " ",
                        "name": "title",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": " ",
                        "name": "content",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Post"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete post data by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Posts"
                ],
                "summary": "Delete post by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Post"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/public/login": {
            "post": {
                "description": "Login to obtain an access token",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Public"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "type": "string",
                        "description": " ",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": " ",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/public/register": {
            "post": {
                "description": "Register a new user",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Public"
                ],
                "summary": "Register User",
                "parameters": [
                    {
                        "type": "string",
                        "description": " ",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": " ",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": " ",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/users": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get user profile data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get Users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/users/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get details of a user by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update user data by ID",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update user by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": " ",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": " ",
                        "name": "email",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete user data by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete user by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": " ",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Post": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "required": [
                "email",
                "name"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 8
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 3
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Golang App",
	Description:      "This is loreme seipsume.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
