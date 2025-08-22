package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Result int `json:"result,omitempty"`
	Data map[string]interface{} `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Data map[string]interface{} `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Result int `json:"result,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Data map[string]interface{} `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Result int `json:"result,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Data map[string]interface{} `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Result int `json:"result,omitempty"`
}

// BalanceResponse represents the BalanceResponse schema from the OpenAPI specification
type BalanceResponse struct {
	Result int `json:"result,omitempty"`
	Data map[string]interface{} `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

// IBANResult represents the IBANResult schema from the OpenAPI specification
type IBANResult struct {
	Message string `json:"message,omitempty"`
	Result int `json:"result,omitempty"`
	Validations []interface{} `json:"validations,omitempty"`
	Data map[string]interface{} `json:"data,omitempty"`
	Expremental int `json:"expremental,omitempty"`
}

// IBANResultBasic represents the IBANResultBasic schema from the OpenAPI specification
type IBANResultBasic struct {
	Message string `json:"message,omitempty"`
	Result int `json:"result,omitempty"`
	Validations []interface{} `json:"validations,omitempty"`
	Data map[string]interface{} `json:"data,omitempty"`
	Expremental int `json:"expremental,omitempty"`
}
