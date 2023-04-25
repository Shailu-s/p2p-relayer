package matching_engine

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/volmexfinance/relayers/relayer-srv/big_ext"

	"github.com/volmexfinance/relayers/relayer-srv/db"
)

func matchAndUpdateOrders(orderLeft, orderRight *db.Order) error {
	orderLeftFill := orderLeft.OrderFills()
	orderRightFill := orderRight.OrderFills()
	println(orderLeft.OrderFills().String(), "===============1", orderRight.OrderFills().String())
	newOrderLeftFill, newOrderRightFill, err := fillOrder(orderLeft, orderRight, orderLeftFill, orderRightFill)
	if err != nil {
		return fmt.Errorf("matchAndUpdateOrders: unable to fill order. %w", err)
	}
	println(orderLeft.OrderFills().String(), "===============2", orderRight.OrderFills().String(), newOrderLeftFill.String(), newOrderRightFill.String())

	// New fill value should be greater than 0 at all times
	if big_ext.LessThanOrEqual(newOrderLeftFill, big.NewInt(0)) || big_ext.LessThanOrEqual(newOrderRightFill, big.NewInt(0)) {
		return errors.New("matchAndUpdateOrders: fill should be greater than 0")
	}

	if err := updateOrderFillValue(orderLeft, newOrderLeftFill); err != nil {
		return err
	}
	if err := updateOrderFillValue(orderRight, newOrderRightFill); err != nil {
		return err
	}
	if err := updateOrderStatus(orderLeft); err != nil {
		return fmt.Errorf("updateOrderStatus: left order failed %w", err)
	}
	if err := updateOrderStatus(orderRight); err != nil {
		return fmt.Errorf("updateOrderStatus: right order failed %w", err)
	}

	return nil
}

func updateOrderFillValue(order *db.Order, newFill *big.Int) error {
	currentFill := order.OrderFills()
	println("currentfill", currentFill.String(), "nfill", newFill.String())
	currentFill = currentFill.Add(currentFill, newFill)
	if big_ext.GreaterThan(currentFill, order.MakeAsset().ValueAsBigInt()) {
		return fmt.Errorf("new fill is greter than matched order %s", currentFill.String())
	}
	order.SetOrderFills(currentFill)
	return nil
}

func updateOrderStatus(order *db.Order) error {
	makeAsset := order.MakeAsset().ValueAsBigInt()
	fills := order.OrderFills()
	println(fills.String(), "fillllllllllllllllllllllllls", makeAsset.String())
	if big_ext.Equals(fills, makeAsset) {
		order.Status = db.MatchedStatusFullMatchFound
	} else if big_ext.LessThanOrEqual(fills, makeAsset) {
		order.Status = db.MatchedStatusPartialMatchFound
	} else {
		return fmt.Errorf("order fill is greater than make asset")
	}

	return nil
}

func fillOrder(leftOrder, rightOrder *db.Order, leftOrderFill, rightOrderFill *big.Int) (*big.Int, *big.Int, error) {
	leftBaseValue, leftQuoteValue, err := calculateRemaining(leftOrder, leftOrderFill)
	if err != nil {
		return nil, nil, fmt.Errorf("fillOrder: unable to calculate remaining fill %w", err)
	}
	rightBaseValue, rightQuoteValue, err := calculateRemaining(rightOrder, rightOrderFill)
	if err != nil {
		return nil, nil, fmt.Errorf("fillOrder: unable to calculate remaining fill %w", err)
	}

	// Perform order fills calculation
	var newLeftOrderFill, newRightOrderFill *big.Int

	rightMakeValue := rightOrder.MakeAsset().ValueAsBigInt()
	rightTakeValue := rightOrder.TakeAsset().ValueAsBigInt()
	leftMakeValue := leftOrder.MakeAsset().ValueAsBigInt()
	leftTakeValue := leftOrder.TakeAsset().ValueAsBigInt()

	// if rightQuoteValue > leftBaseValue, fill left order
	if big_ext.GreaterThan(rightQuoteValue, leftBaseValue) {
		newLeftOrderFill, newRightOrderFill, err = fillLeft(leftBaseValue, leftQuoteValue, rightMakeValue, rightTakeValue)
		if err != nil {
			return nil, nil, fmt.Errorf("fillOrder: unable to filling orderleft %w", err)
		}

		return newLeftOrderFill, newRightOrderFill, nil
	} else {

		// else fill right order
		newLeftOrderFill, newRightOrderFill, err = fillRight(leftMakeValue, leftTakeValue, rightBaseValue, rightQuoteValue)
		if err != nil {
			newLeftOrderFill, newRightOrderFill, err = fillLeft(leftBaseValue, leftQuoteValue, rightMakeValue, rightTakeValue)
			if err != nil {
				return nil, nil, fmt.Errorf("fillOrder: unable to filling orderRight %w", err)
			}

			return newLeftOrderFill, newRightOrderFill, nil
		}

		return newLeftOrderFill, newRightOrderFill, nil
	}
}

func calculateRemaining(order *db.Order, fill *big.Int) (*big.Int, *big.Int, error) {
	makeValue := order.MakeAsset().ValueAsBigInt()
	takeValue := order.TakeAsset().ValueAsBigInt()

	// baseValue = makeValue - fill
	baseValue := new(big.Int).Sub(makeValue, fill)

	quoteValue, err := safeGetPartialAmountFloor(takeValue, makeValue, baseValue)
	if err != nil {
		return nil, nil, fmt.Errorf("CALCULATEREMAINING:UNABLE TO GET  SAFEGETPARTIALAMOUNTFLOOR %w", err)
	}

	return baseValue, quoteValue, nil
}

// safeGetPartialAmountFloor computes the result of multiplying numerator by target and dividing by denominator,
// rounding down to the nearest integer. This function returns an error if the rounding error is greater than 0.1%.
func safeGetPartialAmountFloor(numerator, denominator, target *big.Int) (*big.Int, error) {
	if err := isRoundingErrorFloor(numerator, denominator, target); err != nil {
		return nil, fmt.Errorf("safeGetPartialAmountFloor: %w", err)
	}

	// partialAmount = (numerator * target) / denominator
	partialAmount := new(big.Int).Mul(numerator, target)
	partialAmount = partialAmount.Div(partialAmount, denominator)

	return partialAmount, nil
}

// isRoundingErrorFloor checks if dividing numerator by denominator and rounding down to the nearest integer
// is equal to target, with an error margin of 0.1%. If the rounding error is greater than 0.1%, an error is returned.
func isRoundingErrorFloor(numerator, denominator, target *big.Int) error {
	if big_ext.Equals(denominator, big.NewInt(0)) {
		return errors.New("isRoundingErrorFloor: division by zero")
	}

	if big_ext.Equals(target, big.NewInt(0)) || big_ext.Equals(numerator, big.NewInt(0)) {
		return nil
	}

	// remainder = (numerator * target) % denominator
	remainder := new(big.Int).Mul(numerator, target)
	remainder = remainder.Mod(remainder, denominator)

	// (remainder * 1000) >= (numerator * target)
	if big_ext.GreaterThanOrEqual(new(big.Int).Mul(remainder, big.NewInt(1000)), new(big.Int).Mul(numerator, target)) {
		return fmt.Errorf("isRoundingErrorFloor: rounding error")
	}

	return nil
}

func fillRight(leftMakeValue, leftTakeValue, rightMakeValue, rightTakeValue *big.Int) (*big.Int, *big.Int, error) {
	makerValue, err := safeGetPartialAmountFloor(rightTakeValue, leftMakeValue, leftTakeValue)
	if err != nil {
		return nil, nil, fmt.Errorf("fillRight: unable to get safeGetPartialAmountFloor %w", err)
	}

	if big_ext.GreaterThan(makerValue, rightMakeValue) {
		return nil, nil, errors.New("fillRight: Unable to fill")
	}

	return rightTakeValue, makerValue, nil
}

func fillLeft(leftMakeValue, leftTakeValue, rightMakeValue, rightTakeValue *big.Int) (*big.Int, *big.Int, error) {
	rightTake, err := safeGetPartialAmountFloor(leftTakeValue, rightMakeValue, rightTakeValue)
	if err != nil {
		return nil, nil, fmt.Errorf("fillLeft: unable to get safeGetPartialAmountFloor %w", err)
	}

	if big_ext.GreaterThan(rightTake, leftMakeValue) {
		return nil, nil, errors.New("fillLeft: unable to fill")
	}

	return leftMakeValue, leftTakeValue, nil
}
