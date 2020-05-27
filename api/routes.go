package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CollectRouter return gin router
func CollectRouter(mode, static string) *gin.Engine {
	if len(mode) > 0 {
		gin.SetMode(mode)
	}
	r := gin.Default()
	r.Use(cors.Default())

	r.Static("/", static)
	r.POST("/api/block/file", BlockFile)
	r.POST("/api/block/raw", BlockRaw)
	r.POST("/api/ledger/file", LedgerFile)

	return r
}
