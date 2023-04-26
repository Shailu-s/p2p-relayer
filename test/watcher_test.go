package test

import (
	"fmt"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"github.com/volmexfinance/relayers/internal/testlib"
	"github.com/volmexfinance/relayers/relayer-srv/watcher"

	"github.com/sirupsen/logrus"
	"github.com/volmexfinance/relayers/relayer-srv/worker"
)

func readWorkerTestConfig(chain string) worker.WorkerConfig {
	return worker.WorkerConfig{
		ChainName:           strings.ToUpper(chain),
		WorkerAddr:          common.HexToAddress(viper.GetString(fmt.Sprintf("workers.%s.worker_addr", chain))),
		PositioningContract: common.HexToAddress(viper.GetString(fmt.Sprintf("workers.%s.positioning_contract", chain))),
		PeripheryContract:   common.HexToAddress(viper.GetString(fmt.Sprintf("workers.%s.periphery_contract", chain))),
		GnosisContract:      common.HexToAddress(viper.GetString(fmt.Sprintf("workers.%s.gnosis_contract", chain))),
		MatchingEngine:      common.HexToAddress(viper.GetString(fmt.Sprintf("workers.%s.matching_engine_contract", chain))),
		Provider:            viper.GetString(fmt.Sprintf("workers.%s.provider", chain)),
		GasLimit:            viper.GetInt64(fmt.Sprintf("workers.%s.gas_limit", chain)),
		StartBlockHeight:    big.NewInt(viper.GetInt64(fmt.Sprintf("workers.%s.start_block_height", chain))),
	}
}

func readWorkersTestConfig() []worker.WorkerConfig {
	chains := viper.GetStringSlice("chains")
	configs := make([]worker.WorkerConfig, len(chains))

	for _, chain := range chains {
		configs = append(configs, readWorkerTestConfig(chain))
	}

	return configs
}

func loadTestConfig(t *testing.T) []worker.WorkerConfig {
	confDir := os.Getenv("FILE_PATH")
	require.NotEmptyf(t, confDir, "FILE_PATH env variable is empty")
	confFilename := os.Getenv("FILE_NAME")
	require.NotEmptyf(t, confFilename, "FILE_NAME env variable is empty")

	viper.SetConfigName(confFilename)
	viper.SetConfigType("json")
	viper.AddConfigPath(confDir)

	err := viper.ReadInConfig()
	require.NoError(t, err)

	workersConfig := readWorkersTestConfig()
	require.Greater(t, len(workersConfig), 0)

	return workersConfig
}

func TestWatcherService(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	logger := logrus.New()
	wrkrCfg := loadTestConfig(t)

	testDb := testlib.NewTestDB(t)
	nw := make(map[string]*worker.Worker)

	for _, cfg := range wrkrCfg {
		nw[cfg.ChainName] = worker.NewWorker(logger, cfg, "", testDb)
	}
	wrkrs := nw["ETH"]

	gnosisOwnerRes := make(chan *watcher.GnosisChannel)

	watcherService, err := watcher.NewWatcherSRV(wrkrs, testDb, logger, gnosisOwnerRes, "")

	if err != nil {
		t.Errorf(err.Error())
	}
	watcherService.Run()

	// do contract interactions here to fire event
}
