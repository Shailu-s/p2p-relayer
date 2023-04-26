package relayer_srv

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/volmexfinance/relayers/relayer-srv/db/models"
	"github.com/volmexfinance/relayers/relayer-srv/worker"
)

// TODO: sanitization make this as go routine
// TODO: handle order status as well here
func (r *RelayerSrv) UpdateTxnSentStatus(wrkr *worker.Worker) error {
	transactionStatusType := []models.TransactionStatusType{models.TransactionStatusTypeInit, models.TransactionStatusTypePending, models.TransactionStatusTypeNotFound}
	transactionToCheck, err := r.db.GetTxnsOnStatus(transactionStatusType, wrkr.ChainName)
	if err != nil {
		return fmt.Errorf("failed to get transaction from models")
	}

	for _, txn := range *transactionToCheck {
		txReceipt, er := wrkr.GetTransactionReceipt(txn.TransactionHash) // create context here

		if er != nil {
			if errors.Is(er, ethereum.NotFound) { // Case : Not Found
				err := r.db.UpdateTxnStatus(&txn, models.TransactionStatusTypeNotFound)
				if err != nil {
					return fmt.Errorf("failed to update%w", err)
				}
			} else if strings.Contains(er.Error(), "transaction pending") && txn.TransactionStatus != models.TransactionStatusTypePending { // Case : Pending
				err := r.db.UpdateTxnStatus(&txn, models.TransactionStatusTypePending)
				if err != nil {
					return fmt.Errorf("failed to update %w", err)
				}
			} else {
				return fmt.Errorf("failed to get transaction from client %w", er) // Case : error
			}
		}

		if txReceipt != nil {
			if txReceipt.Status == 1 { // Case : Success
				err := r.db.UpdateTxnStatus(&txn, models.TransactionStatusTypeSuccess)
				if err != nil {
					return fmt.Errorf("failed to update%w", err)
				}
			} else if txReceipt.Status == 0 { // Case : Failed
				err := r.db.UpdateTxnStatus(&txn, models.TransactionStatusTypeFailed)
				if err != nil {
					return fmt.Errorf("failed to update%w", err)
				}
			}
		}
	}
	return nil
}
