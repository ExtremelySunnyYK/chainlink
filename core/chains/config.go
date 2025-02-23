package chains

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

type ChainConfigs[I ID] interface {
	Chains(offset, limit int, ids ...I) ([]ChainConfig, int, error)
}

type NodeConfigs[I ID, N Node] interface {
	Node(name string) (N, error)
	Nodes(chainID I) (nodes []N, err error)

	NodeStatus(name string) (NodeStatus, error)
	NodeStatusesPaged(offset, limit int, chainIDs ...string) (nodes []NodeStatus, count int, err error)
}

// Configs holds chain and node configurations.
type Configs[I ID, N Node] interface {
	ChainConfigs[I]
	NodeConfigs[I, N]
}

type ChainConfig struct {
	ID      string
	Enabled bool
	Cfg     string // TOML
}

type NodeStatus struct {
	ChainID string
	Name    string
	Config  string // TOML
	State   string
}

func EnsureChains[I ID](q pg.Q, prefix string, ids []I) (err error) {
	named := make([]struct{ ID I }, len(ids))
	for i, id := range ids {
		named[i].ID = id
	}
	sql := fmt.Sprintf("INSERT INTO %s_chains (id, created_at, updated_at) VALUES (:id, NOW(), NOW()) ON CONFLICT DO NOTHING;", prefix)

	if _, err := q.NamedExec(sql, named); err != nil {
		return errors.Wrapf(err, "failed to insert chains %v", ids)
	}
	return nil
}
