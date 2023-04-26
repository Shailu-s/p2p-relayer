package relayer_srv

import (
	"fmt"
	"sync"

	"github.com/volmexfinance/relayers/relayer-srv/db"
	"github.com/volmexfinance/relayers/relayer-srv/worker"
)

var lock sync.Mutex

// TODO: relayer sign check should be included here
func (r *RelayerSrv) SendToContract(orderLeft, orderRight []*db.Order, orderIDs []string, relayerSigns [][]byte, wrkr *worker.Worker) (*db.TransactionSent, error) {
	lock.Lock()
	defer lock.Unlock()
	txSent := db.TransactionSent{
		TransactionStatus: db.TransactionStatusTypeInit,
		OrderID:           orderIDs,
	}
	orderCombined := append(orderLeft, orderRight...)

	// TODO: double spending
	txHash, er := wrkr.ExecuteGnosisTx(orderLeft, orderRight, relayerSigns)
	if er != nil {
		txSent.TransactionStatus = db.TransactionStatusTypeSentFailed
		txSent.Error = er.Error()
		r.logger.Warnf("Error sending to contract: %s", er.Error())
		if err := r.postgresDB.UpdateFillAndStatusByTxnLog(orderCombined, db.MatchedStatusSentFailed); err != nil {
			return &txSent, fmt.Errorf("SendToContract: Unable to update %w", er)
		}
		if err := r.sqlitedb.UpdateFillAndStatusByTxnLog(orderCombined, db.MatchedStatusSentFailed); err != nil {
			return &txSent, fmt.Errorf("SendToContract: Unable to update %w", er)
		}
		if err := r.postgresDB.CreateTxSent(&txSent, wrkr.ChainName); err != nil {
			return nil, err
		}
		if err := r.sqlitedb.CreateTxSent(&txSent, wrkr.ChainName); err != nil {
			return nil, err
		}
		return &txSent, fmt.Errorf("SendToContract: %w", er)
	}
	r.logger.Infof("tx sent success for orderIDs: %s, txhash: %s", orderIDs, txHash)
	// updating tx_sents table
	txSent.TransactionHash = txHash
	if err := r.postgresDB.UpdateBatchOrderStatus(orderCombined, db.MatchedStatusSentToContract); err != nil {
		return &txSent, fmt.Errorf("SendToContract: %w", err)
	}

	if err := r.postgresDB.CreateTxSent(&txSent, wrkr.ChainName); err != nil {
		return nil, err
	}

	if err := r.sqlitedb.CreateTxSent(&txSent, wrkr.ChainName); err != nil {
		return nil, err
	}
	return &txSent, nil
}
