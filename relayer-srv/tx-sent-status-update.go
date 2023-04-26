package relayer_srv

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/volmexfinance/relayers/relayer-srv/db"
	"github.com/volmexfinance/relayers/relayer-srv/worker"
)

// TODO: sanitization make this as go routine
// TODO: handle order status as well here
func (r *RelayerSrv) UpdateTxnSentStatus(wrkr *worker.Worker) error {
	transactionStatusType := []db.TransactionStatusType{db.TransactionStatusTypeInit, db.TransactionStatusTypePending, db.TransactionStatusTypeNotFound}
	transactionToCheck, err := r.postgresDB.GetTxnsOnStatus(transactionStatusType, wrkr.ChainName)
	if err != nil {
		return fmt.Errorf("failed to get transaction from db")
	}

	for _, txn := range *transactionToCheck {
		txReceipt, er := wrkr.GetTransactionReceipt(txn.TransactionHash) // create context here

		if er != nil {
			if errors.Is(er, ethereum.NotFound) { // Case : Not Found
				err := r.postgresDB.UpdateTxnStatus(&txn, db.TransactionStatusTypeNotFound)
				if err != nil {
					return fmt.Errorf("failed to update%w", err)
				}
			} else if strings.Contains(er.Error(), "transaction pending") && txn.TransactionStatus != db.TransactionStatusTypePending { // Case : Pending
				err := r.postgresDB.UpdateTxnStatus(&txn, db.TransactionStatusTypePending)
				if err != nil {
					return fmt.Errorf("failed to update %w", err)
				}
			} else {
				return fmt.Errorf("failed to get transaction from client %w", er) // Case : error
			}
		}

		if txReceipt != nil {
			if txReceipt.Status == 1 { // Case : Success
				err := r.postgresDB.UpdateTxnStatus(&txn, db.TransactionStatusTypeSuccess)
				if err != nil {
					return fmt.Errorf("failed to update%w", err)
				}
			} else if txReceipt.Status == 0 { // Case : Failed
				err := r.postgresDB.UpdateTxnStatus(&txn, db.TransactionStatusTypeFailed)
				if err != nil {
					return fmt.Errorf("failed to update%w", err)
				}
			}
		}
	}
	return nil
}
