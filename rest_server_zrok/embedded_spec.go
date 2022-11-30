// Code generated by go-swagger; DO NOT EDIT.

package rest_server_zrok

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/zrok.v1+json"
  ],
  "produces": [
    "application/zrok.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "zrok client access",
    "title": "zrok",
    "version": "0.3.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/access": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "service"
        ],
        "operationId": "access",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/accessRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "access created",
            "schema": {
              "$ref": "#/definitions/accessResponse"
            }
          },
          "401": {
            "description": "unauthorized"
          },
          "404": {
            "description": "not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/disable": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "environment"
        ],
        "operationId": "disable",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/disableRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "environment disabled"
          },
          "401": {
            "description": "invalid environment"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/enable": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "environment"
        ],
        "operationId": "enable",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/enableRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "environment enabled",
            "schema": {
              "$ref": "#/definitions/enableResponse"
            }
          },
          "401": {
            "description": "unauthorized"
          },
          "404": {
            "description": "account not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/invite": {
      "post": {
        "tags": [
          "account"
        ],
        "operationId": "invite",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/inviteRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "invitation created"
          },
          "400": {
            "description": "invitation not created (already exists)"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/login": {
      "post": {
        "tags": [
          "account"
        ],
        "operationId": "login",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/loginRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "login successful",
            "schema": {
              "$ref": "#/definitions/loginResponse"
            }
          },
          "401": {
            "description": "invalid login"
          }
        }
      }
    },
    "/overview": {
      "get": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "metadata"
        ],
        "operationId": "overview",
        "responses": {
          "200": {
            "description": "overview returned",
            "schema": {
              "$ref": "#/definitions/environmentServicesList"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/register": {
      "post": {
        "tags": [
          "account"
        ],
        "operationId": "register",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/registerRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "account created",
            "schema": {
              "$ref": "#/definitions/registerResponse"
            }
          },
          "404": {
            "description": "request not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/service": {
      "get": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "service"
        ],
        "operationId": "getService",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/serviceRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/service03"
            }
          },
          "401": {
            "description": "unauthorized"
          },
          "404": {
            "description": "not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/share": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "service"
        ],
        "operationId": "share",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/shareRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "service created",
            "schema": {
              "$ref": "#/definitions/shareResponse"
            }
          },
          "401": {
            "description": "unauthorized"
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/unaccess": {
      "delete": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "service"
        ],
        "operationId": "unaccess",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/unaccessRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "access removed"
          },
          "401": {
            "description": "unauthorized"
          },
          "404": {
            "description": "not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/unshare": {
      "delete": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "service"
        ],
        "operationId": "unshare",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/unshareRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "service removed"
          },
          "401": {
            "description": "unauthorized"
          },
          "404": {
            "description": "not found"
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/verify": {
      "post": {
        "tags": [
          "account"
        ],
        "operationId": "verify",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/verifyRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "token ready",
            "schema": {
              "$ref": "#/definitions/verifyResponse"
            }
          },
          "404": {
            "description": "token not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/version": {
      "get": {
        "tags": [
          "metadata"
        ],
        "operationId": "version",
        "responses": {
          "200": {
            "description": "current server version",
            "schema": {
              "$ref": "#/definitions/version"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "accessRequest": {
      "type": "object",
      "properties": {
        "envZId": {
          "type": "string"
        },
        "svcToken": {
          "type": "string"
        }
      }
    },
    "accessResponse": {
      "type": "object",
      "properties": {
        "frontendToken": {
          "type": "string"
        }
      }
    },
    "authUser": {
      "type": "object",
      "properties": {
        "password": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "disableRequest": {
      "type": "object",
      "properties": {
        "identity": {
          "type": "string"
        }
      }
    },
    "enableRequest": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string"
        },
        "host": {
          "type": "string"
        }
      }
    },
    "enableResponse": {
      "type": "object",
      "properties": {
        "cfg": {
          "type": "string"
        },
        "identity": {
          "type": "string"
        }
      }
    },
    "environment": {
      "type": "object",
      "properties": {
        "active": {
          "type": "boolean"
        },
        "address": {
          "type": "string"
        },
        "createdAt": {
          "type": "integer"
        },
        "description": {
          "type": "string"
        },
        "host": {
          "type": "string"
        },
        "updatedAt": {
          "type": "integer"
        },
        "zId": {
          "type": "string"
        }
      }
    },
    "environmentServices": {
      "type": "object",
      "properties": {
        "environment": {
          "$ref": "#/definitions/environment"
        },
        "services": {
          "$ref": "#/definitions/services"
        }
      }
    },
    "environmentServicesList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/environmentServices"
      }
    },
    "environments": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/environment"
      }
    },
    "errorMessage": {
      "type": "string"
    },
    "inviteRequest": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        }
      }
    },
    "loginRequest": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "loginResponse": {
      "type": "string"
    },
    "principal": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "registerRequest": {
      "type": "object",
      "properties": {
        "password": {
          "type": "string"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "registerResponse": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "service": {
      "type": "object",
      "properties": {
        "backendProxyEndpoint": {
          "type": "string"
        },
        "createdAt": {
          "type": "integer"
        },
        "frontendEndpoint": {
          "type": "string"
        },
        "metrics": {
          "$ref": "#/definitions/serviceMetrics"
        },
        "token": {
          "type": "string"
        },
        "updatedAt": {
          "type": "integer"
        },
        "zId": {
          "type": "string"
        }
      }
    },
    "service03": {
      "type": "object",
      "properties": {
        "backendMode": {
          "type": "string"
        },
        "backendProxyEndpoint": {
          "type": "string"
        },
        "frontendEndpoint": {
          "type": "string"
        },
        "frontendSelection": {
          "type": "string"
        },
        "reserved": {
          "type": "boolean"
        },
        "shareMode": {
          "type": "string"
        },
        "token": {
          "type": "string"
        },
        "zId": {
          "type": "string"
        }
      }
    },
    "serviceMetrics": {
      "type": "array",
      "items": {
        "type": "integer"
      }
    },
    "serviceRequest": {
      "type": "object",
      "properties": {
        "envZId": {
          "type": "string"
        },
        "svcToken": {
          "type": "string"
        }
      }
    },
    "services": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/service"
      }
    },
    "shareRequest": {
      "type": "object",
      "properties": {
        "authScheme": {
          "type": "string"
        },
        "authUsers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/authUser"
          }
        },
        "backendMode": {
          "type": "string",
          "enum": [
            "proxy",
            "web",
            "dav"
          ]
        },
        "backendProxyEndpoint": {
          "type": "string"
        },
        "envZId": {
          "type": "string"
        },
        "frontendSelection": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "reserved": {
          "type": "boolean"
        },
        "shareMode": {
          "type": "string",
          "enum": [
            "public",
            "private"
          ]
        }
      }
    },
    "shareResponse": {
      "type": "object",
      "properties": {
        "frontendProxyEndpoint": {
          "type": "string"
        },
        "svcToken": {
          "type": "string"
        }
      }
    },
    "unaccessRequest": {
      "type": "object",
      "properties": {
        "envZId": {
          "type": "string"
        },
        "frontendToken": {
          "type": "string"
        },
        "svcToken": {
          "type": "string"
        }
      }
    },
    "unshareRequest": {
      "type": "object",
      "properties": {
        "envZId": {
          "type": "string"
        },
        "reserved": {
          "type": "boolean"
        },
        "svcToken": {
          "type": "string"
        }
      }
    },
    "verifyRequest": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "verifyResponse": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        }
      }
    },
    "version": {
      "type": "string"
    }
  },
  "securityDefinitions": {
    "key": {
      "type": "apiKey",
      "name": "x-token",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/zrok.v1+json"
  ],
  "produces": [
    "application/zrok.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "zrok client access",
    "title": "zrok",
    "version": "0.3.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/access": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "service"
        ],
        "operationId": "access",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/accessRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "access created",
            "schema": {
              "$ref": "#/definitions/accessResponse"
            }
          },
          "401": {
            "description": "unauthorized"
          },
          "404": {
            "description": "not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/disable": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "environment"
        ],
        "operationId": "disable",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/disableRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "environment disabled"
          },
          "401": {
            "description": "invalid environment"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/enable": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "environment"
        ],
        "operationId": "enable",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/enableRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "environment enabled",
            "schema": {
              "$ref": "#/definitions/enableResponse"
            }
          },
          "401": {
            "description": "unauthorized"
          },
          "404": {
            "description": "account not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/invite": {
      "post": {
        "tags": [
          "account"
        ],
        "operationId": "invite",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/inviteRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "invitation created"
          },
          "400": {
            "description": "invitation not created (already exists)"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/login": {
      "post": {
        "tags": [
          "account"
        ],
        "operationId": "login",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/loginRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "login successful",
            "schema": {
              "$ref": "#/definitions/loginResponse"
            }
          },
          "401": {
            "description": "invalid login"
          }
        }
      }
    },
    "/overview": {
      "get": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "metadata"
        ],
        "operationId": "overview",
        "responses": {
          "200": {
            "description": "overview returned",
            "schema": {
              "$ref": "#/definitions/environmentServicesList"
            }
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/register": {
      "post": {
        "tags": [
          "account"
        ],
        "operationId": "register",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/registerRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "account created",
            "schema": {
              "$ref": "#/definitions/registerResponse"
            }
          },
          "404": {
            "description": "request not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/service": {
      "get": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "service"
        ],
        "operationId": "getService",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/serviceRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/service03"
            }
          },
          "401": {
            "description": "unauthorized"
          },
          "404": {
            "description": "not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/share": {
      "post": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "service"
        ],
        "operationId": "share",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/shareRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "service created",
            "schema": {
              "$ref": "#/definitions/shareResponse"
            }
          },
          "401": {
            "description": "unauthorized"
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/unaccess": {
      "delete": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "service"
        ],
        "operationId": "unaccess",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/unaccessRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "access removed"
          },
          "401": {
            "description": "unauthorized"
          },
          "404": {
            "description": "not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/unshare": {
      "delete": {
        "security": [
          {
            "key": []
          }
        ],
        "tags": [
          "service"
        ],
        "operationId": "unshare",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/unshareRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "service removed"
          },
          "401": {
            "description": "unauthorized"
          },
          "404": {
            "description": "not found"
          },
          "500": {
            "description": "internal server error",
            "schema": {
              "$ref": "#/definitions/errorMessage"
            }
          }
        }
      }
    },
    "/verify": {
      "post": {
        "tags": [
          "account"
        ],
        "operationId": "verify",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/verifyRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "token ready",
            "schema": {
              "$ref": "#/definitions/verifyResponse"
            }
          },
          "404": {
            "description": "token not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/version": {
      "get": {
        "tags": [
          "metadata"
        ],
        "operationId": "version",
        "responses": {
          "200": {
            "description": "current server version",
            "schema": {
              "$ref": "#/definitions/version"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "accessRequest": {
      "type": "object",
      "properties": {
        "envZId": {
          "type": "string"
        },
        "svcToken": {
          "type": "string"
        }
      }
    },
    "accessResponse": {
      "type": "object",
      "properties": {
        "frontendToken": {
          "type": "string"
        }
      }
    },
    "authUser": {
      "type": "object",
      "properties": {
        "password": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "disableRequest": {
      "type": "object",
      "properties": {
        "identity": {
          "type": "string"
        }
      }
    },
    "enableRequest": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string"
        },
        "host": {
          "type": "string"
        }
      }
    },
    "enableResponse": {
      "type": "object",
      "properties": {
        "cfg": {
          "type": "string"
        },
        "identity": {
          "type": "string"
        }
      }
    },
    "environment": {
      "type": "object",
      "properties": {
        "active": {
          "type": "boolean"
        },
        "address": {
          "type": "string"
        },
        "createdAt": {
          "type": "integer"
        },
        "description": {
          "type": "string"
        },
        "host": {
          "type": "string"
        },
        "updatedAt": {
          "type": "integer"
        },
        "zId": {
          "type": "string"
        }
      }
    },
    "environmentServices": {
      "type": "object",
      "properties": {
        "environment": {
          "$ref": "#/definitions/environment"
        },
        "services": {
          "$ref": "#/definitions/services"
        }
      }
    },
    "environmentServicesList": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/environmentServices"
      }
    },
    "environments": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/environment"
      }
    },
    "errorMessage": {
      "type": "string"
    },
    "inviteRequest": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        }
      }
    },
    "loginRequest": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "loginResponse": {
      "type": "string"
    },
    "principal": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "registerRequest": {
      "type": "object",
      "properties": {
        "password": {
          "type": "string"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "registerResponse": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "service": {
      "type": "object",
      "properties": {
        "backendProxyEndpoint": {
          "type": "string"
        },
        "createdAt": {
          "type": "integer"
        },
        "frontendEndpoint": {
          "type": "string"
        },
        "metrics": {
          "$ref": "#/definitions/serviceMetrics"
        },
        "token": {
          "type": "string"
        },
        "updatedAt": {
          "type": "integer"
        },
        "zId": {
          "type": "string"
        }
      }
    },
    "service03": {
      "type": "object",
      "properties": {
        "backendMode": {
          "type": "string"
        },
        "backendProxyEndpoint": {
          "type": "string"
        },
        "frontendEndpoint": {
          "type": "string"
        },
        "frontendSelection": {
          "type": "string"
        },
        "reserved": {
          "type": "boolean"
        },
        "shareMode": {
          "type": "string"
        },
        "token": {
          "type": "string"
        },
        "zId": {
          "type": "string"
        }
      }
    },
    "serviceMetrics": {
      "type": "array",
      "items": {
        "type": "integer"
      }
    },
    "serviceRequest": {
      "type": "object",
      "properties": {
        "envZId": {
          "type": "string"
        },
        "svcToken": {
          "type": "string"
        }
      }
    },
    "services": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/service"
      }
    },
    "shareRequest": {
      "type": "object",
      "properties": {
        "authScheme": {
          "type": "string"
        },
        "authUsers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/authUser"
          }
        },
        "backendMode": {
          "type": "string",
          "enum": [
            "proxy",
            "web",
            "dav"
          ]
        },
        "backendProxyEndpoint": {
          "type": "string"
        },
        "envZId": {
          "type": "string"
        },
        "frontendSelection": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "reserved": {
          "type": "boolean"
        },
        "shareMode": {
          "type": "string",
          "enum": [
            "public",
            "private"
          ]
        }
      }
    },
    "shareResponse": {
      "type": "object",
      "properties": {
        "frontendProxyEndpoint": {
          "type": "string"
        },
        "svcToken": {
          "type": "string"
        }
      }
    },
    "unaccessRequest": {
      "type": "object",
      "properties": {
        "envZId": {
          "type": "string"
        },
        "frontendToken": {
          "type": "string"
        },
        "svcToken": {
          "type": "string"
        }
      }
    },
    "unshareRequest": {
      "type": "object",
      "properties": {
        "envZId": {
          "type": "string"
        },
        "reserved": {
          "type": "boolean"
        },
        "svcToken": {
          "type": "string"
        }
      }
    },
    "verifyRequest": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "verifyResponse": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        }
      }
    },
    "version": {
      "type": "string"
    }
  },
  "securityDefinitions": {
    "key": {
      "type": "apiKey",
      "name": "x-token",
      "in": "header"
    }
  }
}`))
}
