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
            "name": "Lionel Wilson",
            "url": "https://github.com/Lionel-Wilson",
            "email": "Lionel_Wilson@outlook.com"
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
        "/health": {
            "get": {
                "description": "Check the health of the API",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
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
        "/payments": {
            "get": {
                "description": "Retrieves the details of all previously made payments.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Payments"
                ],
                "summary": "Retrieve all payments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.PaymentDetails"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Processes a payment through the payment gateway.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Payments"
                ],
                "summary": "Process a Payment",
                "parameters": [
                    {
                        "description": "A JSON body",
                        "name": "ProccessPaymentRequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.ProcessPaymentRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.ProcessPaymentResponse"
                        }
                    },
                    "402": {
                        "description": "Payment Required",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/payments/{id}": {
            "get": {
                "description": "Retrieves the details of a previously made payment using its identifier.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Payments"
                ],
                "summary": "Retrieve Payment Details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Payment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.PaymentDetails"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.ErrorResponse": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "Validation failed"
                },
                "statusCode": {
                    "type": "integer",
                    "example": 422
                }
            }
        },
        "main.PaymentDetails": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number",
                    "example": 500
                },
                "cardNumber": {
                    "type": "string",
                    "example": "************1111"
                },
                "currencyCode": {
                    "type": "string",
                    "example": "GBP"
                },
                "expiryDate": {
                    "type": "string",
                    "example": "12/29"
                },
                "firstName": {
                    "type": "string",
                    "example": "John"
                },
                "id": {
                    "type": "string",
                    "example": "PAY-1625843728243722000"
                },
                "lastName": {
                    "type": "string",
                    "example": "Doe"
                },
                "status": {
                    "type": "string",
                    "example": "payment_paid"
                },
                "statusCode": {
                    "type": "integer",
                    "example": 10000
                }
            }
        },
        "main.ProcessPaymentRequest": {
            "type": "object",
            "required": [
                "amount",
                "cardNumber",
                "currencyCode",
                "cvv",
                "expiryDate",
                "firstName",
                "lastName"
            ],
            "properties": {
                "amount": {
                    "type": "number",
                    "example": 500
                },
                "cardNumber": {
                    "type": "string",
                    "example": "4111111111111111"
                },
                "currencyCode": {
                    "type": "string",
                    "example": "GBP"
                },
                "cvv": {
                    "type": "string",
                    "example": "123"
                },
                "expiryDate": {
                    "description": "Custom validation to ensure it matches \"MM/YY\" format",
                    "type": "string",
                    "example": "12/29"
                },
                "firstName": {
                    "type": "string",
                    "example": "John"
                },
                "lastName": {
                    "type": "string",
                    "example": "Doe"
                }
            }
        },
        "main.ProcessPaymentResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "PAY-1625843728243722000"
                },
                "responseSummary": {
                    "type": "string",
                    "example": "Approved"
                },
                "status": {
                    "type": "string",
                    "example": "payment_paid"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Checkout.com Payment Gateway API",
	Description:      "An API for processing and retrieving payments made by Lionel Wilson",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
