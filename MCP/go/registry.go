package main

import (
	"github.com/ibanapi-openapi-documentation/mcp-server/config"
	"github.com/ibanapi-openapi-documentation/mcp-server/models"
	tools_ibanapi "github.com/ibanapi-openapi-documentation/mcp-server/tools/ibanapi"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_ibanapi.CreateGetbalanceTool(cfg),
		tools_ibanapi.CreateValidateibanTool(cfg),
		tools_ibanapi.CreateValidateibanbasicTool(cfg),
	}
}
