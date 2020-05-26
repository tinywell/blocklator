package block

import (
	"encoding/hex"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/pkg/errors"
)

// Blocklator for block translate
type Blocklator struct {
	block *common.Block
}

// NewBlocklator return new Blocklator
func NewBlocklator(raw []byte) (*Blocklator, error) {
	cb := &common.Block{}
	err := proto.Unmarshal(raw, cb)
	if err != nil {
		return nil, err
	}
	return &Blocklator{
		block: cb,
	}, nil
}

// GetBlockNum return block num
func (bl *Blocklator) GetBlockNum() uint64 {
	return bl.block.Header.Number
}

// GetBlockHash return block hash
func (bl *Blocklator) GetBlockHash() string {
	hash := bl.block.Header.DataHash
	return strings.ToUpper(hex.EncodeToString(hash))
}

// GetBlockPrehash return block previoous hash
func (bl *Blocklator) GetBlockPrehash() string {
	hash := bl.block.Header.PreviousHash
	return strings.ToUpper(hex.EncodeToString(hash))
}

// GetChannel return channel id
func (bl *Blocklator) GetChannel() (string, error) {
	if len(bl.block.Data.Data) < 1 {
		return "", errors.New("invalid block")
	}
	env := &common.Envelope{}
	err := proto.Unmarshal(bl.block.Data.Data[0], env)
	if err != nil {
		return "", err
	}
	payload := &common.Payload{}
	err = proto.Unmarshal(env.Payload, payload)
	if err != nil {
		return "", err
	}
	ch := &common.ChannelHeader{}
	err = proto.Unmarshal(payload.Header.ChannelHeader, ch)
	if err != nil {
		return "", err
	}
	return ch.ChannelId, nil
}

// GetMetaDataLastConfig get the last config block num
func (bl *Blocklator) GetMetaDataLastConfig() (uint64, error) {
	if len(bl.block.Metadata.Metadata) > int(common.BlockMetadataIndex_LAST_CONFIG) {
		meta := bl.block.Metadata.Metadata[common.BlockMetadataIndex_LAST_CONFIG]
		cmeta := &common.Metadata{}
		err := proto.Unmarshal(meta, cmeta)
		if err != nil {
			return 0, errors.Wrap(err, "unmarshal metadata error")
		}
		index := cmeta.Value
		cindex := &common.LastConfig{}
		err = proto.Unmarshal(index, cindex)
		if err != nil {
			return 0, errors.Wrap(err, "unmarshal lastconfig error")
		}
		return cindex.Index, nil
	}
	return 0, errors.New("no metadata for lastconfig")
}

// GetMetaDataTransFilter 获取交易有效性标志
func (bl *Blocklator) GetMetaDataTransFilter() ([]bool, error) {
	txfilters := []bool{}
	if len(bl.block.Metadata.Metadata) > int(common.BlockMetadataIndex_TRANSACTIONS_FILTER) {
		filter := bl.block.Metadata.Metadata[common.BlockMetadataIndex_TRANSACTIONS_FILTER]
		for _, f := range filter {
			tf := false
			if peer.TxValidationCode(f) == peer.TxValidationCode_VALID {
				tf = true
			}
			txfilters = append(txfilters, tf)
		}
		return txfilters, nil
	}
	return nil, errors.New("no metadata for transaction filter")
}

// GetCommitHash get commit hash of the block
func (bl *Blocklator) GetCommitHash() (string, error) {
	if len(bl.block.Metadata.Metadata) > int(common.BlockMetadataIndex_COMMIT_HASH) {
		meta := &common.Metadata{}
		hash := bl.block.Metadata.Metadata[common.BlockMetadataIndex_COMMIT_HASH]
		err := proto.Unmarshal(hash, meta)
		if err != nil {
			return "", err
		}
		return strings.ToUpper(hex.EncodeToString(meta.Value)), nil
	}
	return "", errors.New("no commit hash,invalid block metadata")
}

// GetConfig get config from block
func (bl *Blocklator) GetConfig() *common.Config {
	if len(bl.block.Data.Data) > 1 || len(bl.block.Data.Data) < 1 {
		return nil
	}
	env := &common.Envelope{}
	err := proto.Unmarshal(bl.block.Data.Data[0], env)
	if err != nil {
		return nil
	}
	payload := &common.Payload{}
	err = proto.Unmarshal(env.Payload, payload)
	if err != nil {
		return nil
	}

	cfgenv := &common.ConfigEnvelope{}
	err = proto.Unmarshal(payload.Data, cfgenv)
	if err != nil {
		return nil
	}
	return cfgenv.Config
}

// GetTransactions get transction envelops from block
func (bl *Blocklator) GetTransactions() []*common.Envelope {
	envs := []*common.Envelope{}
	for _, d := range bl.block.Data.Data {
		env := &common.Envelope{}
		err := proto.Unmarshal(d, env)
		if err != nil {
			continue
		}
		envs = append(envs, env)
	}
	return envs
}

// ToDesc block to Desc
func (bl *Blocklator) ToDesc() (*Desc, error) {
	blockdesc := &Desc{}
	blockdesc.BlockNum = bl.GetBlockNum()
	channel, err := bl.GetChannel()
	if err != nil {
		blockdesc.Channel = ""
	}
	blockdesc.Channel = channel
	blockdesc.Hash = bl.GetBlockHash()
	blockdesc.PreHash = bl.GetBlockPrehash()
	if err != nil {
		return nil, err
	}
	blockdesc.LastConfig, err = bl.GetMetaDataLastConfig()
	if err != nil {
		return nil, err
	}
	config := bl.GetConfig()
	if config == nil {
		blockdesc.Type = BlockTypeTrans
		filters, err := bl.GetMetaDataTransFilter()
		if err != nil {
			return nil, err
		}
		trans := bl.GetTransactions()
		for i, t := range trans {
			translator, err := NewTranslator(t)
			if err != nil {
				continue
			}
			desc := translator.ToDesc()
			desc.Filter = filters[i]
			blockdesc.Transactions = append(blockdesc.Transactions, desc)
		}
		blockdesc.TransCount = len(trans)
		blockdesc.CommitHash, err = bl.GetCommitHash()
	} else {
		cfg := NewConfiglator(config)
		cfgdesc := cfg.ToDesc()
		blockdesc.Type = BlockTypeConfig
		blockdesc.Config = cfgdesc
	}
	return blockdesc, nil
}
