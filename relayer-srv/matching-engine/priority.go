package matching_engine

import (
	"github.com/volmexfinance/relayers/relayer-srv/db"
)

// CreatePriorityList creates list of order in priority to match them
func CreatePriorityList(database *db.DataBase, isShort bool, price string, chain string) ([]db.Order, error) {
	return database.GetOrdersSortedByPriority(isShort, price, chain)
}

// GetPriorityOrderVerification returns priority queue of order for verification
func GetPriorityOrderVerification(database *db.DataBase, isShort bool, price string, chain string) ([]db.Order, error) {
	return database.GetOrdersSortedByPriorityForVerification(isShort, price, chain)
}
