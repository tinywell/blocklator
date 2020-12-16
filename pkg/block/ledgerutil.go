package block

import (
	"crypto/sha256"
	"encoding/asn1"
	"math/big"

	"github.com/hyperledger/fabric-protos-go/common"
	ledgerutil "github.com/hyperledger/fabric/common/ledger/util"
	"github.com/pkg/errors"
)

// ExtractLeaderBlock extract block from ledger serialize block
func ExtractLeaderBlock(data []byte) (*common.Block, error) {
	block := &common.Block{}
	var err error
	buf := ledgerutil.NewBuffer(data)
	if block.Header, err = extractHeader(buf); err != nil {
		return nil, errors.WithMessage(err, "extract header error")
	}
	if block.Data, err = extractData(buf); err != nil {
		return nil, errors.WithMessage(err, "extract data error")
	}
	if block.Metadata, err = extractMetadata(buf); err != nil {
		return nil, errors.WithMessage(err, "extract metadata error")
	}
	return block, err
}

func extractHeader(buf *ledgerutil.Buffer) (*common.BlockHeader, error) {
	header := &common.BlockHeader{}
	var err error
	if header.Number, err = buf.DecodeVarint(); err != nil {
		return nil, errors.Wrap(err, "error decoding the block number")
	}
	if header.DataHash, err = buf.DecodeRawBytes(false); err != nil {
		return nil, errors.Wrap(err, "error decoding the data hash")
	}
	if header.PreviousHash, err = buf.DecodeRawBytes(false); err != nil {
		return nil, errors.Wrap(err, "error decoding the previous hash")
	}
	if len(header.PreviousHash) == 0 {
		header.PreviousHash = nil
	}
	return header, nil
}

func extractData(buf *ledgerutil.Buffer) (*common.BlockData, error) {
	data := &common.BlockData{}
	var numItems uint64
	var err error

	if numItems, err = buf.DecodeVarint(); err != nil {
		return nil, errors.Wrap(err, "error decoding the length of block data")
	}
	for i := uint64(0); i < numItems; i++ {
		var txEnvBytes []byte

		if txEnvBytes, err = buf.DecodeRawBytes(false); err != nil {
			return nil, errors.Wrap(err, "error decoding the transaction enevelope")
		}

		data.Data = append(data.Data, txEnvBytes)
	}
	return data, nil
}

func extractMetadata(buf *ledgerutil.Buffer) (*common.BlockMetadata, error) {
	metadata := &common.BlockMetadata{}
	var numItems uint64
	var metadataEntry []byte
	var err error
	if numItems, err = buf.DecodeVarint(); err != nil {
		return nil, errors.Wrap(err, "error decoding the length of block metadata")
	}
	for i := uint64(0); i < numItems; i++ {
		if metadataEntry, err = buf.DecodeRawBytes(false); err != nil {
			return nil, errors.Wrap(err, "error decoding the block metadata")
		}
		metadata.Metadata = append(metadata.Metadata, metadataEntry)
	}
	return metadata, nil
}

type asn1Header struct {
	Number       *big.Int
	PreviousHash []byte
	DataHash     []byte
}

// HeaderBytes ...
func HeaderBytes(b *common.BlockHeader) []byte {
	asn1Header := asn1Header{
		PreviousHash: b.PreviousHash,
		DataHash:     b.DataHash,
		Number:       new(big.Int).SetUint64(b.Number),
	}
	result, err := asn1.Marshal(asn1Header)
	if err != nil {
		// Errors should only arise for types which cannot be encoded, since the
		// BlockHeader type is known a-priori to contain only encodable types, an
		// error here is fatal and should not be propogated
		panic(err)
	}
	return result
}

// HeaderHash ...
func HeaderHash(b *common.BlockHeader) []byte {
	sum := sha256.Sum256(HeaderBytes(b))
	return sum[:]
}
