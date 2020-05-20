package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CollectRouter return gin router
func CollectRouter(mode string) *gin.Engine {
	if len(mode) > 0 {
		gin.SetMode(mode)
	}
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/api/block/file", BlockFile)

	return r
}
