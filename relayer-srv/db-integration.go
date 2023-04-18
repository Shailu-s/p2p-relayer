package relayer_srv

import (
	"time"

	"github.com/volmexfinance/relayers/relayer-srv/db"
	"github.com/volmexfinance/relayers/relayer-srv/worker"
)

// TODO: Update status for which FindFulfilledFailedOrder are implemented
// TODO: Create a go routine for it
func (r *RelayerSrv) MoveSQLiteToPostgres(wrkr *worker.Worker) {
	for {
		currentTime, err := wrkr.GetCurrentTimestamp()
		if err != nil {
			r.logger.Errorf("DBIntegration: Unable to get current timestamp %v", err)
			continue
		}
		if ok := r.db.UpdateDeadlinePassedOrder(currentTime, []db.MatchedStatus{db.MatchedStatusSentToContract, db.MatchedStatusBlocked}, wrkr.ChainName); ok != nil {
			r.logger.Errorf("DBIntegration: Unable to update deadline passed order %v", err)
			continue
		}

		fulfilledOrders, err := r.db.FindFulfilledFailedOrder([]db.MatchedStatus{db.Canceled, db.MatchedStatusFullMatchConfirmed, db.MatchedStatusBlocked}, wrkr.ChainName)
		if err != nil {
			r.logger.Errorf("DBIntegration: %v", err)
			continue
		}

		if len(fulfilledOrders) == 0 {
			r.logger.Errorf("DBIntegration:Did not received any any order from SQLite %v", err)
			continue
		}

		if err := r.postgresDB.CreateOrderBatch(fulfilledOrders); err != nil {
			r.logger.Errorf("DBIntegration:Unable to create %v", err)
			continue
		}

		if err := r.db.DeleteBatchOrder(fulfilledOrders); err != nil {
			r.logger.Errorf("DBIntegration: Unable to delete %v", err)
			continue
		}
		r.logger.Infof("DBIntegration: Orders migration completed from SQLIte to Postgres")
		time.Sleep(10 * time.Second)
	}
}
