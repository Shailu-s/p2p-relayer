package relayer_srv

import (
	"github.com/volmexfinance/relayers/relayer-srv/db/models"
)

func (r *RelayerSrv) InsertOrder(order *models.Order) error {
	return r.db.CreateOrder(order)
}

func (r *RelayerSrv) GetAllOrders(trader string, chainName string) ([]*models.Order, error) {
	return r.db.GetAllOrdersByTraderWithoutSign(trader, chainName)
}

func (r *RelayerSrv) GetOrderQueue(chain string) ([]*models.Order, error) {
	return r.db.GetOrderQueue(chain)
}

func (r *RelayerSrv) CheckHasBaseToken(token, chain string) (bool, error) {
	return r.db.HasBaseToken(token, chain)
}

func (r *RelayerSrv) GetDepthOrderDetails(token, chain string) ([]*models.Order, error) {
	return r.db.DepthOrderDetails(
		token,
		[]models.MatchedStatus{
			models.MatchedStatusInit,
			models.MatchedStatusPartialMatchConfirmed,
			models.MatchedStatusFailedConfirmed,
		},
		chain,
	)
}
