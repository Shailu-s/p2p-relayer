package relayer_srv

import (
	"fmt"
	"sync"

	"github.com/volmexfinance/relayers/relayer-srv/db/models"
	"github.com/volmexfinance/relayers/relayer-srv/worker"
)

var lock sync.Mutex

// TODO: relayer sign check should be included here
func (r *RelayerSrv) SendToContract(orderLeft, orderRight []*models.Order, orderIDs []string, relayerSigns [][]byte, wrkr *worker.Worker) (*models.TransactionSent, error) {
	lock.Lock()
	defer lock.Unlock()
	txSent := models.TransactionSent{
		TransactionStatus: models.TransactionStatusTypeInit,
		OrderID:           orderIDs,
	}
	orderCombined := append(orderLeft, orderRight...)

	// TODO: double spending
	txHash, er := wrkr.ExecuteGnosisTx(orderLeft, orderRight, relayerSigns)
	if er != nil {
		txSent.TransactionStatus = models.TransactionStatusTypeSentFailed
		txSent.Error = er.Error()
		r.logger.Warnf("Error sending to contract: %s", er.Error())
		if err := r.db.UpdateFillAndStatusByTxnLog(orderCombined, models.MatchedStatusSentFailed); err != nil {
			return &txSent, fmt.Errorf("SendToContract: Unable to update %w", er)
		}
		return &txSent, fmt.Errorf("SendToContract: %w", er)
	}
	r.logger.Infof("tx sent success for orderIDs: %s, txhash: %s", orderIDs, txHash)
	// updating tx_sents table
	txSent.TransactionHash = txHash
	if err := r.db.UpdateBatchOrderStatus(orderCombined, models.MatchedStatusSentToContract); err != nil {
		return &txSent, fmt.Errorf("SendToContract: %w", err)
	}

	if err := r.db.CreateTxSent(&txSent, wrkr.ChainName); err != nil {
		return nil, err
	}
	return &txSent, nil
}
