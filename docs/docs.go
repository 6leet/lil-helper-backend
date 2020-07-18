// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/createmission": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "create mission",
                "parameters": [
                    {
                        "description": "set mission params",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apiModel.SetMissionParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/helpermodel.Mission"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin/helpers": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "list helpers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/apiModel.JsonObjectArray"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin/helpers/{uid}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "ban helper",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User uid",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        },
        "/admin/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "user login test",
                "parameters": [
                    {
                        "description": "User login parameters",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apiModel.UserRegistParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/helpermodel.PublicUser"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin/missions": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "list missions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/apiModel.JsonObjectArray"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin/missions/{date}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "list missions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/apiModel.JsonObjectArray"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin/missions/{uid}": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "update mission",
                "parameters": [
                    {
                        "type": "string",
                        "description": "mission uid",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "set mission params",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apiModel.SetMissionParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/helpermodel.Mission"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "delete mission",
                "parameters": [
                    {
                        "type": "string",
                        "description": "mission uid",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        },
        "/admin/regist": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "user registration",
                "parameters": [
                    {
                        "description": "User registration parameters",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apiModel.UserRegistParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/helperModel.PublicUser"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin/screenshots": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "list screenshots",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/apiModel.JsonObjectArray"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin/setscreenshotapprove": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "audit screenshot",
                "parameters": [
                    {
                        "description": "audit screenshot params",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apimodel.AuditScreenshotParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/helpermodel.Screenshot"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apiModel.JsonObjectArray": {
            "type": "object",
            "properties": {
                "keys": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "objects": {
                    "type": "array",
                    "items": {
                        "type": "object"
                    }
                }
            }
        },
        "apiModel.SetMissionParam": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "content": {
                    "type": "string",
                    "example": "this is a content"
                },
                "picture": {
                    "type": "string",
                    "example": "this/is/a/path/of/picture.jpg"
                },
                "score": {
                    "type": "integer"
                },
                "weight": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "apiModel.UserRegistParam": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "my_password"
                },
                "username": {
                    "type": "string",
                    "example": "my_username"
                }
            }
        },
        "apimodel.AuditScreenshotParams": {
            "type": "object",
            "properties": {
                "approve": {
                    "type": "boolean"
                },
                "missionID example:": {
                    "type": "string"
                },
                "userID": {
                    "type": "string",
                    "example": "screenshot_userID"
                }
            }
        },
        "handler.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string",
                    "example": "ok"
                }
            }
        },
        "helperModel.PublicUser": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "admin": {
                    "type": "boolean"
                },
                "userUID": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "helpermodel.Mission": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "content": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                },
                "score": {
                    "type": "integer"
                },
                "uid": {
                    "type": "string"
                },
                "weight": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "helpermodel.PublicUser": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "admin": {
                    "type": "boolean"
                },
                "userUID": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "helpermodel.Screenshot": {
            "type": "object",
            "properties": {
                "approve": {
                    "type": "boolean"
                },
                "audit": {
                    "type": "boolean"
                },
                "date": {
                    "type": "string"
                },
                "missionID": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                },
                "uid": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.0.1",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "lil-helper swagger API",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
