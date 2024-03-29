// Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/v1/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "登录人信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "个人信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/dto.Profile"
                        }
                    }
                }
            }
        },
        "/v1/reg-count": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "按月统计当月每天注册人数",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "注册统计",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Y-d",
                        "name": "month",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "结果按日期分组",
                        "schema": {
                            "allOf": [
                                {
                                    "type": "object"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "integer"
                                        },
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dto.RegisterCount"
                                            }
                                        },
                                        "message": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/register": {
            "post": {
                "description": "手机号码注册",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "注册",
                "parameters": [
                    {
                        "description": "注册表单",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"data\":{\"access_token\":\"\",\"refresh_token\":\"\"},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "{\"code\":400,\"data\":null,\"msg\":\"参数错误\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/v1/users/{id}": {
            "get": {
                "description": "手机号码注册",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户信息",
                        "schema": {
                            "$ref": "#/definitions/dto.Profile"
                        }
                    },
                    "404": {
                        "description": "{\"code\":404,\"data\":null,\"msg\":\"未查询到结果\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.Profile": {
            "type": "object",
            "properties": {
                "avatar": {
                    "description": "头像",
                    "type": "string"
                },
                "birthday": {
                    "description": "🎂",
                    "type": "string"
                },
                "created_at": {
                    "description": "注册时间",
                    "type": "string"
                },
                "email": {
                    "description": "📮",
                    "type": "string"
                },
                "gender": {
                    "description": "性别",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "level": {
                    "description": "等级",
                    "type": "integer"
                },
                "nickname": {
                    "description": "昵称",
                    "type": "string"
                },
                "phone": {
                    "description": "手机号",
                    "type": "string"
                },
                "signature": {
                    "description": "个性签名",
                    "type": "string"
                }
            }
        },
        "dto.RegisterCount": {
            "type": "object",
            "properties": {
                "date": {
                    "description": "年-月-日",
                    "type": "string"
                },
                "total": {
                    "description": "统计结果",
                    "type": "integer"
                }
            }
        },
        "dto.RegisterParam": {
            "type": "object",
            "required": [
                "code",
                "password",
                "phone"
            ],
            "properties": {
                "code": {
                    "description": "短信码",
                    "type": "integer"
                },
                "password": {
                    "type": "string",
                    "maxLength": 18,
                    "minLength": 6
                },
                "phone": {
                    "description": "手机号",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Bingo Example API",
	Description:      "bingo案例",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
