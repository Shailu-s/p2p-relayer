package postgresdb

import (
	"github.com/volmexfinance/relayers/relayer-srv/db/models"
)

func (db *PostgresDataBase) CreateOrderBatch(newOrder []*models.Order) error {
	if err := db.DB.Create(newOrder); err.Error != nil {
		return err.Error
	}
	return nil
}
