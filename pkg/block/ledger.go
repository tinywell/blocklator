package block

import (
	"bufio"
	"io"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

const (
	peekBytes = 8
)

// Ledgerlator translator for ledger data
type Ledgerlator struct {
	ledger *bufio.Reader
}

// NewLedgerlator return new ledgerlator from a Reader
func NewLedgerlator(ledger *bufio.Reader) *Ledgerlator {
	return &Ledgerlator{
		ledger: ledger,
	}
}

// RetriveBlocks retrive blocks from ledger
func (l Ledgerlator) RetriveBlocks() ([][]byte, error) {
	blocks := [][]byte{}
	for {

		peek, err := l.ledger.Peek(peekBytes)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, errors.Wrap(err, "read peek info error")
		}
		length, n := proto.DecodeVarint(peek)
		if n == 0 {
			return nil, errors.New("peek decode not enough")
		}

		l.ledger.Discard(n)
		block := make([]byte, length)
		_, err = io.ReadAtLeast(l.ledger, block, int(length))
		if err != nil {
			return nil, errors.Wrap(err, "read block bytes error")
		}
		blocks = append(blocks, block)
	}
	return blocks, nil
}
