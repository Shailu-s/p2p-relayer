package matching_engine

import (
	"errors"
	"fmt"
	"github.com/volmexfinance/relayers/relayer-srv/db"
	"github.com/volmexfinance/relayers/relayer-srv/db/models"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/volmexfinance/relayers/relayer-srv/worker"
)

// /TODO: Add function so it can revert whole changes if error happened after filling or matching
// MatchOrders takes a order and matches it require order based on price and assets
// returns two matched orders
func MatchOrders(order *models.Order, database *db.DataBase, maxFail int64, w *worker.Worker) (models.Order, models.Order, error) {
	var priorityList []models.Order
	var err error
	if order.IsShort {
		priorityList, err = CreatePriorityList(database, false, "Price Desc", w.ChainName)
		if err != nil {
			return models.Order{}, models.Order{}, fmt.Errorf("UNABLE TO GET PRIORITY QUEUE: %w", err)
		}
		logrus.Infof("Length of buy orders %d", len(priorityList))
	} else {
		priorityList, err = CreatePriorityList(database, true, "Price", w.ChainName)
		if err != nil {
			return models.Order{}, models.Order{}, fmt.Errorf("UNABLE TO GET PRIORITY QUEUE: %w", err)
		}
		logrus.Infof("Length of sell orders %d", len(priorityList))
	}

	for i := 0; i < len(priorityList); i++ {
		matchingOrder := &priorityList[i]
		// logger := logrus.New()

		logrus.WithFields(logrus.Fields{"right_order": matchingOrder.OrderID}).Debugf("matching order with price %f", matchingOrder.Price)
		// switch order.OrderType {
		// case models.STOP_LOSS_LAST_PRICE:
		// 	checkPriceStopLoss(*order, *matchingOrder, logger)

		// case models.TAKE_PROFIT_LAST_PRICE:
		// 	checkPriceTakeProfit(*order, *matchingOrder, logger)
		// }

		// switch matchingOrder.OrderType {
		// case models.STOP_LOSS_LAST_PRICE:
		// 	checkPriceStopLoss(*matchingOrder, *order, logger)

		// case models.TAKE_PROFIT_LAST_PRICE:
		// 	checkPriceTakeProfit(*matchingOrder, *order, logger)
		// }
		validated, err := w.ValidateOrder(matchingOrder)
		if err != nil {
			w.Logger.Errorf("Error in validaing limit order %s", err.Error())
			continue
		}
		if !validated {
			continue
		}

		if order.Trader == matchingOrder.Trader ||
			order.MakeAsset().VirtualToken != matchingOrder.TakeAsset().VirtualToken ||
			order.TakeAsset().VirtualToken != matchingOrder.MakeAsset().VirtualToken {
			logrus.Infof("Order ID %s not matched with order ID %s ", order.OrderID, priorityList[i].OrderID)
			continue
		}

		if !order.IsShort && order.Price < priorityList[i].Price || order.IsShort && order.Price > priorityList[i].Price {
			logrus.WithFields(logrus.Fields{
				"left_order":  order.OrderID,
				"right_order": matchingOrder.OrderID,
			}).Warnf("price is not suitable (%f, %f)", order.Price, matchingOrder.Price)
			continue
		}

		if matchingOrder.FailCount > maxFail {
			err = database.UpdateOrderStatus(priorityList[i].OrderID, []models.MatchedStatus{}, []models.MatchedStatus{}, models.MatchedStatusBlocked)
			if err != nil {
				return models.Order{}, models.Order{}, fmt.Errorf("unable to update failed blocked order: %w", err)
			}
			return models.Order{}, models.Order{}, fmt.Errorf("unable to match failed order: %w", err)
		}

		if order.FailCount > 0 {
			txnSents, err := database.GetTxnByOrderID(order.OrderID)
			if err != nil {
				logrus.Info("Unable to find txn with this order ID in DB")
				continue
			}
			skip := false
			for _, txnSent := range txnSents {
				for _, orderID := range txnSent.OrderID {
					if strings.EqualFold(orderID, priorityList[i].OrderID) {
						skip = true
						break
					}
				}
			}
			if skip {
				logrus.Info("Matched previously with same order")
				continue
			}
		}

		err = matchAndUpdateOrders(order, matchingOrder)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"left_order":  order.OrderID,
				"right_order": matchingOrder.OrderID,
			}).Warnf("unable to match orders: %v", err)
			continue
		}

		logrus.Infof("Matched order %s with order %s", order.OrderID, matchingOrder.OrderID)

		result, err := database.FindOrder(order.OrderID)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"order_id": order.OrderID,
			}).Warnf("order not found in models: %v", err)
			continue
		}

		if result == nil {
			if err := database.CreateOrder(order); err != nil {
				logrus.WithFields(logrus.Fields{
					"order_id": order.OrderID,
				}).Warnf("unable to create models order: %v", err)
				continue
			}
		}

		if err := database.UpdateOrder(order); err != nil {
			logrus.WithFields(logrus.Fields{
				"order_id": order.OrderID,
			}).Warnf("unable to update models order: %v", err)
			continue
		}

		if err := database.UpdateOrder(matchingOrder); err != nil {
			logrus.WithFields(logrus.Fields{
				"order_id": matchingOrder.OrderID,
			}).Warnf("unable to update models order: %v", err)
			continue
		}

		return *order, *matchingOrder, nil
	}

	return models.Order{}, models.Order{}, errors.New("no match found")
}

func MatchBuyDBOrders(database *db.DataBase, w *worker.Worker, maxFail int64) ([]*models.Order, []*models.Order, error) {
	buyPriorityList, err := CreatePriorityList(database, false, "Price desc", w.ChainName)
	if err != nil {
		return nil, nil, fmt.Errorf("UNABLE TO GET ORDER FROM DB: %w", err)
	}
	logrus.Infof("Length of buy orders in DB %d", len(buyPriorityList))

	for i := 0; i < len(buyPriorityList); i++ {
		validated, err := w.ValidateOrder(&buyPriorityList[i])
		if err != nil {
			w.Logger.Errorf("Error in validaing limit order %s", err.Error())
			continue
		}
		if !validated {
			logrus.Infof("Buy limit order price verification failed with  ID: %s trader: %s ", buyPriorityList[i].OrderID, buyPriorityList[i].Trader)
			continue
		}
		logrus.Infof("Working on buy order with price %f; trader: %s", buyPriorityList[i].Price, buyPriorityList[i].Trader)

		if buyPriorityList[i].FailCount > maxFail {
			err = database.UpdateOrderStatus(buyPriorityList[i].OrderID, []models.MatchedStatus{}, []models.MatchedStatus{}, models.MatchedStatusBlocked)
			if err != nil {
				return nil, nil, fmt.Errorf("tried more that twice, can't process order: %w", err)
			}
			return nil, nil, fmt.Errorf("tried more that twice, can't process order: %w", err)
		}
		if buyPriorityList[i].Status == models.MatchedStatusSentFailed {
			if _, err := w.OrderValidation(buyPriorityList[i]); err != nil {
				if err := database.UpdateOrderStatusAndFailCount(buyPriorityList[i].OrderID, models.MatchedStatusBlocked); err != nil {
					return nil, nil, fmt.Errorf("unable to updatestatus %e", err)
				}
			} else {
				if err := database.UpdateOrderStatusAndFailCount(buyPriorityList[i].OrderID, models.MatchedStatusFailedConfirmed); err != nil {
					return nil, nil, fmt.Errorf("unable to updatestatus %e", err)
				}
				continue
			}

		}
		if buyPriorityList[i].OrderID != "" {
			order1, order2, err := MatchOrders(&buyPriorityList[i], database, maxFail, w)
			if err != nil {
				logrus.Infof("Unable to match order ID %s with any order due to error %s", buyPriorityList[i].OrderID, err.Error())
				continue
			}
			if order1.CreatedAt < order2.CreatedAt {
				return []*models.Order{&order1}, []*models.Order{&order2}, nil
			} else {
				return []*models.Order{&order2}, []*models.Order{&order1}, nil
			}
		}
	}
	return nil, nil, fmt.Errorf("NOT MATCHED WITH ANY")
}

// MatchBatchDBOrders fetches buy orders with DB and matched them with proper sell orders
// returns two arrays of matched orders where first array contains buy orders and second contains sell orders
func MatchBatchDBOrders(database *db.DataBase, w *worker.Worker, maxFail int64) ([]*models.Order, []*models.Order, []string, error) {
	buyOrder := make([]*models.Order, 0)
	sellOrder := make([]*models.Order, 0)
	orderIDs := make([]string, 0)
	buyPriorityList, err := CreatePriorityList(database, false, "Price desc", w.ChainName)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("UNABLE TO GET ORDER FROM DB: %w", err)
	}
	w.Logger.Infof("Length of buy orders in DB %d", len(buyPriorityList))

	for i := 0; i < len(buyPriorityList); i++ {

		w.Logger.Infof("Working on buy order with price %f; trader: %s", buyPriorityList[i].Price, buyPriorityList[i].Trader)
		validated, err := w.ValidateOrder(&buyPriorityList[i])
		if err != nil {
			w.Logger.Errorf("Error in validaing limit order %s", err.Error())
			continue
		}
		if !validated {
			continue
		}
		if buyPriorityList[i].FailCount > maxFail {
			err = database.UpdateOrderStatus(buyPriorityList[i].OrderID, []models.MatchedStatus{}, []models.MatchedStatus{}, models.MatchedStatusBlocked)
			if err != nil {
				w.Logger.Error("Tried more that twice, can't process order")
				continue
			}
			w.Logger.Error("Tried more that twice, can't process order")
			continue
		}
		if buyPriorityList[i].Status == models.MatchedStatusSentFailed {
			if _, err := w.OrderValidation(buyPriorityList[i]); err != nil {
				if err := database.UpdateOrderStatusAndFailCount(buyPriorityList[i].OrderID, models.MatchedStatusBlocked); err != nil {
					return nil, nil, nil, fmt.Errorf("unable to updatestatus %e", err)
				}
			} else {
				if err := database.UpdateOrderStatusAndFailCount(buyPriorityList[i].OrderID, models.MatchedStatusFailedConfirmed); err != nil {
					return nil, nil, nil, fmt.Errorf("unable to updatestatus %e", err)
				}
				continue
			}

		}

		if buyPriorityList[i].OrderID != "" {
			order1, order2, err := MatchOrders(&buyPriorityList[i], database, maxFail, w)
			if err != nil {
				w.Logger.Infof("Unable to match order ID %s with any order due to error %s", buyPriorityList[i].OrderID, err.Error())
				continue
			}
			w.Logger.Infof("Matched orderID %s with orderID %s", order1.OrderID, order2.OrderID)
			orderIDs = append(orderIDs, order1.OrderID, order2.OrderID)
			if order1.CreatedAt < order2.CreatedAt {
				buyOrder = append(buyOrder, &order1)
				sellOrder = append(sellOrder, &order2)
			} else {
				buyOrder = append(buyOrder, &order2)
				sellOrder = append(sellOrder, &order1)
			}
		}
	}
	return buyOrder, sellOrder, orderIDs, nil
}

// MatchBatchSellDBOrders fetches buy orders with DB and matched them with proper sell orders
// returns two arrays of matched orders where first array contains buy orders and second contains sell orders
func MatchBatchSellDBOrders(database *db.DataBase, w *worker.Worker, maxFail int64) ([]*models.Order, []*models.Order, error) {
	buyOrder := make([]*models.Order, 0)
	sellOrder := make([]*models.Order, 0)
	sellPriorityList, err := CreatePriorityList(database, true, "Price", w.ChainName)
	if err != nil {
		return nil, nil, fmt.Errorf("UNABLE TO GET ORDER FROM DB: %w", err)
	}
	logrus.Infof("Length of buy orders in DB %d", len(sellPriorityList))
	for i := 0; i < len(sellPriorityList); i++ {

		logrus.Infof("Working on sell order with price %f; trader: %s", sellPriorityList[i].Price, sellPriorityList[i].Trader)
		logrus.Infof("Working on sell order with price %f; trader: %s", sellPriorityList[i].Price, sellPriorityList[i].Trader)
		validated, err := w.ValidateOrder(&sellPriorityList[i])
		if err != nil {
			w.Logger.Errorf("Error in validaing limit order %s", err.Error())
			continue
		}
		if !validated {
			logrus.Infof("Buy limit order price verification failed with  ID: %s trader: %s ", sellPriorityList[i].OrderID, sellPriorityList[i].Trader)
			continue
		}
		if sellPriorityList[i].FailCount > maxFail {
			err = database.UpdateOrderStatus(sellPriorityList[i].OrderID, []models.MatchedStatus{}, []models.MatchedStatus{}, models.MatchedStatusBlocked)
			if err != nil {
				logrus.Error("Tried more that twice, can't process order")
				continue
			}
			logrus.Error("Tried more that twice, can't process order")
			continue
		}
		if sellPriorityList[i].Status == models.MatchedStatusSentFailed {
			if _, err := w.OrderValidation(sellPriorityList[i]); err != nil {
				if err := database.UpdateOrderStatusAndFailCount(sellPriorityList[i].OrderID, models.MatchedStatusBlocked); err != nil {
					return nil, nil, fmt.Errorf("unable to updatestatus %e", err)
				}
			} else {
				if err := database.UpdateOrderStatusAndFailCount(sellPriorityList[i].OrderID, models.MatchedStatusFailedConfirmed); err != nil {
					return nil, nil, fmt.Errorf("unable to updatestatus %e", err)
				}
				continue
			}

		}

		if sellPriorityList[i].OrderID != "" {
			order1, order2, err := MatchOrders(&sellPriorityList[i], database, maxFail, w)
			if err != nil {
				logrus.Infof("Unable to match order ID %s with any order due to error %s", sellPriorityList[i].OrderID, err.Error())
				continue
			}
			if order1.CreatedAt < order2.CreatedAt {
				buyOrder = append(buyOrder, &order1)
				sellOrder = append(sellOrder, &order2)
			} else {
				buyOrder = append(buyOrder, &order2)
				sellOrder = append(sellOrder, &order1)
			}
		}
	}
	return buyOrder, sellOrder, nil
}

// MatchSellDBOrders fetch sell orders with DB and matched them with proper buy orders
// returns two arrays of matched orders where first array contains buy orders and second contains sell orders
// TODO: Match it as buyDBOrders
func MatchSellDBOrders(database *db.DataBase, maxFail int64, w *worker.Worker) ([]*models.Order, []*models.Order, error) {
	buyOrder := make([]*models.Order, 0)
	sellOrder := make([]*models.Order, 0)
	sellPriorityList, err := CreatePriorityList(database, false, "Price desc", w.ChainName)
	if err != nil {
		return nil, nil, fmt.Errorf("UNABLE TO GET ORDER FROM DB: %w", err)
	}
	logrus.Infof("Length of sell orders in DB %d", len(sellPriorityList))
	for i := 0; i < len(sellPriorityList); i++ {

		logrus.Infof("Working on sell order with price %f; trader: %s", sellPriorityList[i].Price, sellPriorityList[i].Trader)
		validated, err := w.ValidateOrder(&sellPriorityList[i])
		if err != nil {
			w.Logger.Errorf("Error in validaing limit order %s", err.Error())
			continue
		}
		if !validated {
			logrus.Infof("Buy limit order price verification failed with  ID: %s trader: %s ", sellPriorityList[i].OrderID, sellPriorityList[i].Trader)
			continue
		}
		if sellPriorityList[i].FailCount > maxFail {
			err = database.UpdateOrderStatus(sellPriorityList[i].OrderID, []models.MatchedStatus{}, []models.MatchedStatus{}, models.MatchedStatusBlocked)
			if err != nil {
				logrus.Error("Tried more that twice, can't process order")
				continue
			}
			logrus.Error("Tried more that twice, can't process order")
			continue
		}
		if sellPriorityList[i].OrderID != "" {
			order1, order2, err := MatchOrders(&sellPriorityList[i], database, maxFail, w)
			if err != nil {
				continue
			}
			if order1.CreatedAt < order2.CreatedAt {
				buyOrder = append(buyOrder, &order1)
				sellOrder = append(sellOrder, &order2)
			} else {
				buyOrder = append(buyOrder, &order2)
				sellOrder = append(sellOrder, &order1)
			}
		}
	}
	return buyOrder, sellOrder, nil
}

// VerifyMatchedOrder verifies two matched orders from database
func VerifyMatchedOrder(order1, order2 *models.Order, database *db.DataBase, maxFail int64, chainName string) error {
	if order1.FailCount > 0 {
		txnSents, err := database.GetTxnByOrderID(order1.OrderID)
		if err != nil {
			return err
		}
		skip := false
		for _, txnSent := range txnSents {
			for _, orderID := range txnSent.OrderID {
				if strings.EqualFold(orderID, order2.OrderID) {
					skip = true
					break
				}
			}
		}
		if skip {
			return fmt.Errorf("matched previously with same order")
		}
	}

	var matchedOrder []models.Order
	var err error
	if order1.IsShort {
		matchedOrder, err = GetPriorityOrderVerification(database, false, "Price Desc", chainName)
		if err != nil {
			return fmt.Errorf("VERIFICATION FAILED: Wrong match %w", err)
		}
	} else {
		matchedOrder, err = GetPriorityOrderVerification(database, true, "Price", chainName)
		if err != nil {
			return fmt.Errorf("VERIFICATION FAILED: GetPriorityOrderVerifications %w", err)
		}
	}

	//TODO: review this extensively
	for i := 0; i < len(matchedOrder); i++ {
		// if i > len(matchedOrder)-1 {
		// 	break
		// }
		if order2.OrderID == matchedOrder[i].OrderID {
			return nil
		}
	}
	return fmt.Errorf("VERIFICATION FAILED ORDER NOT MATCHED")
}

// func checkPriceStopLoss(order, matchingOrder models.Order, logger *logrus.Logger) {
// 	if order.IsShort && order.Price < matchingOrder.Price ||
// 		!order.IsShort && order.Price > matchingOrder.Price {
// 		logger.WithFields(logrus.Fields{
// 			"left_order":  order.OrderID,
// 			"right_order": matchingOrder.OrderID,
// 		}).Warnf("price is not suitable (%f, %f)", order.Price, matchingOrder.Price)
// 		return
// 	}
// }

// func checkPriceTakeProfit(order, matchingOrder models.Order, logger *logrus.Logger) {
// 	if order.IsShort && order.Price > matchingOrder.Price ||
// 		!order.IsShort && order.Price < matchingOrder.Price {
// 		logger.WithFields(logrus.Fields{
// 			"left_order":  order.OrderID,
// 			"right_order": matchingOrder.OrderID,
// 		}).Warnf("price is not suitable (%f, %f)", order.Price, matchingOrder.Price)
// 		return
// 	}
// }
