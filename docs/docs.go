// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Supuwoerc",
            "url": "https://github.com/supuwoerc",
            "email": "zhangzhouou@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/public/npm/downloads": {
            "get": {
                "description": "根据指定的时间范围和包名获取下载数据",
                "tags": [
                    "NPM数据查询"
                ],
                "summary": "获取指定时间范围内的下载数据",
                "parameters": [
                    {
                        "type": "string",
                        "description": "包名",
                        "name": "package",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "yyyy-mm-dd",
                        "description": "开始日期 (默认为7天前)",
                        "name": "start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "yyyy-mm-dd",
                        "description": "结束日期 (默认为今天)",
                        "name": "end",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功返回数据",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    }
                }
            }
        },
        "/api/public/npm/info": {
            "get": {
                "description": "获取包meta信息",
                "tags": [
                    "NPM数据查询"
                ],
                "summary": "获取包meta信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "包名",
                        "name": "package",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功返回数据",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    }
                }
            }
        },
        "/api/public/user/add": {
            "post": {
                "description": "用于添加用户",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户管理模块"
                ],
                "summary": "添加用户",
                "parameters": [
                    {
                        "description": "ADD USER INFO",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserAddDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully add",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    },
                    "400": {
                        "description": "Invalid parameters",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    }
                }
            }
        },
        "/api/public/user/list": {
            "post": {
                "description": "查询用户列表",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户管理模块"
                ],
                "summary": "查询用户列表",
                "parameters": [
                    {
                        "description": "GET USER LIST",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserListDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully get user list",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    },
                    "400": {
                        "description": "Invalid parameters",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    }
                }
            }
        },
        "/api/public/user/login": {
            "post": {
                "description": "用于用户登录系统",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户管理模块"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "User Login Info",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserLoginDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully login",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    },
                    "400": {
                        "description": "Invalid parameters",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    }
                }
            }
        },
        "/api/user/delete/{id}": {
            "delete": {
                "description": "根据用户ID删除用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理模块"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    }
                }
            }
        },
        "/api/user/update": {
            "patch": {
                "description": "更新用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理模块"
                ],
                "summary": "更新用户信息",
                "parameters": [
                    {
                        "description": "Update User Info",
                        "name": "userUpdateDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserUpdateDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully updated",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    },
                    "400": {
                        "description": "Invalid parameters",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    }
                }
            }
        },
        "/api/user/{id}": {
            "get": {
                "description": "用于根据ID查询用户",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户管理模块"
                ],
                "summary": "查询用户",
                "parameters": [
                    {
                        "description": "GET USER INFO",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.BasicIdDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully get user info",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    },
                    "400": {
                        "description": "Invalid parameters",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/serializer.BasicResponse-any"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.BasicIdDTO": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "dto.UserAddDTO": {
            "type": "object",
            "required": [
                "name",
                "password"
            ],
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "real_name": {
                    "type": "string"
                }
            }
        },
        "dto.UserListDTO": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                }
            }
        },
        "dto.UserLoginDTO": {
            "type": "object",
            "required": [
                "name",
                "password"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.UserUpdateDTO": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "real_name": {
                    "type": "string"
                }
            }
        },
        "serializer.BasicResponse-any": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Go-Parrot",
	Description:      "鹦鹉学舌",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
