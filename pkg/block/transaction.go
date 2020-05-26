package block

import (
	"encoding/base64"
	"errors"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric-protos-go/peer"
)

// Translator translator for transaction envelope
type Translator struct {
	env      *common.Envelope
	innerEnv *Envelope
}

// NewTranslator return new Translator
func NewTranslator(env *common.Envelope) (*Translator, error) {
	translator := &Translator{
		env: env,
		innerEnv: &Envelope{
			Signature: env.Signature,
		},
	}
	if err := translator.init(); err != nil {
		return nil, err
	}
	return translator, nil
}

func (t *Translator) init() error {
	payload := &common.Payload{}
	err := proto.Unmarshal(t.env.Payload, payload)
	if err != nil {
		return err
	}
	trans := &peer.Transaction{}
	err = proto.Unmarshal(payload.Data, trans)
	if err != nil {
		return err
	}
	if err = t.unHeader(payload.Header); err != nil {
		return err
	}
	if len(trans.Actions) < 1 {
		return errors.New("no transaction in block")
	}
	if err = t.unAction(trans.Actions[0]); err != nil {
		return err
	}
	return nil
}

func (t *Translator) unHeader(header *common.Header) error {
	ch := &common.ChannelHeader{}
	err := proto.Unmarshal(header.ChannelHeader, ch)
	if err != nil {
		return err
	}
	t.innerEnv.Payload.Header.ChannelHeader = ch
	sh := &common.SignatureHeader{}
	err = proto.Unmarshal(header.SignatureHeader, sh)
	if err != nil {
		return err
	}
	t.innerEnv.Payload.Header.SignatureHeader = sh
	return nil
}

func (t *Translator) unAction(action *peer.TransactionAction) error {
	ccp := &peer.ChaincodeActionPayload{}
	err := proto.Unmarshal(action.Payload, ccp)
	if err != nil {
		return err
	}
	t.innerEnv.Payload.Transaction.ChaincodeAction.Endorses = ccp.Action.Endorsements
	pp := &peer.ChaincodeProposalPayload{}
	err = proto.Unmarshal(ccp.ChaincodeProposalPayload, pp)
	if err != nil {
		return err
	}
	ppi := &peer.ChaincodeInvocationSpec{}
	err = proto.Unmarshal(pp.Input, ppi)
	if err != nil {
		return err
	}
	t.innerEnv.Payload.Transaction.ChaincodeAction.Proposal.Input = ppi.ChaincodeSpec

	ccresp := &peer.ProposalResponsePayload{}
	err = proto.Unmarshal(ccp.Action.ProposalResponsePayload, ccresp)
	if err != nil {
		return err
	}
	t.innerEnv.Payload.Transaction.ChaincodeAction.Response.ProposalHash = ccresp.ProposalHash
	cca := &peer.ChaincodeAction{}
	err = proto.Unmarshal(ccresp.Extension, cca)
	if err != nil {
		return err
	}
	t.innerEnv.Payload.Transaction.ChaincodeAction.Response.ChaincodeAction = cca
	return nil
}

// ToDesc transaction block to TranDesc
func (t *Translator) ToDesc() *TranDesc {
	td := &TranDesc{}
	td.Channel = t.innerEnv.Payload.Header.ChannelHeader.ChannelId
	td.TxID = t.innerEnv.Payload.Header.ChannelHeader.TxId
	ts := t.innerEnv.Payload.Header.ChannelHeader.Timestamp
	td.Time = time.Unix(ts.Seconds, int64(ts.Nanos))
	td.Chaincode = t.innerEnv.Payload.Transaction.ChaincodeAction.Proposal.Input.ChaincodeId.Name
	td.Func = string(t.innerEnv.Payload.Transaction.ChaincodeAction.Proposal.Input.Input.Args[0])
	for _, a := range t.innerEnv.Payload.Transaction.ChaincodeAction.Proposal.Input.Input.Args[1:] {
		td.Args = append(td.Args, string(a))
	}
	td.Resp.Status = t.innerEnv.Payload.Transaction.ChaincodeAction.Response.ChaincodeAction.Response.Status
	td.Resp.Message = t.innerEnv.Payload.Transaction.ChaincodeAction.Response.ChaincodeAction.Response.Message
	td.Resp.Data = string(t.innerEnv.Payload.Transaction.ChaincodeAction.Response.ChaincodeAction.Response.Payload)

	seri := t.innerEnv.Payload.Header.SignatureHeader.Creator
	creator := &msp.SerializedIdentity{}
	err := proto.Unmarshal(seri, creator)
	if err != nil {
		//TODO:
	}
	si := &SignInfo{
		MSPID:     creator.Mspid,
		Signature: base64.StdEncoding.EncodeToString(t.innerEnv.Signature),
	}
	si.Cert, err = NewCert(creator.IdBytes)
	if err != nil {
		//TODO:
	}
	td.Signer = si
	for _, e := range t.innerEnv.Payload.Transaction.ChaincodeAction.Endorses {
		er := &msp.SerializedIdentity{}
		err := proto.Unmarshal(e.Endorser, er)
		if err != nil {
			//TODO:
		}
		si := &SignInfo{
			MSPID:     er.Mspid,
			Signature: base64.StdEncoding.EncodeToString(e.Signature),
		}
		cert, err := NewCert(er.IdBytes)
		if err != nil {
			//TODO:
		}
		si.Cert = cert
		td.Endorsers = append(td.Endorsers, si)
	}
	return td
}

// GetTxID return TxID
func (t *Translator) GetTxID() string {
	return t.innerEnv.Payload.Header.ChannelHeader.TxId
}

// GetChannel return channel id
func (t *Translator) GetChannel() string {
	return t.innerEnv.Payload.Header.ChannelHeader.ChannelId
}
