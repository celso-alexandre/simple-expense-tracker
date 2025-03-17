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
        "/expense-plan-record/list": {
            "post": {
                "description": "List all expense-plan-record items (using cursor-based pagination)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ExpensePlanRecord"
                ],
                "summary": "List all expense-plan-record items",
                "parameters": [
                    {
                        "description": "ListExpensePlanRecordRequest",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/handler.ListExpensePlanRecordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ListExpensePlanRecordResponse"
                        }
                    }
                }
            }
        },
        "/expense-plan/create": {
            "post": {
                "description": "Create expense-plan item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ExpensePlan"
                ],
                "summary": "Create expense-plan item",
                "parameters": [
                    {
                        "description": "CreateExpensePlanRequest",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/handler.CreateExpensePlanRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.CreateExpensePlanResponse"
                        }
                    }
                }
            }
        },
        "/expense-plan/delete": {
            "post": {
                "description": "Delete expense-plan item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ExpensePlan"
                ],
                "summary": "Delete expense-plan item",
                "parameters": [
                    {
                        "description": "DeleteExpensePlanRequest",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/handler.DeleteExpensePlanRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.DeleteExpensePlanResponse"
                        }
                    }
                }
            }
        },
        "/expense-plan/get": {
            "post": {
                "description": "Get all expense-plan items (using cursor-based pagination)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ExpensePlan"
                ],
                "summary": "Get all expense-plan items",
                "parameters": [
                    {
                        "description": "GetExpensePlanRequest",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/handler.GetExpensePlanRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.GetExpensePlanResponse"
                        }
                    }
                }
            }
        },
        "/expense-plan/list": {
            "post": {
                "description": "List all expense-plan items (using cursor-based pagination)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ExpensePlan"
                ],
                "summary": "List all expense-plan items",
                "parameters": [
                    {
                        "description": "ListExpensePlanRequest",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/handler.ListExpensePlanRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ListExpensePlanResponse"
                        }
                    }
                }
            }
        },
        "/expense-plan/update": {
            "post": {
                "description": "Update expense-plan item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ExpensePlan"
                ],
                "summary": "Update expense-plan item",
                "parameters": [
                    {
                        "description": "UpdateExpensePlanRequest",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/handler.UpdateExpensePlanRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.UpdateExpensePlanResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.CreateExpensePlanRequest": {
            "type": "object",
            "required": [
                "amount_planned",
                "category",
                "title"
            ],
            "properties": {
                "amount_planned": {
                    "type": "integer"
                },
                "category": {
                    "$ref": "#/definitions/query.ExpensePlanCategory"
                },
                "recurrency_type": {
                    "$ref": "#/definitions/query.RecurrencyType"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "handler.CreateExpensePlanResponse": {
            "type": "object",
            "properties": {
                "expense_plan_id": {
                    "type": "integer"
                }
            }
        },
        "handler.DeleteExpensePlanRequest": {
            "type": "object",
            "required": [
                "expense_plan_id"
            ],
            "properties": {
                "expense_plan_id": {
                    "type": "integer"
                }
            }
        },
        "handler.DeleteExpensePlanResponse": {
            "type": "object"
        },
        "handler.GetExpensePlanRequest": {
            "type": "object",
            "required": [
                "expense_plan_id"
            ],
            "properties": {
                "expense_plan_id": {
                    "type": "integer"
                }
            }
        },
        "handler.GetExpensePlanResponse": {
            "type": "object",
            "properties": {
                "amount_planned": {
                    "type": "integer"
                },
                "category": {
                    "$ref": "#/definitions/query.ExpensePlanCategory"
                },
                "created_at": {
                    "type": "string"
                },
                "expense_plan_id": {
                    "type": "integer"
                },
                "first_payment_date": {
                    "type": "string"
                },
                "last_amount_spent": {
                    "type": "integer"
                },
                "last_paid_date": {
                    "type": "string"
                },
                "last_payment_date": {
                    "type": "string"
                },
                "paid_count": {
                    "type": "integer"
                },
                "recurrency_interval": {
                    "type": "integer"
                },
                "recurrency_type": {
                    "$ref": "#/definitions/query.RecurrencyType"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "handler.ListExpensePlanRecordRequest": {
            "type": "object"
        },
        "handler.ListExpensePlanRecordResponse": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handler.ListExpensePlanRecordResponse_ListExpensePlanRecord"
                    }
                }
            }
        },
        "handler.ListExpensePlanRecordResponse_ListExpensePlan": {
            "type": "object",
            "properties": {
                "amount_planned": {
                    "type": "integer"
                },
                "category": {
                    "type": "string"
                },
                "expense_plan_id": {
                    "type": "integer"
                },
                "recurrency_interval": {
                    "type": "integer"
                },
                "recurrency_type": {
                    "$ref": "#/definitions/query.RecurrencyType"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "handler.ListExpensePlanRecordResponse_ListExpensePlanRecord": {
            "type": "object",
            "properties": {
                "amount_paid": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "expense_plan": {
                    "$ref": "#/definitions/handler.ListExpensePlanRecordResponse_ListExpensePlan"
                },
                "expense_plan_id": {
                    "type": "integer"
                },
                "expense_plan_record_id": {
                    "type": "integer"
                },
                "expense_plan_sequence": {
                    "type": "integer"
                },
                "paid_date": {
                    "type": "string"
                },
                "payment_date": {
                    "type": "string"
                },
                "previous_payment_amount": {
                    "type": "integer"
                },
                "previous_payment_date": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "handler.ListExpensePlanRequest": {
            "type": "object"
        },
        "handler.ListExpensePlanResponse": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handler.ListExpensePlanResponse_ListExpensePlan"
                    }
                }
            }
        },
        "handler.ListExpensePlanResponse_ListExpensePlan": {
            "type": "object",
            "properties": {
                "amount_planned": {
                    "type": "integer"
                },
                "category": {
                    "$ref": "#/definitions/query.ExpensePlanCategory"
                },
                "created_at": {
                    "type": "string"
                },
                "expense_plan_id": {
                    "type": "integer"
                },
                "first_payment_date": {
                    "type": "string"
                },
                "last_amount_spent": {
                    "type": "integer"
                },
                "last_paid_date": {
                    "type": "string"
                },
                "last_payment_date": {
                    "type": "string"
                },
                "paid_count": {
                    "type": "integer"
                },
                "recurrency_interval": {
                    "type": "integer"
                },
                "recurrency_type": {
                    "$ref": "#/definitions/query.RecurrencyType"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "handler.UpdateExpensePlanRequest": {
            "type": "object",
            "required": [
                "amount_planned",
                "category",
                "expense_plan_id",
                "title"
            ],
            "properties": {
                "amount_planned": {
                    "type": "integer"
                },
                "category": {
                    "$ref": "#/definitions/query.ExpensePlanCategory"
                },
                "expense_plan_id": {
                    "type": "integer"
                },
                "recurrency_type": {
                    "$ref": "#/definitions/query.RecurrencyType"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "handler.UpdateExpensePlanResponse": {
            "type": "object",
            "properties": {
                "expense_plan_id": {
                    "type": "integer"
                }
            }
        },
        "query.ExpensePlanCategory": {
            "type": "string",
            "enum": [
                "FOOD",
                "TRANSPORT",
                "PROPERTY",
                "TAX",
                "ENTERTAINMENT",
                "OTHER"
            ],
            "x-enum-varnames": [
                "ExpensePlanCategoryFOOD",
                "ExpensePlanCategoryTRANSPORT",
                "ExpensePlanCategoryPROPERTY",
                "ExpensePlanCategoryTAX",
                "ExpensePlanCategoryENTERTAINMENT",
                "ExpensePlanCategoryOTHER"
            ]
        },
        "query.RecurrencyType": {
            "type": "string",
            "enum": [
                "MONTHLY",
                "YEARLY"
            ],
            "x-enum-varnames": [
                "RecurrencyTypeMONTHLY",
                "RecurrencyTypeYEARLY"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "My API",
	Description:      "This is a sample API using Swagger.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
