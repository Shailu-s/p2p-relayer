package relayer_srv

import (
	"github.com/volmexfinance/relayers/relayer-srv/db"
)

func (r *RelayerSrv) InsertOrder(order *db.Order) error {
	return r.postgresDB.CreateOrder(order)
}

func (r *RelayerSrv) GetAllOrders(trader string, chainName string) ([]*db.Order, error) {
	return r.postgresDB.GetAllOrdersByTraderWithoutSign(trader, chainName)
}

func (r *RelayerSrv) GetOrderQueue(chain string) ([]*db.Order, error) {
	return r.postgresDB.GetOrderQueue(chain)
}

func (r *RelayerSrv) CheckHasBaseToken(token, chain string) (bool, error) {
	return r.postgresDB.HasBaseToken(token, chain)
}

func (r *RelayerSrv) GetDepthOrderDetails(token, chain string) ([]*db.Order, error) {
	return r.postgresDB.DepthOrderDetails(
		token,
		[]db.MatchedStatus{
			db.MatchedStatusInit,
			db.MatchedStatusPartialMatchConfirmed,
			db.MatchedStatusFailedConfirmed,
		},
		chain,
	)
}
