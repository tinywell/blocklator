package block

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric-protos-go/orderer"
	"github.com/hyperledger/fabric-protos-go/orderer/etcdraft"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/pkg/errors"
)

// Configlator 配置信息解析
type Configlator struct {
	config *common.Config
}

// NewConfiglator 返回一个 Configlator
func NewConfiglator(config *common.Config) *Configlator {
	return &Configlator{
		config: config,
	}
}

// ToDesc return ConfigDesc
func (c *Configlator) ToDesc() *ConfigDesc {
	desc := &ConfigDesc{}
	desc.ApplicationOrgs = c.GetApplicationOrgs()
	desc.ConsortiumOrgs = c.GetConsortiumOrgs()
	desc.OrdererOrgs = c.GetOrdererOrgs()
	desc.Consensus = c.GetConsensusInfo()
	desc.Values = c.GetValues()
	return desc
}

// GetValues 解析配置块中的基本参数信息
func (c *Configlator) GetValues() *ConfigValues {
	values := &ConfigValues{}
	pha := &common.HashingAlgorithm{}
	err := proto.Unmarshal(c.config.ChannelGroup.Values[HashingAlgorithmKey].Value, pha)
	if err != nil {
		return nil
	}
	values.HashingAlgorithm = pha.Name

	if ct, ok := c.config.ChannelGroup.Values[ConsortiumKey]; ok {
		pct := &common.Consortium{}
		err = proto.Unmarshal(ct.Value, pct)
		if err != nil {
			return nil
		}
		values.Consortium = pct.Name
	}
	pod := &common.OrdererAddresses{}
	if oa, ok := c.config.ChannelGroup.Values[OrdererAddressesKey]; ok {
		err = proto.Unmarshal(oa.Value, pod)
		if err != nil {
			return nil
		}
	}
	values.OrdererAddresses = pod.Addresses
	cap := &common.Capabilities{}
	err = proto.Unmarshal(c.config.ChannelGroup.Values[CapabilitiesKey].Value, cap)
	if err != nil {
		return nil
	}
	caps := []string{}
	for k := range cap.Capabilities {
		caps = append(caps, k)
	}
	values.Capabilities = caps
	return values
}

// GetConsensusInfo 解析配置块中的共识相关配置信息
func (c *Configlator) GetConsensusInfo() *ConsensusInfo {
	info := &ConsensusInfo{}
	ct := &orderer.ConsensusType{}
	err := proto.Unmarshal(c.config.ChannelGroup.Groups[OrdererGroupKey].Values[ConsensusTypeKey].Value, ct)
	if err != nil {
		return nil
	}
	info.Type = ct.Type
	switch ct.Type {
	case "etcdraft":
		meta := &etcdraft.ConfigMetadata{}
		err = proto.Unmarshal(ct.Metadata, meta)
		if err != nil {
			break
		}
		info.RaftMetadata = meta
	case "kafka":
		brokers := &orderer.KafkaBrokers{}
		err = proto.Unmarshal(c.config.ChannelGroup.Groups[OrdererGroupKey].Values[KafkaBrokersKey].Value, brokers)
		if err != nil {
			break
		}
		info.Borkers = brokers.Brokers
	default:
		break
	}
	bs := &orderer.BatchSize{}
	err = proto.Unmarshal(c.config.ChannelGroup.Groups[OrdererGroupKey].Values[BatchSizeKey].Value, bs)
	if err != nil {
		return nil
	}
	info.MaxMessageCount = bs.MaxMessageCount
	info.PreferredMaxBytes = bs.PreferredMaxBytes
	info.AbsoluteMaxBytes = bs.AbsoluteMaxBytes
	bt := &orderer.BatchTimeout{}
	err = proto.Unmarshal(c.config.ChannelGroup.Groups[OrdererGroupKey].Values[BatchTimeoutKey].Value, bt)
	if err != nil {
		return nil
	}
	info.BatchTimeOut = bt.Timeout
	return info
}

// GetOrdererOrgs 解析配置块中的 orderer 组织
func (c *Configlator) GetOrdererOrgs() []*GroupOrg {
	orgs := []*GroupOrg{}
	if _, ok := c.config.ChannelGroup.Groups[OrdererGroupKey]; ok {
		for _, g := range c.config.ChannelGroup.Groups[OrdererGroupKey].Groups {
			mspvalue := g.Values[MSPKey].Value
			org, err := c.getOrg(mspvalue)
			if err != nil {
				fmt.Println("getorg error:", err.Error())
				continue
			}
			if _, ok := g.Values[EndpointsKey]; ok {
				ep := g.Values[EndpointsKey].Value
				pbep := &common.OrdererAddresses{}
				err = proto.Unmarshal(ep, pbep)
				if err != nil {
					fmt.Println(err)
					continue
				}
				org.Endpoints = append(org.Endpoints, pbep.Addresses...)
			}
			orgs = append(orgs, org)
		}
	}
	return orgs
}

// GetApplicationOrgs 解析配置块中的应用组织
func (c *Configlator) GetApplicationOrgs() []*GroupOrg {
	orgs := []*GroupOrg{}
	if _, ok := c.config.ChannelGroup.Groups[ApplicationGroupKey]; ok {
		for _, g := range c.config.ChannelGroup.Groups[ApplicationGroupKey].Groups {
			mspvalue := g.Values[MSPKey].Value
			org, err := c.getOrg(mspvalue)
			if err != nil {
				continue
			}
			if _, ok := g.Values[AnchorPeersKey]; ok {
				acp := g.Values[AnchorPeersKey].Value
				pbacp := &peer.AnchorPeers{}
				err = proto.Unmarshal(acp, pbacp)
				if err != nil {
					continue
				}

				for _, p := range pbacp.AnchorPeers {
					org.Endpoints = append(org.Endpoints, fmt.Sprintf("%s:%d", p.Host, p.Port))
				}
			}
			orgs = append(orgs, org)
		}
	}
	return orgs
}

// GetConsortiumOrgs 解析系统配置块中的联盟组织
func (c *Configlator) GetConsortiumOrgs() map[string][]*GroupOrg {
	corgs := make(map[string][]*GroupOrg)

	if _, ok := c.config.ChannelGroup.Groups[ConsortiumsGroupKey]; ok {
		for cn, g := range c.config.ChannelGroup.Groups[ConsortiumsGroupKey].Groups {
			orgs := []*GroupOrg{}
			for _, o := range g.Groups {
				mspvalue := o.Values[MSPKey].Value
				org, err := c.getOrg(mspvalue)
				if err != nil {
					continue
				}
				orgs = append(orgs, org)
			}
			corgs[cn] = orgs
		}
	}
	return corgs
}

func (c *Configlator) getOrg(mspvalue []byte) (*GroupOrg, error) {
	mspc := &msp.MSPConfig{}
	err := proto.Unmarshal(mspvalue, mspc)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal MSPConfig")
	}
	org := &GroupOrg{}
	switch mspc.Type {
	case FABRIC:
		mspcfg := &msp.FabricMSPConfig{}
		err = proto.Unmarshal(mspc.Config, mspcfg)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal FabricMSPConfig error")
		}
		org.Type = mspc.Type
		org.Name = mspcfg.Name
		if len(mspcfg.Admins) > 0 {
			org.Admin = string(mspcfg.Admins[0])
		}
		if len(mspcfg.RootCerts) > 0 {
			org.RootCert = string(mspcfg.RootCerts[0])
		}
		if len(mspcfg.TlsRootCerts) > 0 {
			org.TLSRootCert = string(mspcfg.TlsRootCerts[0])
		}
		org.RevocationList = mspcfg.RevocationList
	case IDEMIX:
		mspcfg := &msp.IdemixMSPConfig{}
		err = proto.Unmarshal(mspc.Config, mspcfg)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal IdemixMSPConfig error")
		}
		org.Type = mspc.Type
		org.Name = mspcfg.Name
	default:
		return nil, errors.Errorf("unexpected msp type:%d", mspc.Type)
	}
	return org, nil
}
