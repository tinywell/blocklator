package block

import (
	"time"

	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/orderer/etcdraft"
	"github.com/hyperledger/fabric-protos-go/peer"
)

// block type
const (
	BlockTypeTrans = iota
	BlockTypeConfig
)

// GroupOrg 配置块中的组织信息
type GroupOrg struct {
	Type           int32    `json:"type,omitempty" db:"type"`
	Name           string   `json:"name,omitempty" db:"name"`
	RootCert       string   `json:"root_cert,omitempty" db:"root_cert"`
	TLSRootCert    string   `json:"tls_root_cert,omitempty" db:"tls_root_cert"`
	Admin          string   `json:"admin,omitempty" db:"admin"`
	RevocationList [][]byte `json:"revocation_list,omitempty" db:"revocation_list"`
	Endpoints      []string `json:"endpoints,omitempty" db:"endpoints"` // peer: anchorpeers  orderer: ordereraddress
}

// ConsensusInfo 配置快中的共识信息
type ConsensusInfo struct {
	Type              string                   `json:"type,omitempty" db:"type"`
	RaftMetadata      *etcdraft.ConfigMetadata `json:"raft_metadata,omitempty" db:"raft_metadata"`
	MaxMessageCount   uint32                   `json:"max_message_count,omitempty" db:"max_message_count"`
	AbsoluteMaxBytes  uint32                   `json:"absolute_max_bytes,omitempty" db:"absolute_max_bytes"`
	PreferredMaxBytes uint32                   `json:"preferred_max_bytes,omitempty" db:"preferred_max_bytes"`
	BatchTimeOut      string                   `json:"batch_time_out,omitempty" db:"batch_time_out"`
	Borkers           []string                 `json:"borkers,omitempty" db:"borkers"`
}

// ConfigValues config values
type ConfigValues struct {
	Consortium            string   `json:"consortium,omitempty" db:"consortium"`
	HashingAlgorithm      string   `json:"hashing_algorithm,omitempty" db:"hashing_algorithm"`
	OrdererAddresses      []string `json:"orderer_addresses,omitempty" db:"orderer_addresses"`
	BlockDataHashingWidth int      `json:"block_data_hashing_width,omitempty" db:"block_data_hashing_width"`
}

// Desc block description
type Desc struct {
	BlockNum     uint64      `json:"block_num" db:"block_num"`
	Hash         string      `json:"hash" db:"hash"`
	PreHash      string      `json:"pre_hash" db:"pre_hash"`
	Channel      string      `json:"channel" db:"channel"`
	Type         int         `json:"type" db:"type"` // 0: transaction 1: config
	Config       *ConfigDesc `json:"config" db:"config"`
	Transactions []*TranDesc `json:"transactions" db:"transactions"`
	TransCount   int         `json:"trans_count" db:"trans_count"`
	CommitHash   string      `json:"commit_hash" db:"commit_hash"`
	LastConfig   uint64      `json:"last_config" db:"last_config"`
}

// ConfigDesc config description
type ConfigDesc struct {
	OrdererOrgs     []*GroupOrg            `json:"orderer_orgs" db:"orderer_orgs"`
	ConsortiumOrgs  map[string][]*GroupOrg `json:"consortium_orgs" db:"consortium_orgs"`
	ApplicationOrgs []*GroupOrg            `json:"application_orgs" db:"application_orgs"`
	Values          *ConfigValues          `json:"values" db:"values"`
	Consensus       *ConsensusInfo         `json:"consensus" db:"consensus"`
}

// SignInfo sign info
type SignInfo struct {
	MSPID     string `json:"mspid,omitempty" db:"mspid"`
	Cert      *Cert  `json:"cert,omitempty" db:"cert"`
	Signature string `json:"signature,omitempty" db:"signature"`
}

// TranDesc transaction description
type TranDesc struct {
	Channel   string    `json:"channel," db:"channel"`
	TxID      string    `json:"tx_id," db:"tx_id"`
	Time      time.Time `json:"time," db:"time"`
	Chaincode string    `json:"chaincode," db:"chaincode"`
	Func      string    `json:"func," db:"func"`
	Args      []string  `json:"args," db:"args"`
	Resp      struct {
		Status  int32  `json:"status," db:"status"`
		Message string `json:"message," db:"message"`
		Data    string `json:"data," db:"data"`
	} `json:"resp," db:"resp"`
	Filter    bool        `json:"filter," db:"filter"`
	Signer    *SignInfo   `json:"signer," db:"signer"`
	Endorsers []*SignInfo `json:"endorsers," db:"endorsers"`
}

// Envelope clean struct for envelope
type Envelope struct {
	Payload struct {
		Header struct {
			ChannelHeader   *common.ChannelHeader
			SignatureHeader *common.SignatureHeader
		}
		Transaction struct {
			Header          *common.SignatureHeader
			ChaincodeAction struct {
				Proposal struct {
					Input *peer.ChaincodeSpec
				}
				Response struct {
					ProposalHash    []byte
					ChaincodeAction *peer.ChaincodeAction
				}
				Endorses []*peer.Endorsement
			}
		}
	}
	Signature []byte
}
