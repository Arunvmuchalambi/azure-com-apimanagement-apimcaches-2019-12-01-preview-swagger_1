package main

import (
	"github.com/apimanagementclient/mcp-server/config"
	"github.com/apimanagementclient/mcp-server/models"
	tools_cache "github.com/apimanagementclient/mcp-server/tools/cache"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_cache.CreateCache_listbyserviceTool(cfg),
		tools_cache.CreateCache_getentitytagTool(cfg),
		tools_cache.CreateCache_updateTool(cfg),
		tools_cache.CreateCache_createorupdateTool(cfg),
		tools_cache.CreateCache_deleteTool(cfg),
		tools_cache.CreateCache_getTool(cfg),
	}
}
