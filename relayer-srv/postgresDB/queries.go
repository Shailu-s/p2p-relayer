package postgresdb

import (
	storage "github.com/volmexfinance/relayers/relayer-srv/db"
)

func (db *PostgresDataBase) CreateOrderBatch(newOrder []*storage.Order) error {
	if err := db.DB.Create(newOrder); err.Error != nil {
		return err.Error
	}
	return nil
}
