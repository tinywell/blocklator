package api

import (
	"github.com/gin-contrib/cors"
	ginstatic "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// CollectRouter return gin router
func CollectRouter(mode, static string) *gin.Engine {
	if len(mode) > 0 {
		gin.SetMode(mode)
	}
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(ginstatic.Serve("/", ginstatic.LocalFile(static, false)))
	// r.Static("/front", static)
	r.POST("/api/block/file", BlockFile)
	r.POST("/api/block/raw", BlockRaw)
	r.POST("/api/ledger/file", LedgerFile)
	r.GET("/api/ledger/blocks", LedgerBlocks)

	return r
}
