package config

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// WorkerConfig worker configuration struct
type WorkerConfig struct {
	ChainName           string         `json:"chain_name"`
	Provider            string         `json:"provider"`
	MatchingEngine      common.Address `json:"matching_engine_contract"`
	GnosisContract      common.Address `json:"gnosis_contract"`
	PositioningContract common.Address `json:"positioning_contract"`
	IndexPriceOracle    common.Address `json:"index_price_oracle"`
	MarkPriceOracle     common.Address `json:"mark_price_oracle"`
	PeripheryContract   common.Address `json:"periphery_contract"`
	GasLimit            int64          `json:"gas_limit"`
	WorkerAddr          common.Address `json:"worker_addr"`
	StartBlockHeight    *big.Int       `json:"from_block"`
}

// readWorkerConfig reads ethereum chain worker params from config.json
func (v *viperConfig) readWorkerConfig(chain string) WorkerConfig {
	return WorkerConfig{
		ChainName:           strings.ToUpper(chain),
		WorkerAddr:          common.HexToAddress(v.GetString(fmt.Sprintf("workers.%s.worker_addr", chain))),
		PositioningContract: common.HexToAddress(v.GetString(fmt.Sprintf("workers.%s.positioning_contract", chain))),
		IndexPriceOracle:    common.HexToAddress(v.GetString(fmt.Sprintf("workers.%s.index_price_oracle", chain))),
		MarkPriceOracle:     common.HexToAddress(v.GetString(fmt.Sprintf("workers.%s.mark_price_oracle", chain))),
		PeripheryContract:   common.HexToAddress(v.GetString(fmt.Sprintf("workers.%s.periphery_contract", chain))),
		GnosisContract:      common.HexToAddress(v.GetString(fmt.Sprintf("workers.%s.gnosis_contract", chain))),
		MatchingEngine:      common.HexToAddress(v.GetString(fmt.Sprintf("workers.%s.matching_engine_contract", chain))),
		Provider:            v.GetString(fmt.Sprintf("workers.%s.provider", chain)),
		GasLimit:            v.GetInt64(fmt.Sprintf("workers.%s.gas_limit", chain)),
		StartBlockHeight:    big.NewInt(v.GetInt64(fmt.Sprintf("workers.%s.start_block_height", chain))),
	}
}

func (v *viperConfig) ReadWorkersConfig() []WorkerConfig {
	chains := v.GetStringSlice("chains")
	chainCfgs := make([]WorkerConfig, 0, len(chains))

	for _, chain := range chains {
		chainCfgs = append(chainCfgs, v.readWorkerConfig(chain))
	}

	return chainCfgs
}
