package api

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

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

// BlockRaw translate block data in base64 format
func BlockRaw(c *gin.Context) {
	data := c.PostForm("block")
	blockraw, err := base64.StdEncoding.DecodeString(data)
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

// LedgerFile retrive ledger file into blocks
func LedgerFile(c *gin.Context) {
	lfile, err := c.FormFile("ledger")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	ledger, err := lfile.Open()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	ledgerLator := block.NewLedgerlator(bufio.NewReader(ledger))
	blocks, err := ledgerLator.RetriveBlocks()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
		return
	}

	// blockStrs := []string{}
	blockSums := []*block.Desc{}
	for _, b := range blocks {
		// blockStrs = append(blockStrs, base64.StdEncoding.EncodeToString(b))
		bl, err := block.NewBlocklatorFromLedgerRaw(b)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": errors.WithMessage(err, "create blocklator error").Error()})
			return
		}
		sum, err := bl.ToDesc()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": errors.WithMessage(err, "get block summary").Error()})
			return
		}
		blockSums = append(blockSums, sum)
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "ok", "data": blockSums})
	return
}
