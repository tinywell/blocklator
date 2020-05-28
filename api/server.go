package api

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/tinywell/blocklator/pkg/block"
	"github.com/tinywell/blocklator/pkg/store"
)

// Pagination params
const (
	PageSize = 20
)

var (
	// Cache cache
	Cache store.Cache
)

func init() {
	Cache = store.NewCache()
}

// Ledger ledger data
type Ledger struct {
	Pagination bool          `json:"pagination" db:"pagination"`
	Key        string        `json:"key,omitempty" db:"key"`
	Blocks     []*block.Desc `json:"blocks" db:"blocks"`
	CurPage    int           `json:"cur_page" db:"cur_page"`
	Total      int           `json:"total" db:"total"`
}

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
	ledgerRsp := &Ledger{
		Total:  len(blockSums),
		Blocks: blockSums,
	}
	data, err := json.Marshal(blockSums)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": errors.WithMessage(err, "marsh block summary error").Error()})
		return
	}
	if len(data) > 4*1024*1024 || len(blockSums) > PageSize {
		key, err := Cache.KeyGen(data)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": errors.WithMessage(err, "generate block summar keyy error").Error()})
			return
		}
		Cache.Store(key, blockSums)
		ledgerRsp.Key = key
		ledgerRsp.CurPage = 1
		ledgerRsp.Pagination = true
		ledgerRsp.Blocks = blockSums[0:PageSize]
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "ok", "data": ledgerRsp})
	return
}

// LedgerBlocks get blocks by pagination
func LedgerBlocks(c *gin.Context) {
	key := c.Query("key")
	page := c.Query("page")
	if len(key) == 0 || len(page) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": errors.New("paaination query need key and page").Error()})
		return
	}
	cp, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": errors.WithMessage(err, "decode page param error").Error()})
		return
	}
	blocks, err := Cache.Load(key)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": errors.WithMessage(err, "get data cache error").Error()})
		return
	}
	if blockSum, ok := blocks.([]*block.Desc); ok {
		ledgerRsp := Ledger{
			Pagination: true,
			CurPage:    cp,
			Key:        key,
			Total:      len(blockSum),
		}
		start := (cp - 1) * PageSize
		end := cp * PageSize
		if end > len(blockSum) {
			end = len(blockSum)
		}
		ledgerRsp.Blocks = blockSum[start:end]
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "ok", "data": ledgerRsp})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 500, "msg": "data cache  invalid ,please reload your ledger file"})
	return
}
