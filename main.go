package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	relayer_srv "github.com/volmexfinance/relayers/relayer-srv"
	"github.com/volmexfinance/relayers/relayer-srv/db"
	"github.com/volmexfinance/relayers/relayer-srv/p2p"
	"github.com/volmexfinance/relayers/relayer-srv/worker"

	"github.com/sirupsen/logrus"
	"github.com/volmexfinance/relayers/relayer-srv/app"
	"github.com/volmexfinance/relayers/relayer-srv/config"
)

func toDbConfig(config config.SqliteDbConfig) db.Config {
	return db.Config{
		FileName: config.FileName,
	}
}

func toMatchingConfig(config config.MatchingConfig) relayer_srv.MatchingConfig {
	return relayer_srv.MatchingConfig{
		Leader:         config.Leader,
		MaxFailAllowed: config.MaxFailAllowed,
	}
}

func toWorkerConfig(config config.WorkerConfig) worker.WorkerConfig {
	return worker.WorkerConfig{
		ChainName:           config.ChainName,
		Provider:            config.Provider,
		MatchingEngine:      config.MatchingEngine,
		GnosisContract:      config.GnosisContract,
		PositioningContract: config.PositioningContract,
		PeripheryContract:   config.PeripheryContract,
		IndexPriceOracle:    config.IndexPriceOracle,
		MarkPriceOracle:     config.MarkPriceOracle,
		GasLimit:            config.GasLimit,
		WorkerAddr:          config.WorkerAddr,
		StartBlockHeight:    config.StartBlockHeight,
	}
}

func toNodeDetailsConfig(config config.NodeDetailsConfig) relayer_srv.NodeDetails {
	return relayer_srv.NodeDetails{
		WorkerAddress: config.WorkerAddress,
		PrivateKey:    config.PrivateKey,
	}
}

func toWorkerConfigs(configs []config.WorkerConfig) []worker.WorkerConfig {
	workerConfs := make([]worker.WorkerConfig, 0, len(configs))
	for _, conf := range configs {
		workerConfs = append(workerConfs, toWorkerConfig(conf))
	}
	return workerConfs
}

func toP2pConfig(config config.P2PConfig) p2p.Config {
	return p2p.Config{
		Port:           config.Port,
		ListenAddress:  config.ListenAddress,
		DiscoveryTag:   config.DiscoveryTag,
		RelayerAddress: config.RelayerAddress,
	}
}

func main() {
	cfg := config.NewViperConfig()
	loggerLevel := cfg.ReadLogLevelConfig()
	sqliteConfig := cfg.GetSqliteConfig()
	workersCfg := cfg.ReadWorkersConfig()
	nodeDetailsCfg := cfg.ReadNodeDetailsConfig()
	matchingCfg := cfg.ReadMatchingConfig()
	srvURL := cfg.ReadServiceConfig()
	postgresDbConfig := cfg.ReadDBConfig()
	p2pCfg := cfg.ReadP2PConfig()

	logger := NewRootLogger(loggerLevel)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sign := <-c
		logger.Infof("System signal: %+v\n", sign)
		cancel()
	}()

	nodeDetails := toNodeDetailsConfig(nodeDetailsCfg)
	dbConfig := toDbConfig(sqliteConfig)
	matchingConfig := toMatchingConfig(matchingCfg)
	workerConfigs := toWorkerConfigs(workersCfg)
	p2pConfig := toP2pConfig(p2pCfg)

	dbURL := postgresDbConfig.AsPostgresDbUrl()

	relayerSrvInstance := relayer_srv.NewRelayerSrv(ctx, logger, dbURL, dbConfig, workerConfigs, p2pConfig, matchingConfig, nodeDetails)
	appInstance := app.NewApp(logger, srvURL, relayerSrvInstance)

	appInstance.Run(ctx)
}

func NewRootLogger(loggerLevel string) *logrus.Logger {
	// Initialize root logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
	})

	logLevel, err := logrus.ParseLevel(loggerLevel)
	if err != nil {
		panic(fmt.Errorf("failed to parse log level: %v", err))
	}
	logger.SetLevel(logLevel)

	err = os.MkdirAll("./logs", os.ModePerm)
	if err != nil {
		panic(fmt.Errorf("failed to create logs directory: %v", err))
	}

	logFile, err := os.OpenFile("./logs/relayer.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(fmt.Errorf("error opening file: %v", err))
	}
	logger.SetOutput(logFile)
	return logger
}
