package db

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/sirupsen/logrus"
)

// CreateOrder creates an order in the DB
func (db *SQLiteDataBase) CreateOrder(order *Order) error {
	order.CreatedAt = time.Now().Unix()
	order.UpdatedAt = time.Now().Unix()

	if order.Status == 0 {
		order.Status = MatchedStatusInit
	}

	if err := db.DB.Create(&order); err.Error != nil {
		return err.Error
	}
	return nil
}

// CreateOrder creates an order in the DB
func (db *PostgresDataBase) CreateOrder(order *Order) error {
	order.CreatedAt = time.Now().Unix()
	order.UpdatedAt = time.Now().Unix()

	if order.Status == 0 {
		order.Status = MatchedStatusInit
	}

	if err := db.DB.Create(&order); err.Error != nil {
		return err.Error
	}
	return nil
}

// CreateOrderInBatch creates multiple orders in the DB
func (db *SQLiteDataBase) CreateOrderInBatch(orders []*Order) error {
	result := db.DB.Create(&orders)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetZeroOrders returns orders from DB with status zero
func (db *SQLiteDataBase) GetZeroOrders(chain string) ([]*Order, error) {
	var orders []*Order
	result := db.DB.
		Model(Order{}).
		Where(
			"chain_name = ? AND status in (?)",
			chain, []MatchedStatus{MatchedStatusZero},
		).
		Preload("Assets").
		Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

// GetZeroOrders returns orders from DB with status zero
func (db *PostgresDataBase) GetZeroOrders(chain string) ([]*Order, error) {
	var orders []*Order
	result := db.DB.
		Model(Order{}).
		Where(
			"chain_name = ? AND status in (?)",
			chain, []MatchedStatus{MatchedStatusZero},
		).
		Preload("Assets").
		Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

// GetOrdersOnStatus returns orders from DB with given statuses
func (db *SQLiteDataBase) GetOrdersOnStatus(chain string, statuses []MatchedStatus) ([]*Order, error) {
	var orders []*Order
	result := db.DB.
		Model(Order{}).
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

// GetOrdersOnStatus returns orders from DB with given statuses
func (db *PostgresDataBase) GetOrdersOnStatus(chain string, statuses []MatchedStatus) ([]*Order, error) {
	var orders []*Order
	result := db.DB.
		Model(Order{}).
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
func (db *SQLiteDataBase) DeleteBatchOrder(order []*Order) error {
	result := db.DB.Delete(&order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// UpdateOrder take an updated order and update it
func (db *SQLiteDataBase) UpdateOrder(order *Order) error {
	order.UpdatedAt = time.Now().Unix()

	result := db.DB.
		Model(Order{}).
		Where("order_id = ?", order.OrderID).
		Updates(order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// UpdateOrCreateOrder updates the order or creates a new one if it doesn't exist
func (db *SQLiteDataBase) UpdateOrCreateOrder(order *Order) error {
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
func (db *SQLiteDataBase) UpdateOrderStatusAndFailCount(orderID string, status MatchedStatus) error {
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

// UpdateOrderStatusAndFailCount updates order status and its fail count
func (db *PostgresDataBase) UpdateOrderStatusAndFailCount(orderID string, status MatchedStatus) error {
	order, err := db.FindOrder(orderID)
	if err != nil {
		return errors.New("Order not Found")
	}

	order.UpdatedAt = time.Now().Unix()
	order.Status = status
	order.FailCount = order.FailCount + 1
	order.Fills = ""

	if result := db.DB.Save(&order); result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateOrderStatusByMinSalt updates order status with condition
func (db *PostgresDataBase) UpdateOrderStatusByMinSalt(trader string, minSalt string, inStatuses, notInStatuses []MatchedStatus, updateStatus MatchedStatus, chain string) error {
	query := db.DB.
		Model(Order{}).
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

// UpdateOrderStatusByMinSalt updates order status with condition
func (db *SQLiteDataBase) UpdateOrderStatusByMinSalt(trader string, minSalt string, inStatuses, notInStatuses []MatchedStatus, updateStatus MatchedStatus, chain string) error {
	query := db.DB.
		Model(Order{}).
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
func (db *PostgresDataBase) UpdateBatchOrderStatus(newOrder []*Order, newStatus MatchedStatus) error {
	for _, order := range newOrder {
		order.UpdatedAt = time.Now().Unix()
		order.Status = newStatus

		result := db.DB.
			Model(Order{}).
			Where("order_id = ?", order.OrderID).
			Updates(order)

		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

// UpdateBatchOrderStatus updates status of order in batch
// TODO: Don't use for loop to batch update order
func (db *SQLiteDataBase) UpdateBatchOrderStatus(newOrder []*Order, newStatus MatchedStatus) error {
	for _, order := range newOrder {
		order.UpdatedAt = time.Now().Unix()
		order.Status = newStatus

		result := db.DB.
			Model(Order{}).
			Where("order_id = ?", order.OrderID).
			Updates(order)

		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

// UpdateFillAndStatusByTxnLog updates status of order in batch and fills by transaction log
func (db *PostgresDataBase) UpdateFillAndStatusByTxnLog(newOrder []*Order, newStatus MatchedStatus) error {
	// TODO: Don't use for loop to batch update order
	for _, order := range newOrder {
		query := db.DB.Model(Order{}).Where("order_id = ?", order.OrderID)

		var txnLog TransactionLog
		fills := ""
		result := db.DB.
			Model(TransactionLog{}).
			Where("order_id LIKE ?", fmt.Sprintf("%%%s%%", order.OrderID)).
			Last(&txnLog)
		if result.Error != nil {
			logrus.Infof("No txn log history found in DB")
		} else {
			if txnLog.OrderID[0] == order.OrderID {
				fills = txnLog.NewLeftFill
			} else {
				fills = txnLog.NewRightFill
			}
		}

		toUpdate := map[string]interface{}{
			"status":     newStatus,
			"updated_at": time.Now().Unix(),
			"fills":      fills,
		}
		if err := query.Updates(toUpdate).Error; err != nil {
			return err
		}
	}
	return nil
}

// UpdateFillAndStatusByTxnLog updates status of order in batch and fills by transaction log
func (db *SQLiteDataBase) UpdateFillAndStatusByTxnLog(newOrder []*Order, newStatus MatchedStatus) error {
	// TODO: Don't use for loop to batch update order
	for _, order := range newOrder {
		order.UpdatedAt = time.Now().Unix()
		order.Status = newStatus
		order.Fills = ""

		var txnLog TransactionLog
		result := db.DB.
			Model(TransactionLog{}).
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
func (db *SQLiteDataBase) FindOrder(orderID string) (*Order, error) {
	var order Order
	result := db.DB.
		Model(Order{}).
		Where("order_id = ?", orderID).
		Preload("Assets").
		Limit(1).
		Find(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

// FindOrder return order from DB with given order ID
func (db *PostgresDataBase) FindOrder(orderID string) (*Order, error) {
	var order Order
	result := db.DB.
		Model(Order{}).
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
func (db *SQLiteDataBase) FindOrderStatus(orderID string) (MatchedStatus, error) {
	var order Order

	result := db.DB.
		Model(Order{}).
		Where("order_id = ?", orderID).
		Limit(1).
		Find(&order)
	if result.Error != nil {
		return 0, result.Error
	}

	return order.Status, nil
}

// FindFulfilledFailedOrder return order with given status
func (db *SQLiteDataBase) FindFulfilledFailedOrder(statuses []MatchedStatus, chain string) ([]*Order, error) {
	var orders []*Order

	result := db.DB.
		Model(Order{}).
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
func (db *SQLiteDataBase) GetOrdersSortedByPriority(isShort bool, price string, chain string) ([]Order, error) {
	var PriorityList []Order

	result := db.DB.
		Model(&Order{}).
		Where(
			"chain_name = ? AND is_short = ? AND status IN (?)",
			chain,
			isShort,
			[]MatchedStatus{
				MatchedStatusZero,
				MatchedStatusInit,
				MatchedStatusPartialMatchFound, //TODO: this would NOT allow same order to be again matched in same tx
				MatchedStatusPartialMatchConfirmed,
				MatchedStatusFailedConfirmed, //TODO: once failed confirmed where should it go?
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
func (db *SQLiteDataBase) GetOrdersSortedByPriorityForVerification(isShort bool, price string, chain string) ([]Order, error) {
	var PriorityOrder []Order

	result := db.DB.
		Model(&Order{}).
		Where(
			"chain_name = ? AND is_short = ? AND status IN (?)",
			chain,
			isShort,
			[]MatchedStatus{
				MatchedStatusInit,
				MatchedStatusPartialMatchConfirmed,
				MatchedStatusFullMatchFound,
				MatchedStatusPartialMatchFound,
				MatchedStatusValidated,
				MatchedStatusValidationConfirmed,
				MatchedStatusFailedConfirmed,
			}).
		Preload("Assets").
		Order(price).
		Find(&PriorityOrder)
	if result.Error != nil {
		return []Order{}, fmt.Errorf("GetPriorityOrderVerification: unable to get order from db %w", result.Error)
	}

	return PriorityOrder, nil
}

// HandleOrderStatusAndFills updates status and fills of order on given statuses
func (db *PostgresDataBase) HandleOrderStatusAndFills(orderID string, fills string, inStatuses, notInStatuses []MatchedStatus, updateStatus MatchedStatus) error {
	query := db.DB.Model(Order{}).Where("order_id = ?", orderID)
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

// HandleOrderStatusAndFills updates status and fills of order on given statuses
func (db *SQLiteDataBase) HandleOrderStatusAndFills(orderID string, fills string, inStatuses, notInStatuses []MatchedStatus, updateStatus MatchedStatus) error {
	query := db.DB.Model(Order{}).Where("order_id = ?", orderID)
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
func (db *PostgresDataBase) HasBaseToken(baseToken string, chain string) (bool, error) {
	var order Order
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
func (db *PostgresDataBase) DepthOrderDetails(baseToken string, inStatuses []MatchedStatus, chain string) ([]*Order, error) {
	var orders []*Order
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
func (db *SQLiteDataBase) UpdateOrderStatus(orderID string, inStatuses, notInStatuses []MatchedStatus, updateStatus MatchedStatus) error {
	query := db.DB.Model(Order{}).Where("order_id = ?", orderID)
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

// UpdateOrderStatus updates status of order on given statuses
func (db *PostgresDataBase) UpdateOrderStatus(orderID string, inStatuses, notInStatuses []MatchedStatus, updateStatus MatchedStatus) error {
	query := db.DB.Model(Order{}).Where("order_id = ?", orderID)
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
func (db *SQLiteDataBase) UpdateDeadlinePassedOrder(currentTime uint64, notInStatuses []MatchedStatus, chain string) error {
	query := db.DB.Model(Order{}).Where("deadline < ?", currentTime)

	if len(notInStatuses) != 0 {
		query = query.Where(
			"chain_name = ? AND status not in (?)",
			chain, notInStatuses,
		)
	}

	toUpdate := map[string]interface{}{
		"status":     MatchedStatusBlocked,
		"updated_at": time.Now().Unix(),
	}
	return query.Updates(toUpdate).Error
}

// GetAllOrdersByTraderWithoutSign return order without trader sign, used for API
func (db *PostgresDataBase) GetAllOrdersByTraderWithoutSign(trader string, chain string) ([]*Order, error) {
	var orders []*Order

	result := db.DB.
		Model(Order{}).
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
func (db *PostgresDataBase) GetOrderQueue(chain string) ([]*Order, error) {
	var orders []*Order

	result := db.DB.
		Model(Order{}).
		Select("order_id", "is_short", "trigger_price", "created_at", "status").
		Where(
			"chain_name = ? AND status in (?)",
			chain,
			[]MatchedStatus{
				MatchedStatusInit,
				MatchedStatusPartialMatchConfirmed,
				MatchedStatusFailedConfirmed,
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
func (db *SQLiteDataBase) RemoveOrder(orderID string, status []MatchedStatus) error {
	result := db.DB.
		Model(Order{}).
		Where(
			"order_id = ? AND status in ?",
			orderID, status,
		).
		Delete(Order{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
