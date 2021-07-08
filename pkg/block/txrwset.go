package block

import (
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/ledger/rwset"
	"github.com/hyperledger/fabric-protos-go/ledger/rwset/kvrwset"
)

// FromProtoBytes 解析 TxRwSet
func (txRwSet *TxRwSet) FromProtoBytes(protoBytes []byte) error {
	protoMsg := &rwset.TxReadWriteSet{}
	var err error
	var txRwSetTemp *TxRwSet
	if err = proto.Unmarshal(protoBytes, protoMsg); err != nil {
		return err
	}
	if txRwSetTemp, err = TxRwSetFromProtoMsg(protoMsg); err != nil {
		return err
	}
	txRwSet.NsRwSets = txRwSetTemp.NsRwSets
	return nil
}

// TxRwSetFromProtoMsg ...
func TxRwSetFromProtoMsg(protoMsg *rwset.TxReadWriteSet) (*TxRwSet, error) {
	txRwSet := &TxRwSet{}
	var nsRwSet *NsRwSet
	var err error
	for _, nsRwSetProtoMsg := range protoMsg.NsRwset {
		if nsRwSet, err = nsRwSetFromProtoMsg(nsRwSetProtoMsg); err != nil {
			return nil, err
		}
		if nsRwSet.NameSpace == "lscc" {
			continue
		}
		txRwSet.NsRwSets = append(txRwSet.NsRwSets, nsRwSet)
	}
	return txRwSet, nil
}

func nsRwSetFromProtoMsg(protoMsg *rwset.NsReadWriteSet) (*NsRwSet, error) {
	nsRwSet := &NsRwSet{NameSpace: protoMsg.Namespace, KvRwSet: &kvrwset.KVRWSet{}}
	if err := proto.Unmarshal(protoMsg.Rwset, nsRwSet.KvRwSet); err != nil {
		return nil, err
	}
	return nsRwSet, nil
}
