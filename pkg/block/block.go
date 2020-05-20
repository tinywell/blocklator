package block

import (
	"encoding/base64"
	"errors"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/common"
)


// Blocklator for block translate
type Blocklator struct {
	block *common.Block
}

// NewBlocklator return new Blocklator
func NewBlocklator(raw []byte) *Blocklator {
	cb := &common.Block{}
	err:= proto.Unmarshal(raw,cb)
	if err != nil {
		return nil
	}
	return &Blocklator{
		block:cb,
	}
}

// GetBlockNum return block num
func (bl *Blocklator) GetBlockNum() uint64  {
	return bl.block.Header.Number
}

// GetBlockHash return block hash
func (bl *Blocklator) GetBlockHash() string  {
	hash := bl.block.Header.DataHash	
	return base64.StdEncoding.EncodeToString(hash)	
}

// GetBlockPrehash return block previoous hash
func (bl *Blocklator) GetBlockPrehash() string  {
	hash := bl.block.Header.PreviousHash	
	return base64.StdEncoding.EncodeToString(hash)		
}


// GetChannel return channel id
func (bl *Blocklator) GetChannel()  (string,error) {
	if len(bl.block.Data.Data) < 1{
		return "",errors.New("invalid block")
	}
	env := &common.Envelope{}
	err:= proto.Unmarshal(bl.block.Data.Data[0],env)
	if err != nil {
		return "",err
	}
	payload := &common.Payload{}
	err = proto.Unmarshal(env.Payload,payload)
	if err != nil {
		return "",err
	}
	ch := &common.ChannelHeader{}
	err=proto.Unmarshal(payload.Header.ChannelHeader,ch)
	if err != nil {
		return "", err
	}
	return ch.ChannelId,nil	
}

// GetConfig get config from block
func (bl *Blocklator) GetConfig() *common.Config {
	if len(bl.block.Data.Data) > 1 || len(bl.block.Data.Data) < 1{
		return nil
	}
	env := &common.Envelope{}
	err:= proto.Unmarshal(bl.block.Data.Data[0],env)
	if err != nil {
		return nil
	}
	payload := &common.Payload{}
	err = proto.Unmarshal(env.Payload,payload)
	if err != nil {
		return nil
	}

	cfgenv := &common.ConfigEnvelope{}
	err=proto.Unmarshal(payload.Data,cfgenv)
	if err != nil {
		return nil
	}
	return cfgenv.Config
}

func (bl *Blocklator) GetTransactions() {
	
}