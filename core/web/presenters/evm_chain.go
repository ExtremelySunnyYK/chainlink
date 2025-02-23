package presenters

import (
	"github.com/smartcontractkit/chainlink/v2/core/chains"
)

// EVMChainResource is an EVM chain JSONAPI resource.
type EVMChainResource struct {
	ChainResource
}

// GetName implements the api2go EntityNamer interface
func (r EVMChainResource) GetName() string {
	return "evm_chain"
}

// NewEVMChainResource returns a new EVMChainResource for chain.
func NewEVMChainResource(chain chains.ChainConfig) EVMChainResource {
	return EVMChainResource{ChainResource{
		JAID:    NewJAID(chain.ID),
		Config:  chain.Cfg,
		Enabled: chain.Enabled,
	}}
}

// EVMNodeResource is an EVM node JSONAPI resource.
type EVMNodeResource struct {
	NodeResource
}

// GetName implements the api2go EntityNamer interface
func (r EVMNodeResource) GetName() string {
	return "evm_node"
}

// NewEVMNodeResource returns a new EVMNodeResource for node.
func NewEVMNodeResource(node chains.NodeStatus) EVMNodeResource {
	return EVMNodeResource{NodeResource{
		JAID:    NewJAID(node.Name),
		ChainID: node.ChainID,
		Name:    node.Name,
		State:   node.State,
		Config:  node.Config,
	}}
}
