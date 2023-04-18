package relayer_srv

import (
	"context"
	"time"

	cr "github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
	protocols_p2p "github.com/volmexfinance/relayers/relayer-srv/chat"
	"github.com/volmexfinance/relayers/relayer-srv/db"
	matching_engine "github.com/volmexfinance/relayers/relayer-srv/matching-engine"
	postgresDB "github.com/volmexfinance/relayers/relayer-srv/postgresDB"
	"github.com/volmexfinance/relayers/relayer-srv/utils"

	"github.com/volmexfinance/relayers/relayer-srv/p2p"
	"github.com/volmexfinance/relayers/relayer-srv/watcher"
	"github.com/volmexfinance/relayers/relayer-srv/worker"
)

const TimeOut = 30 * time.Second

// MatchingConfig stores required config for matching criteria and consensus
type MatchingConfig struct {
	Leader         bool  `json:"leader"`
	MaxFailAllowed int64 `json:"max_fail_allowed"`
}

// NodeDetails stores node details
type NodeDetails struct {
	WorkerAddress string `json:"worker_address"`
	PrivateKey    string `json:"private_key"`
}

type RelayerSrv struct {
	ctx            context.Context
	logger         *logrus.Entry
	node           *p2p.Node
	db             *db.DataBase
	postgresDB     *postgresDB.PostgresDataBase
	matchingCfg    MatchingConfig
	workers        map[string]*worker.Worker
	obsRecieveRes  <-chan *protocols_p2p.GossipMessage
	obsSendReq     chan<- *protocols_p2p.GossipMessage
	watcher        map[string]*watcher.WatcherSRV
	GnosisOwnerRes chan *watcher.GnosisChannel
}

// /TODO: Leader handling should be done by DB queries
func NewRelayerSrv(ctx context.Context, logger *logrus.Logger, dbUrl string, dbConfig db.Config, wrkrCfg []worker.WorkerConfig, p2pCfg p2p.Config, matchingCfg MatchingConfig, nodeDetails NodeDetails) *RelayerSrv {

	dbConn, err := db.NewDataBase(logger, dbConfig)
	if err != nil {
		logger.Panicf("Unable to create db connection: %v", err)
	}

	err = dbConn.InitialMigration()
	if err != nil {
		logger.Warnf("Unable to run migration: %v", err)
	}

	//TODO: Create compose for postgresDB
	// postgresDB, err := postgresDB.InitialMigration(dbUrl)
	// if err != nil {
	// 	logger.Warnf("Unable to run migration: %s", err)
	// }

	obsRecieveRes := make(chan *protocols_p2p.GossipMessage, 1)
	obsSendReq := make(chan *protocols_p2p.GossipMessage, 1)

	workers := make(map[string]*worker.Worker)
	watchers := make(map[string]*watcher.WatcherSRV)
	gnosisOwnerRes := make(chan *watcher.GnosisChannel)
	guardianSetMapping := make(map[string][]p2p.GuardianInfo)
	for _, cfg := range wrkrCfg {
		workers[cfg.ChainName] = worker.NewWorker(logger, cfg, nodeDetails.PrivateKey, dbConn)
		wrkr := workers[cfg.ChainName]
		owners, err := wrkr.GetGnosisOwners()
		if err != nil {
			logger.Panicf("Unable to get gnosis owners: %v", err)
		}
		var guardianSet []p2p.GuardianInfo
		for i := 0; i < len(owners); i++ {
			guardianSet = append(guardianSet, p2p.GuardianInfo{Address: string(owners[i].Hex())})
		}
		guardianSetMapping[cfg.ChainName] = guardianSet
		threshold, err := wrkr.GetThreshold()
		if err != nil {
			logger.Panicf("Unable to get gnosis owners: %v", err)
		}
		wrkr.Threshold = threshold.Int64()
		watcherService, err := watcher.NewWatcherSRV(wrkr, dbConn, logger, gnosisOwnerRes, cfg.ChainName)
		if err != nil {
			logger.Warnf("Failed to create watcher srv: %v", err)
		}
		watchers[cfg.ChainName] = watcherService
	}
	inst := &RelayerSrv{
		ctx:           ctx,
		logger:        logger.WithField("layer", "relayer"),
		node:          p2p.NewP2PNode(ctx, logger, obsRecieveRes, obsSendReq, nodeDetails.PrivateKey, p2pCfg, guardianSetMapping),
		db:            dbConn,
		postgresDB:    nil,
		workers:       workers,
		obsRecieveRes: obsRecieveRes,
		obsSendReq:    obsSendReq,
		matchingCfg:   matchingCfg,
		watcher:       watchers,
	}
	return inst
}

// sign binary data using the local node's private key
func signData(data []byte, privKey string) ([]byte, error) {
	privateKey, _ := utils.GetPrivateKey(privKey)
	sig, err := cr.Sign(data, privateKey)
	return sig, err
}

func (r *RelayerSrv) Run() {
	r.node.Run()

	for i, worker := range r.workers {
		go r.MatchAndSendToP2P(worker)
		go r.RetryMatching(worker)
		r.watcher[i].Run()
		go r.UpdateTx(worker)
		// go r.MoveSQLiteToPostgres(worker)
	}

	// TODO: Solve error in this

	<-r.ctx.Done()
	r.Stop()
}

func (r *RelayerSrv) Stop() {
	r.node.Stop()
}

func (r *RelayerSrv) UpdateTx(worker *worker.Worker) {
	for {
		er := r.UpdateTxnSentStatus(worker)
		if er != nil {
			r.logger.Warnf("UpdateTxnSentStatus: %s", er.Error())
		}
		time.Sleep(15 * time.Second)
	}
}

func (r *RelayerSrv) MatchAndSendToP2P(wrkr *worker.Worker) {
	//TODO: error handling
	for {
		//TODO: to be changed to batch again
		orders, err := r.db.GetZeroOrders(wrkr.ChainName)
		if err != nil {
			r.logger.Warnf("Not any new order found in DB %v", err)
			continue
		}
		if len(orders) == 0 {
			r.logger.Warn("Not any new order found in Database")
		} else {
			err = r.db.UpdateBatchOrderStatus(orders, db.MatchedStatusInit)
			if err != nil {
				r.logger.Errorf("Run: Found error in update %s", err.Error())
			}
		}
		time.Sleep(TimeOut)
		//TODO: to be changed to batch again
		order1, order2, err := matching_engine.MatchBatchDBOrders(r.db, wrkr, r.matchingCfg.MaxFailAllowed)
		if err != nil {
			r.logger.Warnf("Not any match found yet %v", err)
			continue
		}
		if len(order1) == 0 || len(order2) == 0 {
			time.Sleep(TimeOut)
			// time.Sleep(2 * time.Minute)
			r.logger.Warn("Not any match found yet")
			continue
		}
		orderID := []string{}
		_, hash, err := wrkr.CreateGnosisTxAndHash(order1, order2)
		if err != nil {
			r.logger.Errorf("Error in MatchAndSendToP2P: %s", err.Error())
		}
		sign, err := signData(hash, wrkr.PrivateKey)
		if err != nil {
			r.logger.Errorf("Error in MatchAndSendToP2P: signData%s", err.Error())
		}
		signToSend := [][]byte{}
		signToSend = append(signToSend, sign)
		_, err = r.SendToContract(order1, order2, orderID, signToSend, wrkr)
		if err != nil {
			r.logger.Errorf("Error in MatchAndSendToP2P: SendToContract%s", err.Error())
		}
		time.Sleep(TimeOut)
		continue
	}
}

// TODO: Handle lost and not found transaction/// if not found, then check if txn after that is success or failed, then convert not found to lost,
func (r *RelayerSrv) RetryMatching(wrkr *worker.Worker) {
	for {
		orders, err := r.db.GetOrdersOnStatus(wrkr.ChainName, []db.MatchedStatus{db.MatchedStatusSentFailed})
		if err != nil {
			r.logger.Warnf("Unable to get MatchedStatusSentFailed: %v", err)
			time.Sleep(10 * time.Second)
		}
		if len(orders) > 0 {
			for _, order := range orders {
				if _, err := wrkr.OrderValidation(*order); err != nil {
					if err := r.db.UpdateOrderStatusAndFailCount(order.OrderID, db.MatchedStatusBlocked); err != nil {
						r.logger.Warnf("UpdateOrderStatusById : %s", err)
						continue
					}
				} else {
					if err := r.db.UpdateOrderStatusAndFailCount(order.OrderID, db.MatchedStatusFailedConfirmed); err != nil {
						r.logger.Warnf("UpdateOrderStatusById : %s", err)
						continue
					}
				}
			}
		}
		txns, err := r.db.GetTxnsOnStatus([]db.TransactionStatusType{db.TransactionStatusTypeNotFound, db.TransactionStatusTypeFailed}, wrkr.ChainName)
		if err != nil {
			r.logger.Warnf("Unable to get any not found transaction: %v", err)
			time.Sleep(10 * time.Second)
			continue
		}
		if len(*txns) > 0 {
			r.logger.Infof("Found %d transactions to retry in DB", len(*txns))
		}

		for _, txn := range *txns {
			if txn.TransactionStatus == db.TransactionStatusTypeNotFound {
				// nextTxn, err := r.db.GetTxnByNonce(txn.Nonce + 1)
				// if err != nil {
				// 	r.logger.Warnf("Unable to get txn from next nonce : %v", err)
				// 	continue
				// }
				// if nextTxn.TransactionStatus == db.TransactionStatusTypeFailed || nextTxn.TransactionStatus == db.TransactionStatusTypeSuccess {
				// 	for _, orderId := range txn.OrderID {
				// 		if err := r.db.UpdateOrderStatusAndResetFill(orderId, db.MatchedStatusFailedConfirmed); err != nil {
				// 			r.logger.Warnf("UpdateOrderStatusById : %s", err)
				// 			continue
				// 		}
				// 	}

				// 	if err := r.db.UpdateTxnStatus(&txn, db.TransactionStatusTypeLost); err != nil {
				// 		r.logger.Warnf("Unable to update lost transaction : %v", err)
				// 		continue
				// 	}
				// }
			} else {
				for _, orderId := range txn.OrderID {
					order, err := r.db.FindOrder(orderId)
					if err != nil {
						r.logger.Warnf("Unable to get order from order id: %v", err)
						continue
					}

					if _, err := wrkr.OrderValidation(*order); err != nil {
						if err := r.db.UpdateOrderStatusAndFailCount(orderId, db.MatchedStatusBlocked); err != nil {
							r.logger.Warnf("UpdateOrderStatusById : %s", err)
							continue
						}
					} else {
						if err := r.db.UpdateOrderStatusAndFailCount(orderId, db.MatchedStatusFailedConfirmed); err != nil {
							r.logger.Warnf("UpdateOrderStatusById : %s", err)
							continue
						}
					}
					if err := r.db.UpdateTxnStatus(&txn, db.TransactionStatusTypeFailedConfirmed); err != nil {
						r.logger.Warnf("UpdateTxnStatus : %s", err)
						continue
					}
				}
			}
		}
		time.Sleep(20 * time.Second)
	}
}

func (r *RelayerSrv) GetChainID(chainName string) int64 {
	wrkr := r.workers[chainName]
	return wrkr.GetChainID()
}

func (r *RelayerSrv) GetPeripheryContract(chain string) string {
	wrkr := r.workers[chain]
	return wrkr.GetPeripheryContract().String()
}

func (r *RelayerSrv) GetPositioningContract(chain string) string {
	wrkr := r.workers[chain]
	return wrkr.GetPositioningContract().String()
}
