package api

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tinywell/blocklator/pkg/block"
)

// BlockFile translate block data in file
func BlockFile(c *gin.Context) {
	rfile, err := c.FormFile("block")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	file, err := rfile.Open()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	blockraw, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	blocklator, err := block.NewBlocklator(blockraw)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	desc, err := blocklator.ToDesc()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "ok", "data": desc})
	return
}
