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
	desc := forblock(blockraw)
	if desc == nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": "invalid block data"})
		return
	}
	// res, err := json.Marshal(desc)
	// if err != nil {
	// 	c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "ok", "data": desc})
	return
}

func forblock(data []byte) *block.Desc {
	blockdesc := &block.Desc{}
	blocklator := block.NewBlocklator(data)
	if blocklator == nil {
		return nil
	}
	blockdesc.BlockNum = blocklator.GetBlockNum()
	channel, err := blocklator.GetChannel()
	if err != nil {
		blockdesc.Channel = ""
	}
	blockdesc.Channel = channel
	blockdesc.Hash = blocklator.GetBlockHash()
	blockdesc.PreHash = blocklator.GetBlockPrehash()

	config := blocklator.GetConfig()
	if config == nil {
		blockdesc.Type = block.BlockTypeTrans
		filters, err := blocklator.GetMetaDataTransFilter()
		if err != nil {
			return nil
		}
		trans := blocklator.GetTransactions()
		for i, t := range trans {
			translator, err := block.NewTranslator(t)
			if err != nil {
				continue
			}
			desc := fortransactions(translator)
			desc.Filter = filters[i]
			blockdesc.Transactions = append(blockdesc.Transactions, desc)
		}
	} else {
		cfg := block.NewConfiglator(config)
		cfgdesc := forconfig(cfg)
		blockdesc.Type = block.BlockTypeConfig
		blockdesc.Config = cfgdesc
	}
	return blockdesc
}

func forconfig(configlator *block.Configlator) *block.ConfigDesc {
	return configlator.ToDesc()
}

func fortransactions(translator *block.Translator) *block.TranDesc {
	return translator.ToDesc()
}
