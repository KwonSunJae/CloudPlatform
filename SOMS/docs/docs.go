// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/mockup/200": {
            "get": {
                "description": "200 OK 응답을 반환합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mockup"
                ],
                "summary": "200 OK 응답",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "200 OK 응답을 반환합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mockup"
                ],
                "summary": "200 OK 응답",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/mockup/400": {
            "get": {
                "description": "40X 에러 응답을 반환합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mockup"
                ],
                "summary": "40X 에러 응답",
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "40X 에러 응답을 반환합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mockup"
                ],
                "summary": "40X 에러 응답",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/mockup/500": {
            "get": {
                "description": "50X 에러 응답을 반환합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mockup"
                ],
                "summary": "50X 에러 응답",
                "responses": {
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "50X 에러 응답을 반환합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mockup"
                ],
                "summary": "50X 에러 응답",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "사용자의 정보를 전체 조회합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "사용자 정보 전체 조회",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "사용자 가입을 진행합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "사용자 가입",
                "parameters": [
                    {
                        "description": "User Name",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UserRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "사용자 로그인을 진행합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "사용자 로그인",
                "parameters": [
                    {
                        "description": "User Login Info",
                        "name": "UserLogin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UserLoginRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/user/validate/{userID}": {
            "get": {
                "description": "사용자 ID의 유효성을 검사합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "사용자 ID 유효성 검사",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "사용자의 정보를 조회합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "사용자 정보 조회",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User uuid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "사용자의 정보를 삭제합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "사용자 정보 삭제",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User uuid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/user/{userID}": {
            "patch": {
                "description": "사용자의 정보를 수정합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "사용자 정보 수정",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "UserID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User Name",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UserRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/vm": {
            "get": {
                "description": "VM의 정보를 전체 조회합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vm"
                ],
                "summary": "VM 정보 전체 조회",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "VM을 등록합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vm"
                ],
                "summary": "VM 등록",
                "parameters": [
                    {
                        "description": "VM 정보",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vm.CreateVmBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/vm/console/{id}": {
            "get": {
                "description": "VM의 콘솔을 조회합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vm"
                ],
                "summary": "VM 콘솔 조회",
                "parameters": [
                    {
                        "type": "string",
                        "description": "VM uuid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/vm/reboot/{id}": {
            "post": {
                "description": "VM을 재부팅합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vm"
                ],
                "summary": "VM 재부팅",
                "parameters": [
                    {
                        "type": "string",
                        "description": "VM uuid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/vm/start/{id}": {
            "post": {
                "description": "VM을 시작합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vm"
                ],
                "summary": "VM 시작",
                "parameters": [
                    {
                        "type": "string",
                        "description": "VM uuid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/vm/stop/{id}": {
            "post": {
                "description": "VM을 중지합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vm"
                ],
                "summary": "VM 중지",
                "parameters": [
                    {
                        "type": "string",
                        "description": "VM uuid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/vm/vnc/{id}": {
            "get": {
                "description": "VM의 VNC를 조회합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vm"
                ],
                "summary": "VM VNC 조회",
                "parameters": [
                    {
                        "type": "string",
                        "description": "VM uuid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/vm/{id}": {
            "get": {
                "description": "VM의 정보를 조회합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vm"
                ],
                "summary": "VM 정보 조회",
                "parameters": [
                    {
                        "type": "string",
                        "description": "VM uuid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "VM의 정보를 삭제합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vm"
                ],
                "summary": "VM 정보 삭제",
                "parameters": [
                    {
                        "type": "string",
                        "description": "VM uuid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "VM의 정보를 수정합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vm"
                ],
                "summary": "VM 정보 수정",
                "parameters": [
                    {
                        "type": "string",
                        "description": "VM uuid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "VM 정보",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vm.CreateVmBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/vmstat": {
            "get": {
                "description": "VM의 상태를 조회합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vm"
                ],
                "summary": "VM 상태 조회",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "response.CommonResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {},
                "status": {
                    "type": "integer"
                }
            }
        },
        "user.UserLoginRequestBody": {
            "type": "object",
            "properties": {
                "pw": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "user.UserRequestBody": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "priority": {
                    "type": "string"
                },
                "pw": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "spot": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "vm.CreateVmBody": {
            "type": "object",
            "properties": {
                "externalIP": {
                    "type": "string"
                },
                "flavorID": {
                    "type": "string"
                },
                "internalIP": {
                    "type": "string"
                },
                "keypair": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "selectedOS": {
                    "type": "string"
                },
                "selectedSecuritygroup": {
                    "type": "string"
                },
                "unionmountImage": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "SOMS API",
	Description:      "Cloud Platform API Server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
