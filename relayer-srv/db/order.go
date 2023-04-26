package db

import (
	"errors"
	"fmt"
	"github.com/volmexfinance/relayers/relayer-srv/db/models"
	"gorm.io/gorm"
	"time"

	"github.com/sirupsen/logrus"
)

// CreateOrder creates an order in the DB
func (db *DataBase) CreateOrder(order *models.Order) error {
	order.CreatedAt = time.Now().Unix()
	order.UpdatedAt = time.Now().Unix()

	if order.Status == 0 {
		order.Status = models.MatchedStatusInit
	}

	if err := db.DB.Create(&order); err.Error != nil {
		return err.Error
	}
	return nil
}

// CreateOrderInBatch creates multiple orders in the DB
func (db *DataBase) CreateOrderInBatch(orders []*models.Order) error {
	result := db.DB.Create(&orders)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetZeroOrders returns orders from DB with status zero
func (db *DataBase) GetZeroOrders(chain string) ([]*models.Order, error) {
	var orders []*models.Order
	result := db.DB.
		Model(models.Order{}).
		Where(
			"chain_name = ? AND status in (?)",
			chain, []models.MatchedStatus{models.MatchedStatusZero},
		).
		Preload("Assets").
		Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

// GetOrdersOnStatus returns orders from DB with given statuses
func (db *DataBase) GetOrdersOnStatus(chain string, statuses []models.MatchedStatus) ([]*models.Order, error) {
	var orders []*models.Order
	result := db.DB.
		Model(models.Order{}).
		Where(
			"chain_name = ? AND status in (?)",
			chain, statuses,
		).
		Preload("Assets").
		Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

// DeleteBatchOrder deletes order from DB
func (db *DataBase) DeleteBatchOrder(order []*models.Order) error {
	result := db.DB.Delete(&order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// UpdateOrder take an updated order and update it
func (db *DataBase) UpdateOrder(order *models.Order) error {
	order.UpdatedAt = time.Now().Unix()

	result := db.DB.
		Model(models.Order{}).
		Where("order_id = ?", order.OrderID).
		Updates(order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// UpdateOrCreateOrder updates the order or creates a new one if it doesn't exist
func (db *DataBase) UpdateOrCreateOrder(order *models.Order) error {
	result, err := db.FindOrder(order.OrderID)
	if err != nil {
		return err
	}

	if result == nil {
		if err = db.CreateOrder(order); err != nil {
			return err
		}
	}

	if err = db.UpdateOrder(order); err != nil {
		return err
	}

	return nil
}

// UpdateOrderStatusAndFailCount updates order status and its fail count
func (db *DataBase) UpdateOrderStatusAndFailCount(orderID string, status models.MatchedStatus) error {
	order, err := db.FindOrder(orderID)
	if err != nil {
		return errors.New("Order not Found")
	}

	order.UpdatedAt = time.Now().Unix()
	order.Status = status
	order.FailCount = order.FailCount + 1

	if result := db.DB.Save(&order); result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateOrderStatusByMinSalt updates order status with condition
func (db *DataBase) UpdateOrderStatusByMinSalt(trader string, minSalt string, inStatuses, notInStatuses []models.MatchedStatus, updateStatus models.MatchedStatus, chain string) error {
	query := db.DB.
		Model(models.Order{}).
		Where(
			"chain_name = ? AND trader = ? AND salt <= ?",
			chain, trader, minSalt,
		)

	if len(inStatuses) != 0 {
		query = query.Where("status in (?)", inStatuses)
	}

	if len(notInStatuses) != 0 {
		query = query.Where("status not in (?)", notInStatuses)
	}

	toUpdate := map[string]interface{}{
		"status":     updateStatus,
		"updated_at": time.Now().Unix(),
	}
	return query.Updates(toUpdate).Error
}

// UpdateBatchOrderStatus updates status of order in batch
// TODO: Don't use for loop to batch update order
func (db *DataBase) UpdateBatchOrderStatus(newOrder []*models.Order, newStatus models.MatchedStatus) error {
	for _, order := range newOrder {
		order.UpdatedAt = time.Now().Unix()
		order.Status = newStatus

		result := db.DB.
			Model(models.Order{}).
			Where("order_id = ?", order.OrderID).
			Updates(order)

		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

// UpdateFillAndStatusByTxnLog updates status of order in batch and fills by transaction log
func (db *DataBase) UpdateFillAndStatusByTxnLog(newOrder []*models.Order, newStatus models.MatchedStatus) error {
	// TODO: Don't use for loop to batch update order
	for _, order := range newOrder {
		order.UpdatedAt = time.Now().Unix()
		order.Status = newStatus
		order.Fills = ""

		var txnLog models.TransactionLog
		result := db.DB.
			Model(models.TransactionLog{}).
			Where("order_id LIKE ?", fmt.Sprintf("%%%s%%", order.OrderID)).
			Last(&txnLog)
		if result.Error != nil {
			logrus.Infof("No txn log history found in DB")
		} else {
			if txnLog.OrderID[0] == order.OrderID {
				order.Fills = txnLog.NewLeftFill
			} else {
				order.Fills = txnLog.NewRightFill
			}
		}

		if result := db.DB.Save(&order); result.Error != nil {
			return result.Error
		}
	}
	return nil
}

// FindOrder return order from DB with given order ID
func (db *DataBase) FindOrder(orderID string) (*models.Order, error) {
	var order models.Order
	result := db.DB.
		Model(models.Order{}).
		Where("order_id = ?", orderID).
		Preload("Assets").
		Limit(1).
		Find(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

// FindOrderStatus returns the status of given order ID in DB
func (db *DataBase) FindOrderStatus(orderID string) (models.MatchedStatus, error) {
	var order models.Order

	result := db.DB.
		Model(models.Order{}).
		Where("order_id = ?", orderID).
		Limit(1).
		Find(&order)
	if result.Error != nil {
		return 0, result.Error
	}

	return order.Status, nil
}

// FindFulfilledFailedOrder return order with given status
func (db *DataBase) FindFulfilledFailedOrder(statuses []models.MatchedStatus, chain string) ([]*models.Order, error) {
	var orders []*models.Order

	result := db.DB.
		Model(models.Order{}).
		Where(
			"chain_name = ? AND status in (?)",
			chain, statuses,
		).
		Preload("Assets").
		Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

// GetOrdersSortedByPriority returns orders sorted by priority
func (db *DataBase) GetOrdersSortedByPriority(isShort bool, price string, chain string) ([]models.Order, error) {
	var PriorityList []models.Order

	result := db.DB.
		Model(&models.Order{}).
		Where(
			"chain_name = ? AND is_short = ? AND status IN (?)",
			chain,
			isShort,
			[]models.MatchedStatus{
				models.MatchedStatusInit,
				models.MatchedStatusPartialMatchFound, //TODO: this would NOT allow same order to be again matched in same tx
				models.MatchedStatusPartialMatchConfirmed,
				models.MatchedStatusFailedConfirmed, //TODO: once failed confirmed where should it go?
			},
		).
		Preload("Assets").
		Order(price).
		Find(&PriorityList)
	if result.Error != nil {
		return nil, fmt.Errorf("CreatePriorityList: unable to get list from db %w", result.Error)
	}

	return PriorityList, nil
}

// GetOrdersSortedByPriorityForVerification returns orders sorted by priority
func (db *DataBase) GetOrdersSortedByPriorityForVerification(isShort bool, price string, chain string) ([]models.Order, error) {
	var PriorityOrder []models.Order

	result := db.DB.
		Model(&models.Order{}).
		Where(
			"chain_name = ? AND is_short = ? AND status IN (?)",
			chain,
			isShort,
			[]models.MatchedStatus{
				models.MatchedStatusInit,
				models.MatchedStatusPartialMatchConfirmed,
				models.MatchedStatusFullMatchFound,
				models.MatchedStatusPartialMatchFound,
				models.MatchedStatusValidated,
				models.MatchedStatusValidationConfirmed,
				models.MatchedStatusFailedConfirmed,
			}).
		Preload("Assets").
		Order(price).
		Find(&PriorityOrder)
	if result.Error != nil {
		return []models.Order{}, fmt.Errorf("GetPriorityOrderVerification: unable to get order from db %w", result.Error)
	}

	return PriorityOrder, nil
}

// HandleOrderStatusAndFills updates status and fills of order on given statuses
func (db *DataBase) HandleOrderStatusAndFills(orderID string, fills string, inStatuses, notInStatuses []models.MatchedStatus, updateStatus models.MatchedStatus) error {
	query := db.DB.Model(models.Order{}).Where("order_id = ?", orderID)
	if len(inStatuses) != 0 {
		query = query.Where("status in (?)", inStatuses)
	}
	if len(notInStatuses) != 0 {
		query = query.Where("status not in (?)", notInStatuses)
	}

	toUpdate := map[string]interface{}{
		"status":     updateStatus,
		"updated_at": time.Now().Unix(),
		"fills":      fills,
	}
	return query.Updates(toUpdate).Error
}

// HasBaseToken returns true if there is an order with given base token
func (db *DataBase) HasBaseToken(baseToken string, chain string) (bool, error) {
	var order models.Order
	result := db.DB.
		Joins("JOIN assets ON assets.orderbook_id = orders.order_id AND assets.virtual_token = ?", baseToken).
		Where("chain_name = ?", chain).
		First(&order)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return false, result.Error
		}
	}

	return true, nil
}

// DepthOrderDetails returns order details with given base token and status
func (db *DataBase) DepthOrderDetails(baseToken string, inStatuses []models.MatchedStatus, chain string) ([]*models.Order, error) {
	var orders []*models.Order
	result := db.DB.
		Joins("JOIN assets ON assets.orderbook_id = orders.order_id AND assets.virtual_token = ?", baseToken).
		Where("chain_name = ? AND status in (?)", chain, inStatuses).
		Preload("Assets").
		Order("price").
		Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

// UpdateOrderStatus updates status of order on given statuses
func (db *DataBase) UpdateOrderStatus(orderID string, inStatuses, notInStatuses []models.MatchedStatus, updateStatus models.MatchedStatus) error {
	query := db.DB.Model(models.Order{}).Where("order_id = ?", orderID)
	if len(inStatuses) != 0 {
		query = query.Where("status in (?)", inStatuses)
	}
	if len(notInStatuses) != 0 {
		query = query.Where("status not in (?)", notInStatuses)
	}

	toUpdate := map[string]interface{}{
		"status":     updateStatus,
		"updated_at": time.Now().Unix(),
	}
	return query.Updates(toUpdate).Error
}

// UpdateDeadlinePassedOrder updates the status of and order that has passed deadline
func (db *DataBase) UpdateDeadlinePassedOrder(currentTime uint64, notInStatuses []models.MatchedStatus, chain string) error {
	query := db.DB.Model(models.Order{}).Where("deadline < ?", currentTime)

	if len(notInStatuses) != 0 {
		query = query.Where(
			"chain_name = ? AND status not in (?)",
			chain, notInStatuses,
		)
	}

	toUpdate := map[string]interface{}{
		"status":     models.MatchedStatusBlocked,
		"updated_at": time.Now().Unix(),
	}
	return query.Updates(toUpdate).Error
}

// GetAllOrdersByTraderWithoutSign return order without trader sign, used for API
func (db *DataBase) GetAllOrdersByTraderWithoutSign(trader string, chain string) ([]*models.Order, error) {
	var orders []*models.Order

	result := db.DB.
		Model(models.Order{}).
		Select("order_id", "order_type", "trader", "deadline", "is_short", "salt", "trigger_price", "fills", "created_at", "price", "status").
		Where(
			"chain_name = ? AND trader = ?",
			chain, trader,
		).
		Order("created_at desc").
		Preload("Assets").
		Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

// GetOrderQueue return order queue, used for API
func (db *DataBase) GetOrderQueue(chain string) ([]*models.Order, error) {
	var orders []*models.Order

	result := db.DB.
		Model(models.Order{}).
		Select("order_id", "is_short", "trigger_price", "created_at", "status").
		Where(
			"chain_name = ? AND status in (?)",
			chain,
			[]models.MatchedStatus{
				models.MatchedStatusInit,
				models.MatchedStatusPartialMatchConfirmed,
				models.MatchedStatusFailedConfirmed,
			},
		).
		Order("created_at desc").
		Preload("Assets").
		Limit(100).
		Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

// RemoveOrder deletes order from DB
func (db *DataBase) RemoveOrder(orderID string, status []models.MatchedStatus) error {
	result := db.DB.
		Model(models.Order{}).
		Where(
			"order_id = ? AND status in ?",
			orderID, status,
		).
		Delete(models.Order{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
