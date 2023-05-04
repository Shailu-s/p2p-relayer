package watcher

import (
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"github.com/volmexfinance/relayers/relayer-srv/big_ext"
	"github.com/volmexfinance/relayers/relayer-srv/db"
	"github.com/volmexfinance/relayers/relayer-srv/utils"
	"github.com/volmexfinance/relayers/relayer-srv/worker"
	"gorm.io/gorm"
)

// WatcherSRV is an instance of watcher and stores related information
type WatcherSRV struct {
	Logs             chan types.Log
	Sub              ethereum.Subscription
	Logger           *logrus.Logger
	PostgresDataBase *db.PostgresDataBase
	SQLiteDataBase   *db.SQLiteDataBase
	Worker           *worker.Worker
	GnosisOwnerRes   chan *GnosisChannel
}
type GnosisChannel struct {
	Address   common.Address
	Threshold *big.Int
	EventType string
	Chain     string
}

var DUST = big.NewInt(1000)

// NewWatcherSRV returns new instance of watcher
func NewWatcherSRV(wrkr *worker.Worker, db *db.PostgresDataBase, sqlDB *db.SQLiteDataBase, logger *logrus.Logger, gnosisOwnerRes chan *GnosisChannel, chain string) (*WatcherSRV, error) {
	logs := make(chan types.Log)

	subs, err := wrkr.SubscribeToLogs(logs)
	if err != nil {
		return nil, err
	}

	return &WatcherSRV{
		Logs:             logs,
		Sub:              subs,
		Logger:           logger,
		PostgresDataBase: db,
		SQLiteDataBase:   sqlDB,
		Worker:           wrkr,
		GnosisOwnerRes:   gnosisOwnerRes,
	}, nil
}

// getSubscribedEventLog used to subscribe events and handle events according to its use case
func (w *WatcherSRV) getSubscribedEventLog(logs chan types.Log) error {
	// ticker := time.NewTicker(1 * time.Minute)
	// defer ticker.Stop()
	for {
		var tLog db.TransactionLog

		select {
		// case <-ticker.C:

		// 	for {
		// 		lock.Lock()
		// 		w.Logger.Infof("Working on RedialClient")
		// 		err := w.Worker.RedialClient()
		// 		if err != nil {
		// 			w.Logger.Errorf("Found error in RedialClient")
		// 		} else {
		// 			w.Logger.Infof("RedialClient complete")
		// 			lock.Unlock()
		// 			break
		// 		}
		// 		lock.Unlock()
		// 	}

		case <-w.Sub.Err():
			w.Logger.Errorf("Error received in event subscription")
			var err error
			for {
				w.Logger.Errorf("Trying to renew subscription")
				w.Sub, err = w.Worker.SubscribeToLogs(w.Logs)
				if err != nil {
					w.Logger.Errorf("Found error in renewing subscription")
					time.Sleep(10 * time.Second)
				} else {
					break
				}
			}
			// continue
		case vLog := <-w.Logs:
			tLog.BlockHeight = vLog.BlockNumber
			tLog.TransactionHash = vLog.TxHash.Hex()
			event, eventType, err := worker.ParseEvent(&vLog)
			if err != nil {
				w.Logger.Infof("Error: %v\n", err)
				continue
			}
			lastTxnLog, err := w.PostgresDataBase.FindLatestTxnLogByHash(tLog.TransactionHash)
			switch eventType {
			case "OrdersFilled":
				MatchingEngineOrdersFilled := event.(worker.MatchingEngineAbiOrdersFilled)
				for i := 0; i < 2; i++ {
					tLog.Traders = append(tLog.Traders, MatchingEngineOrdersFilled.Traders[i].String())
					tLog.Salt = append(tLog.Salt, MatchingEngineOrdersFilled.Salts[i].String())
				}

				tLog.NewLeftFill = MatchingEngineOrdersFilled.Fills[0].String()
				tLog.NewRightFill = MatchingEngineOrdersFilled.Fills[1].String()
				tLog.Status = db.TransactionLogStatusTypeConfirmed

				// Get order IDs
				orderId1, orderId2 := utils.GetOrderIdsFromTLog(&tLog, w.Worker.ChainName)

				tLog.OrderID = append(tLog.OrderID, orderId1, orderId2)
				for i, order_id := range tLog.OrderID {
					order, err := w.PostgresDataBase.FindOrder(order_id)
					if err != nil {
						if errors.Is(err, gorm.ErrRecordNotFound) {
							continue
						}
						w.Logger.Warnf("getSubscribedEventLog :%v", err)
						continue
					}

					currentFill := order.OrderFills()

					var assetValue *big.Int
					if order.CreatedAt == 0 {
						continue
					}
					assetValue = order.MakeAsset().ValueAsBigInt()

					var newFill *big.Int
					if i == 0 {
						newFill = MatchingEngineOrdersFilled.Fills[0]
					} else {
						newFill = MatchingEngineOrdersFilled.Fills[1]
					}
					if lastTxnLog.OrderID[0] == order_id {
						fill, err := new(big.Int).SetString(lastTxnLog.NewLeftFill, 10)
						if big_ext.LessThan(newFill, fill) {
							w.Logger.Warnf("getSubscribedEventLog :%v", err)
							continue
						}
					} else if lastTxnLog.OrderID[1] == order_id {
						fill, err := new(big.Int).SetString(lastTxnLog.NewRightFill, 10)
						if big_ext.LessThan(newFill, fill) {
							w.Logger.Warnf("getSubscribedEventLog :%v", err)
							continue
						}
					}
					result := new(big.Int).Sub(currentFill, newFill)
					if result.Cmp(DUST) > 0 {
						w.Logger.Warnf("getSubscribedEventLog: Fill stored %s greater than fill fetched from event %s", currentFill.String(), newFill.String())
					}
					resultFills := new(big.Int).Sub(assetValue, newFill)

					if resultFills.Cmp(DUST) <= 0 {
						erOrder := w.PostgresDataBase.HandleOrderStatusAndFills(order.OrderID, newFill.String(), []db.MatchedStatus{}, []db.MatchedStatus{}, db.MatchedStatusFullMatchConfirmed)
						if erOrder != nil {
							w.Logger.Warnf("getSubscribedEventLog :%v", err)
							continue
						}
						errOrder := w.SQLiteDataBase.HandleOrderStatusAndFills(order.OrderID, newFill.String(), []db.MatchedStatus{}, []db.MatchedStatus{}, db.MatchedStatusFullMatchConfirmed)
						if errOrder != nil {
							w.Logger.Warnf("getSubscribedEventLog :%v", err)
							continue
						}
					} else {
						erOrder := w.PostgresDataBase.HandleOrderStatusAndFills(order.OrderID, newFill.String(), []db.MatchedStatus{}, []db.MatchedStatus{}, db.MatchedStatusPartialMatchConfirmed)
						if erOrder != nil {
							w.Logger.Warnf("getSubscribedEventLog :%v", err)
							continue
						}
						errOrder := w.SQLiteDataBase.HandleOrderStatusAndFills(order.OrderID, newFill.String(), []db.MatchedStatus{}, []db.MatchedStatus{}, db.MatchedStatusPartialMatchConfirmed)
						if errOrder != nil {
							w.Logger.Warnf("getSubscribedEventLog :%v", err)
							continue
						}
					}
				}
				// Create TransactionLog
				er := w.PostgresDataBase.CreateTransactionLog(&tLog)
				if er != nil {
					// wg.Done()
					w.Logger.Infof("Error: %v\n", er)
					continue
				}
				err := w.SQLiteDataBase.CreateTransactionLog(&tLog)
				if err != nil {
					// wg.Done()
					w.Logger.Infof("Error: %v\n", er)
					continue
				}
				w.Logger.Infof("Found matched event and handled successfully with orderID %s and order ID %s", orderId1, orderId2)
			case "Canceled":
				canceledOrder := event.(worker.MatchingEngineAbiCanceled)

				orderID := db.CreateOrderID(canceledOrder.Trader.String(), canceledOrder.Salt.String(), w.Worker.ChainName)

				if ok := w.PostgresDataBase.UpdateOrderStatus(orderID, []db.MatchedStatus{}, []db.MatchedStatus{}, db.Canceled); ok != nil {
					w.Logger.Warnf("getSubscribedEventLog : Error in handle cancel order %v", err)
					continue
				}
				if ok := w.SQLiteDataBase.UpdateOrderStatus(orderID, []db.MatchedStatus{}, []db.MatchedStatus{}, db.Canceled); ok != nil {
					w.Logger.Warnf("getSubscribedEventLog : Error in handle cancel order %v", err)
					continue
				}
				w.Logger.Infof("Found canceledOrder event and handled successfully with orderID %s", orderID)
			case "CanceledAll":
				canceledAllOrder := event.(worker.MatchingEngineAbiCanceledAll)
				if ok := w.PostgresDataBase.UpdateOrderStatusByMinSalt(canceledAllOrder.Trader.String(), canceledAllOrder.MinSalt.String(), []db.MatchedStatus{db.MatchedStatusInit, db.MatchedStatusBlocked, db.MatchedStatusFailedConfirmed, db.MatchedStatusFullMatchFound, db.MatchedStatusPartialMatchFound, db.MatchedStatusPartialMatchConfirmed, db.MatchedStatusSentFailed}, []db.MatchedStatus{db.MatchedStatusValidated, db.MatchedStatusValidationConfirmed, db.MatchedStatusSentToContract, db.MatchedStatusFullMatchConfirmed}, db.Canceled, w.Worker.ChainName); ok != nil {
					w.Logger.Warnf("getSubscribedEventLog : Error in handle cancel order %v", ok)
					continue
				}
				if ok := w.SQLiteDataBase.UpdateOrderStatusByMinSalt(canceledAllOrder.Trader.String(), canceledAllOrder.MinSalt.String(), []db.MatchedStatus{db.MatchedStatusInit, db.MatchedStatusBlocked, db.MatchedStatusFailedConfirmed, db.MatchedStatusFullMatchFound, db.MatchedStatusPartialMatchFound, db.MatchedStatusPartialMatchConfirmed, db.MatchedStatusSentFailed}, []db.MatchedStatus{db.MatchedStatusValidated, db.MatchedStatusValidationConfirmed, db.MatchedStatusSentToContract, db.MatchedStatusFullMatchConfirmed}, db.Canceled, w.Worker.ChainName); ok != nil {
					w.Logger.Warnf("getSubscribedEventLog : Error in handle cancel order %v", ok)
					continue
				}
				w.Logger.Infof("Found CanceledAll event and handled successfully with trader %s", canceledAllOrder.Trader.String())
			case "AddedOwner":
				newOwner := event.(worker.GnosisAddedOwner)
				w.GnosisOwnerRes <- &GnosisChannel{Address: newOwner.Owner, EventType: "AddedOwner", Chain: w.Worker.ChainName}
				w.Logger.Infof("Found AddedOwner event and handled successfully with owner %s", newOwner.Owner.String())

			case "RemovedOwner":
				oldOwner := event.(worker.GnosisRemovedOwner)
				w.GnosisOwnerRes <- &GnosisChannel{Address: oldOwner.Owner, EventType: "RemovedOwner", Chain: w.Worker.ChainName}
				w.Logger.Infof("Found RemovedOwner event and handled successfully with owner %s", oldOwner.Owner.String())
			case "ChangedThreshold":
				changedThreshold := event.(worker.GnosisChangedThreshold)
				w.GnosisOwnerRes <- &GnosisChannel{Threshold: changedThreshold.Threshold, EventType: "ChangedThreshold", Chain: w.Worker.ChainName}
				w.Logger.Infof("Found ChangedThreshold event and handled successfully with threshold %d", changedThreshold.Threshold.Int64())
			}
		}
		time.Sleep(2 * time.Second)
	}
}

// create watcher service first
func (w *WatcherSRV) Run() {

	// var wg sync.WaitGroup
	// defer wg.Wait()
	go func() {
		err := w.getSubscribedEventLog(w.Logs)
		if err != nil {
			w.Logger.Infof("Error : %v\n", err)
			// wg.Done()
		}
		time.Sleep(2 * time.Second)
	}()
	// wg.Add(1)
}
