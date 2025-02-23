package v2

import (
	"fmt"
	"net/url"

	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/pelletier/go-toml/v2"
	"github.com/shopspring/decimal"
	"go.uber.org/multierr"
	"golang.org/x/exp/slices"
	"gopkg.in/guregu/null.v4"

	"github.com/smartcontractkit/chainlink/v2/core/assets"
	"github.com/smartcontractkit/chainlink/v2/core/chains"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/config"
	v2 "github.com/smartcontractkit/chainlink/v2/core/config/v2"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ethkey"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

type HasEVMConfigs interface {
	EVMConfigs() EVMConfigs
}

type EVMConfigs []*EVMConfig

func (cs EVMConfigs) ValidateConfig() (err error) {
	// Unique chain IDs
	chainIDs := v2.UniqueStrings{}
	for i, c := range cs {
		if chainIDs.IsDupeFmt(c.ChainID) {
			err = multierr.Append(err, v2.NewErrDuplicate(fmt.Sprintf("%d.ChainID", i), c.ChainID.String()))
		}
	}

	// Unique node names
	names := v2.UniqueStrings{}
	for i, c := range cs {
		for j, n := range c.Nodes {
			if names.IsDupe(n.Name) {
				err = multierr.Append(err, v2.NewErrDuplicate(fmt.Sprintf("%d.Nodes.%d.Name", i, j), *n.Name))
			}
		}
	}

	// Unique node WSURLs
	wsURLs := v2.UniqueStrings{}
	for i, c := range cs {
		for j, n := range c.Nodes {
			u := (*url.URL)(n.WSURL)
			if wsURLs.IsDupeFmt(u) {
				err = multierr.Append(err, v2.NewErrDuplicate(fmt.Sprintf("%d.Nodes.%d.WSURL", i, j), u.String()))
			}
		}
	}

	// Unique node HTTPURLs
	httpURLs := v2.UniqueStrings{}
	for i, c := range cs {
		for j, n := range c.Nodes {
			u := (*url.URL)(n.HTTPURL)
			if httpURLs.IsDupeFmt(u) {
				err = multierr.Append(err, v2.NewErrDuplicate(fmt.Sprintf("%d.Nodes.%d.HTTPURL", i, j), u.String()))
			}
		}
	}
	return
}

func (cs *EVMConfigs) SetFrom(fs *EVMConfigs) {
	for _, f := range *fs {
		if f.ChainID == nil {
			*cs = append(*cs, f)
		} else if i := slices.IndexFunc(*cs, func(c *EVMConfig) bool {
			return c.ChainID != nil && c.ChainID.Cmp(f.ChainID) == 0
		}); i == -1 {
			*cs = append(*cs, f)
		} else {
			(*cs)[i].SetFrom(f)
		}
	}
}

func (cs EVMConfigs) Chains(ids ...utils.Big) (r []chains.ChainConfig, err error) {
	for _, ch := range cs {
		if ch == nil {
			continue
		}
		if len(ids) > 0 {
			var match bool
			for _, id := range ids {
				if id.Cmp(ch.ChainID) == 0 {
					match = true
					break
				}
			}
			if !match {
				continue
			}
		}
		ch2 := chains.ChainConfig{
			ID:      ch.ChainID.String(),
			Enabled: ch.IsEnabled(),
		}
		ch2.Cfg, err = ch.TOMLString()
		if err != nil {
			return
		}
		r = append(r, ch2)
	}
	return
}

func (cs EVMConfigs) Node(name string) (types.Node, error) {
	for i := range cs {
		for _, n := range cs[i].Nodes {
			if n.Name != nil && *n.Name == name {
				return legacyNode(n, cs[i].ChainID), nil
			}
		}
	}
	return types.Node{}, chains.ErrNotFound
}

func (cs EVMConfigs) NodeStatus(name string) (chains.NodeStatus, error) {
	for i := range cs {
		for _, n := range cs[i].Nodes {
			if n.Name != nil && *n.Name == name {
				return nodeStatus(n, cs[i].ChainID.String())
			}
		}
	}
	return chains.NodeStatus{}, chains.ErrNotFound
}

func legacyNode(n *Node, chainID *utils.Big) (v2 types.Node) {
	v2.Name = *n.Name
	v2.EVMChainID = *chainID
	if n.HTTPURL != nil {
		v2.HTTPURL = null.StringFrom(n.HTTPURL.String())
	}
	if n.WSURL != nil {
		v2.WSURL = null.StringFrom(n.WSURL.String())
	}
	if n.SendOnly != nil {
		v2.SendOnly = *n.SendOnly
	}
	return
}

func nodeStatus(n *Node, chainID string) (chains.NodeStatus, error) {
	var s chains.NodeStatus
	s.ChainID = chainID
	s.Name = *n.Name
	b, err := toml.Marshal(n)
	if err != nil {
		return chains.NodeStatus{}, err
	}
	s.Config = string(b)
	return s, nil
}

func (cs EVMConfigs) nodes(chainID string) (ns EVMNodes) {
	for _, c := range cs {
		if c.ChainID.String() == chainID {
			return c.Nodes
		}
	}
	return nil
}

func (cs EVMConfigs) Nodes(chainID utils.Big) (ns []types.Node, err error) {
	id := chainID.String()
	nodes := cs.nodes(id)
	if nodes == nil {
		err = chains.ErrNotFound
		return
	}
	for _, n := range nodes {
		if n == nil {
			continue
		}
		ns = append(ns, legacyNode(n, &chainID))
	}
	return
}

func (cs EVMConfigs) NodeStatuses(chainIDs ...string) (ns []chains.NodeStatus, err error) {
	if len(chainIDs) == 0 {
		for i := range cs {
			for _, n := range cs[i].Nodes {
				if n == nil {
					continue
				}
				n2, err := nodeStatus(n, cs[i].ChainID.String())
				if err != nil {
					return nil, err
				}
				ns = append(ns, n2)
			}
		}
		return
	}
	for _, id := range chainIDs {
		for _, n := range cs.nodes(id) {
			if n == nil {
				continue
			}
			n2, err := nodeStatus(n, id)
			if err != nil {
				return nil, err
			}
			ns = append(ns, n2)
		}
	}
	return
}

type EVMNodes []*Node

func (ns *EVMNodes) SetFrom(fs *EVMNodes) {
	for _, f := range *fs {
		if f.Name == nil {
			*ns = append(*ns, f)
		} else if i := slices.IndexFunc(*ns, func(n *Node) bool {
			return n.Name != nil && *n.Name == *f.Name
		}); i == -1 {
			*ns = append(*ns, f)
		} else {
			(*ns)[i].SetFrom(f)
		}
	}
}

type EVMConfig struct {
	ChainID *utils.Big
	Enabled *bool
	Chain
	Nodes EVMNodes
}

func (c *EVMConfig) IsEnabled() bool {
	return c.Enabled == nil || *c.Enabled
}

func (c *EVMConfig) SetFrom(f *EVMConfig) {
	if f.ChainID != nil {
		c.ChainID = f.ChainID
	}
	if f.Enabled != nil {
		c.Enabled = f.Enabled
	}
	c.Chain.SetFrom(&f.Chain)
	c.Nodes.SetFrom(&f.Nodes)
}

func (c *EVMConfig) ValidateConfig() (err error) {
	if c.ChainID == nil {
		err = multierr.Append(err, v2.ErrMissing{Name: "ChainID", Msg: "required for all chains"})
	} else if c.ChainID.String() == "" {
		err = multierr.Append(err, v2.ErrEmpty{Name: "ChainID", Msg: "required for all chains"})
	} else if must, ok := ChainTypeForID(c.ChainID); ok { // known chain id
		if c.ChainType == nil && must != "" {
			err = multierr.Append(err, v2.ErrMissing{Name: "ChainType",
				Msg: fmt.Sprintf("only %q can be used with this chain id", must)})
		} else if c.ChainType != nil && *c.ChainType != string(must) {
			if *c.ChainType == "" {
				err = multierr.Append(err, v2.ErrEmpty{Name: "ChainType",
					Msg: fmt.Sprintf("only %q can be used with this chain id", must)})
			} else if must == "" {
				err = multierr.Append(err, v2.ErrInvalid{Name: "ChainType", Value: *c.ChainType,
					Msg: "must not be set with this chain id"})
			} else {
				if config.ChainType(*c.ChainType) != config.ChainOptimismBedrock {
					err = multierr.Append(err, v2.ErrInvalid{Name: "ChainType", Value: *c.ChainType,
						Msg: fmt.Sprintf("only %q can be used with this chain id", must)})
				}
			}
		}
	}

	if len(c.Nodes) == 0 {
		err = multierr.Append(err, v2.ErrMissing{Name: "Nodes", Msg: "must have at least one node"})
	} else {
		var hasPrimary bool
		for _, n := range c.Nodes {
			if n.SendOnly != nil && *n.SendOnly {
				continue
			}
			hasPrimary = true
			break
		}
		if !hasPrimary {
			err = multierr.Append(err, v2.ErrMissing{Name: "Nodes",
				Msg: "must have at least one primary node with WSURL"})
		}
	}

	err = multierr.Append(err, c.Chain.ValidateConfig())

	return
}

func (c *EVMConfig) TOMLString() (string, error) {
	b, err := toml.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type Chain struct {
	AutoCreateKey            *bool
	BlockBackfillDepth       *uint32
	BlockBackfillSkip        *bool
	ChainType                *string
	FinalityDepth            *uint32
	FlagsContractAddress     *ethkey.EIP55Address
	LinkContractAddress      *ethkey.EIP55Address
	LogBackfillBatchSize     *uint32
	LogPollInterval          *models.Duration
	LogKeepBlocksDepth       *uint32
	MinIncomingConfirmations *uint32
	MinContractPayment       *assets.Link
	NonceAutoSync            *bool
	NoNewHeadsThreshold      *models.Duration
	OperatorFactoryAddress   *ethkey.EIP55Address
	RPCDefaultBatchSize      *uint32
	RPCBlockQueryDelay       *uint16

	Transactions   Transactions      `toml:",omitempty"`
	BalanceMonitor BalanceMonitor    `toml:",omitempty"`
	GasEstimator   GasEstimator      `toml:",omitempty"`
	HeadTracker    HeadTracker       `toml:",omitempty"`
	KeySpecific    KeySpecificConfig `toml:",omitempty"`
	NodePool       NodePool          `toml:",omitempty"`
	OCR            OCR               `toml:",omitempty"`
	OCR2           OCR2              `toml:",omitempty"`
}

func (c *Chain) ValidateConfig() (err error) {
	var chainType config.ChainType
	if c.ChainType != nil {
		chainType = config.ChainType(*c.ChainType)
	}
	if !chainType.IsValid() {
		err = multierr.Append(err, v2.ErrInvalid{Name: "ChainType", Value: *c.ChainType,
			Msg: config.ErrInvalidChainType.Error()})

	} else {
		switch chainType {
		case config.ChainOptimism, config.ChainMetis:
			gasEst := *c.GasEstimator.Mode
			switch gasEst {
			case "Optimism2", "L2Suggested":
				// valid
			case "Optimism":
				err = multierr.Append(err, v2.ErrInvalid{Name: "GasEstimator.Mode", Value: gasEst,
					Msg: "unsupported since OVM 1.0 was discontinued - use L2Suggested"})
			default:
				err = multierr.Append(err, v2.ErrInvalid{Name: "GasEstimator.Mode", Value: gasEst,
					Msg: fmt.Sprintf("not allowed with ChainType %q - use L2Suggested", chainType)})
			}
		case config.ChainArbitrum, config.ChainXDai:
		}
	}

	if uint32(*c.GasEstimator.BumpTxDepth) > *c.Transactions.MaxInFlight {
		err = multierr.Append(err, v2.ErrInvalid{Name: "GasEstimator.BumpTxDepth", Value: *c.GasEstimator.BumpTxDepth,
			Msg: "must be less than or equal to Transactions.MaxInFlight"})
	}
	if *c.HeadTracker.HistoryDepth < *c.FinalityDepth {
		err = multierr.Append(err, v2.ErrInvalid{Name: "HeadTracker.HistoryDepth", Value: *c.HeadTracker.HistoryDepth,
			Msg: "must be equal to or reater than FinalityDepth"})
	}
	if *c.FinalityDepth < 1 {
		err = multierr.Append(err, v2.ErrInvalid{Name: "FinalityDepth", Value: *c.FinalityDepth,
			Msg: "must be greater than or equal to 1"})
	}
	if *c.MinIncomingConfirmations < 1 {
		err = multierr.Append(err, v2.ErrInvalid{Name: "MinIncomingConfirmations", Value: *c.MinIncomingConfirmations,
			Msg: "must be greater than or equal to 1"})
	}
	return
}

type Transactions struct {
	ForwardersEnabled    *bool
	MaxInFlight          *uint32
	MaxQueued            *uint32
	ReaperInterval       *models.Duration
	ReaperThreshold      *models.Duration
	ResendAfterThreshold *models.Duration
}

func (t *Transactions) setFrom(f *Transactions) {
	if v := f.ForwardersEnabled; v != nil {
		t.ForwardersEnabled = v
	}
	if v := f.MaxInFlight; v != nil {
		t.MaxInFlight = v
	}
	if v := f.MaxQueued; v != nil {
		t.MaxQueued = v
	}
	if v := f.ReaperInterval; v != nil {
		t.ReaperInterval = v
	}
	if v := f.ReaperThreshold; v != nil {
		t.ReaperThreshold = v
	}
	if v := f.ResendAfterThreshold; v != nil {
		t.ResendAfterThreshold = v
	}
}

type OCR2 struct {
	Automation Automation `toml:",omitempty"`
}

func (o *OCR2) setFrom(f *OCR2) {
	o.Automation.setFrom(&f.Automation)
}

type Automation struct {
	GasLimit *uint32
}

func (a *Automation) setFrom(f *Automation) {
	if v := f.GasLimit; v != nil {
		a.GasLimit = v
	}
}

type BalanceMonitor struct {
	Enabled *bool
}

func (m *BalanceMonitor) setFrom(f *BalanceMonitor) {
	if v := f.Enabled; v != nil {
		m.Enabled = v
	}
}

type GasEstimator struct {
	Mode *string

	PriceDefault *assets.Wei
	PriceMax     *assets.Wei
	PriceMin     *assets.Wei

	LimitDefault    *uint32
	LimitMax        *uint32
	LimitMultiplier *decimal.Decimal
	LimitTransfer   *uint32
	LimitJobType    GasLimitJobType `toml:",omitempty"`

	BumpMin       *assets.Wei
	BumpPercent   *uint16
	BumpThreshold *uint32
	BumpTxDepth   *uint16

	EIP1559DynamicFees *bool

	FeeCapDefault *assets.Wei
	TipCapDefault *assets.Wei
	TipCapMin     *assets.Wei

	BlockHistory BlockHistoryEstimator `toml:",omitempty"`
}

func (e *GasEstimator) ValidateConfig() (err error) {
	if uint64(*e.BumpPercent) < txpool.DefaultConfig.PriceBump {
		err = multierr.Append(err, v2.ErrInvalid{Name: "BumpPercent", Value: *e.BumpPercent,
			Msg: fmt.Sprintf("may not be less than Geth's default of %d", txpool.DefaultConfig.PriceBump)})
	}
	if e.TipCapDefault.Cmp(e.TipCapMin) < 0 {
		err = multierr.Append(err, v2.ErrInvalid{Name: "TipCapDefault", Value: e.TipCapDefault,
			Msg: "must be greater than or equal to TipCapMinimum"})
	}
	if e.FeeCapDefault.Cmp(e.TipCapDefault) < 0 {
		err = multierr.Append(err, v2.ErrInvalid{Name: "FeeCapDefault", Value: e.TipCapDefault,
			Msg: "must be greater than or equal to TipCapDefault"})
	}
	if *e.Mode == "FixedPrice" && *e.BumpThreshold == 0 && *e.EIP1559DynamicFees && e.FeeCapDefault.Cmp(e.PriceMax) != 0 {
		err = multierr.Append(err, v2.ErrInvalid{Name: "FeeCapDefault", Value: e.FeeCapDefault,
			Msg: fmt.Sprintf("must be equal to PriceMax (%s) since you are using FixedPrice estimation with gas bumping disabled in "+
				"EIP1559 mode - PriceMax will be used as the FeeCap for transactions instead of FeeCapDefault", e.PriceMax)})
	} else if e.FeeCapDefault.Cmp(e.PriceMax) > 0 {
		err = multierr.Append(err, v2.ErrInvalid{Name: "FeeCapDefault", Value: e.FeeCapDefault,
			Msg: fmt.Sprintf("must be less than or equal to PriceMax (%s)", e.PriceMax)})
	}

	if e.PriceMin.Cmp(e.PriceDefault) > 0 {
		err = multierr.Append(err, v2.ErrInvalid{Name: "PriceMin", Value: e.PriceMin,
			Msg: "must be less than or equal to PriceDefault"})
	}
	if e.PriceMax.Cmp(e.PriceDefault) < 0 {
		err = multierr.Append(err, v2.ErrInvalid{Name: "PriceMax", Value: e.PriceMin,
			Msg: "must be greater than or equal to PriceDefault"})
	}
	if *e.Mode == "BlockHistory" && *e.BlockHistory.BlockHistorySize <= 0 {
		err = multierr.Append(err, v2.ErrInvalid{Name: "BlockHistory.BlockHistorySize", Value: *e.BlockHistory.BlockHistorySize,
			Msg: "must be greater than or equal to 1 with BlockHistory Mode"})
	}

	return
}

func (e *GasEstimator) setFrom(f *GasEstimator) {
	if v := f.Mode; v != nil {
		e.Mode = v
	}
	if v := f.EIP1559DynamicFees; v != nil {
		e.EIP1559DynamicFees = v
	}
	if v := f.BumpPercent; v != nil {
		e.BumpPercent = v
	}
	if v := f.BumpThreshold; v != nil {
		e.BumpThreshold = v
	}
	if v := f.BumpTxDepth; v != nil {
		e.BumpTxDepth = v
	}
	if v := f.BumpMin; v != nil {
		e.BumpMin = v
	}
	if v := f.FeeCapDefault; v != nil {
		e.FeeCapDefault = v
	}
	if v := f.LimitDefault; v != nil {
		e.LimitDefault = v
	}
	if v := f.LimitMax; v != nil {
		e.LimitMax = v
	}
	if v := f.LimitMultiplier; v != nil {
		e.LimitMultiplier = v
	}
	if v := f.LimitTransfer; v != nil {
		e.LimitTransfer = v
	}
	if v := f.PriceDefault; v != nil {
		e.PriceDefault = v
	}
	if v := f.TipCapDefault; v != nil {
		e.TipCapDefault = v
	}
	if v := f.TipCapMin; v != nil {
		e.TipCapMin = v
	}
	if v := f.PriceMax; v != nil {
		e.PriceMax = v
	}
	if v := f.PriceMin; v != nil {
		e.PriceMin = v
	}
	e.LimitJobType.setFrom(&f.LimitJobType)
	e.BlockHistory.setFrom(&f.BlockHistory)
}

type GasLimitJobType struct {
	OCR    *uint32 `toml:",inline"`
	DR     *uint32 `toml:",inline"`
	VRF    *uint32 `toml:",inline"`
	FM     *uint32 `toml:",inline"`
	Keeper *uint32 `toml:",inline"`
}

func (t *GasLimitJobType) setFrom(f *GasLimitJobType) {
	if f.OCR != nil {
		t.OCR = f.OCR
	}
	if f.DR != nil {
		t.DR = f.DR
	}
	if f.VRF != nil {
		t.VRF = f.VRF
	}
	if f.FM != nil {
		t.FM = f.FM
	}
	if f.Keeper != nil {
		t.Keeper = f.Keeper
	}
}

type BlockHistoryEstimator struct {
	BatchSize                 *uint32
	BlockHistorySize          *uint16
	CheckInclusionBlocks      *uint16
	CheckInclusionPercentile  *uint16
	EIP1559FeeCapBufferBlocks *uint16
	TransactionPercentile     *uint16
}

func (e *BlockHistoryEstimator) setFrom(f *BlockHistoryEstimator) {
	if v := f.BatchSize; v != nil {
		e.BatchSize = v
	}
	if v := f.BlockHistorySize; v != nil {
		e.BlockHistorySize = v
	}
	if v := f.CheckInclusionBlocks; v != nil {
		e.CheckInclusionBlocks = v
	}
	if v := f.CheckInclusionPercentile; v != nil {
		e.CheckInclusionPercentile = v
	}
	if v := f.EIP1559FeeCapBufferBlocks; v != nil {
		e.EIP1559FeeCapBufferBlocks = v
	}
	if v := f.TransactionPercentile; v != nil {
		e.TransactionPercentile = v
	}
}

type KeySpecificConfig []KeySpecific

func (ks KeySpecificConfig) ValidateConfig() (err error) {
	addrs := map[string]struct{}{}
	for _, k := range ks {
		addr := k.Key.String()
		if _, ok := addrs[addr]; ok {
			err = multierr.Append(err, v2.NewErrDuplicate("Key", addr))
		} else {
			addrs[addr] = struct{}{}
		}
	}
	return
}

type KeySpecific struct {
	Key          *ethkey.EIP55Address
	GasEstimator KeySpecificGasEstimator `toml:",omitempty"`
}

type KeySpecificGasEstimator struct {
	PriceMax *assets.Wei
}

func (e *KeySpecificGasEstimator) setFrom(f *KeySpecificGasEstimator) {
	if v := f.PriceMax; v != nil {
		e.PriceMax = v
	}
}

type HeadTracker struct {
	HistoryDepth     *uint32
	MaxBufferSize    *uint32
	SamplingInterval *models.Duration
}

func (t *HeadTracker) setFrom(f *HeadTracker) {
	if v := f.HistoryDepth; v != nil {
		t.HistoryDepth = v
	}
	if v := f.MaxBufferSize; v != nil {
		t.MaxBufferSize = v
	}
	if v := f.SamplingInterval; v != nil {
		t.SamplingInterval = v
	}
}

type NodePool struct {
	PollFailureThreshold *uint32
	PollInterval         *models.Duration
	SelectionMode        *string
	SyncThreshold        *uint32
}

func (p *NodePool) setFrom(f *NodePool) {
	if v := f.PollFailureThreshold; v != nil {
		p.PollFailureThreshold = v
	}
	if v := f.PollInterval; v != nil {
		p.PollInterval = v
	}
	if v := f.SelectionMode; v != nil {
		p.SelectionMode = v
	}
	if v := f.SyncThreshold; v != nil {
		p.SyncThreshold = v
	}
}

type OCR struct {
	ContractConfirmations              *uint16
	ContractTransmitterTransmitTimeout *models.Duration
	DatabaseTimeout                    *models.Duration
	ObservationGracePeriod             *models.Duration
}

func (o *OCR) setFrom(f *OCR) {
	if v := f.ContractConfirmations; v != nil {
		o.ContractConfirmations = v
	}
	if v := f.ContractTransmitterTransmitTimeout; v != nil {
		o.ContractTransmitterTransmitTimeout = v
	}
	if v := f.DatabaseTimeout; v != nil {
		o.DatabaseTimeout = v
	}
	if v := f.ObservationGracePeriod; v != nil {
		o.ObservationGracePeriod = v
	}
}

type Node struct {
	Name     *string
	WSURL    *models.URL
	HTTPURL  *models.URL
	SendOnly *bool
}

func (n *Node) ValidateConfig() (err error) {
	if n.Name == nil {
		err = multierr.Append(err, v2.ErrMissing{Name: "Name", Msg: "required for all nodes"})
	} else if *n.Name == "" {
		err = multierr.Append(err, v2.ErrEmpty{Name: "Name", Msg: "required for all nodes"})
	}

	var sendOnly bool
	if n.SendOnly != nil {
		sendOnly = *n.SendOnly
	}
	if n.WSURL == nil {
		if !sendOnly {
			err = multierr.Append(err, v2.ErrMissing{Name: "WSURL", Msg: "required for primary nodes"})
		}
	} else if n.WSURL.IsZero() {
		if !sendOnly {
			err = multierr.Append(err, v2.ErrEmpty{Name: "WSURL", Msg: "required for primary nodes"})
		}
	} else {
		switch n.WSURL.Scheme {
		case "ws", "wss":
		default:
			err = multierr.Append(err, v2.ErrInvalid{Name: "WSURL", Value: n.WSURL.Scheme, Msg: "must be ws or wss"})
		}
	}

	if n.HTTPURL == nil {
		err = multierr.Append(err, v2.ErrMissing{Name: "HTTPURL", Msg: "required for all nodes"})
	} else if n.HTTPURL.IsZero() {
		err = multierr.Append(err, v2.ErrEmpty{Name: "HTTPURL", Msg: "required for all nodes"})
	} else {
		switch n.HTTPURL.Scheme {
		case "http", "https":
		default:
			err = multierr.Append(err, v2.ErrInvalid{Name: "HTTPURL", Value: n.HTTPURL.Scheme, Msg: "must be http or https"})
		}
	}

	return
}

func (n *Node) SetFrom(f *Node) {
	if f.Name != nil {
		n.Name = f.Name
	}
	if f.WSURL != nil {
		n.WSURL = f.WSURL
	}
	if f.HTTPURL != nil {
		n.HTTPURL = f.HTTPURL
	}
	if f.SendOnly != nil {
		n.SendOnly = f.SendOnly
	}
}
