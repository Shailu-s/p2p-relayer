package relayer_srv

import (
	"github.com/volmexfinance/relayers/relayer-srv/db"
)

func (r *RelayerSrv) InsertOrder(order *db.Order) error {
	return r.db.CreateOrder(order)
}

func (r *RelayerSrv) GetAllOrders(trader string, chainName string) ([]*db.Order, error) {
	return r.db.GetAllOrdersByTraderWithoutSign(trader, chainName)
}

func (r *RelayerSrv) GetOrderQueue(chain string) ([]*db.Order, error) {
	return r.db.GetOrderQueue(chain)
}

func (r *RelayerSrv) CheckHasBaseToken(token, chain string) (bool, error) {
	return r.db.HasBaseToken(token, chain)
}

func (r *RelayerSrv) GetDepthOrderDetails(token, chain string) ([]*db.Order, error) {
	return r.db.DepthOrderDetails(
		token,
		[]db.MatchedStatus{
			db.MatchedStatusInit,
			db.MatchedStatusPartialMatchConfirmed,
			db.MatchedStatusFailedConfirmed,
		},
		chain,
	)
}
