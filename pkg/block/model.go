package block

import (
	"github.com/hyperledger/fabric-protos-go/orderer/etcdraft"
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
	RootCerts      [][]byte `json:"root_certs,omitempty" db:"root_certs"`
	TLSRootCerts   [][]byte `json:"tls_root_certs,omitempty" db:"tls_root_certs"`
	Admins         [][]byte `json:"admins,omitempty" db:"admins"`
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
	BlockNum     uint64      `json:"block_num,omitempty" db:"block_num"`
	Hash         string      `json:"hash,omitempty" db:"hash"`
	PreHash      string      `json:"pre_hash,omitempty" db:"pre_hash"`
	Channel      string      `json:"channel,omitempty" db:"channel"`
	Type         int         `json:"type,omitempty" db:"type"` // 0: transaction 1: config
	Config       *ConfigDesc `json:"config,omitempty" db:"config"`
	Transactions []*TranDesc `json:"transactions,omitempty" db:"transactions"`
}

// ConfigDesc config description
type ConfigDesc struct {
	OrdererOrgs     []*GroupOrg            `json:"orderer_orgs,omitempty" db:"orderer_orgs"`
	ConsortiumOrgs  map[string][]*GroupOrg `json:"consortium_orgs,omitempty" db:"consortium_orgs"`
	ApplicationOrgs []*GroupOrg            `json:"application_orgs,omitempty" db:"application_orgs"`
	Values          *ConfigValues          `json:"values,omitempty" db:"values"`
	Consensus       *ConsensusInfo         `json:"consensus,omitempty" db:"consensus"`
}

// TranDesc transaction description
type TranDesc struct {
}
