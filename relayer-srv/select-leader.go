package relayer_srv

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/volmexfinance/relayers/relayer-srv/db"
	"github.com/volmexfinance/relayers/relayer-srv/worker"
)

// This is for selection for leader
// TODO: Add more leader checks here
// TODO: Check for previous leader
// TODO: this should not return error but select other leader
// TODO: select leader from only connected nodes
func (r *RelayerSrv) SelectLeader(wrkr *worker.Worker, currentLeader string) (db.TransactionSent, error) {

	guardianSet := r.node.GuardianSetMapping[wrkr.ChainName]
	var nextLeader string
	var leaderIndex int
	for i := 0; i < len(guardianSet); i++ {
		if strings.EqualFold(wrkr.GetWorkerAddress(), guardianSet[i].Address) {
			if i == len(guardianSet)-1 {
				leaderIndex = 0
			} else {
				leaderIndex = i + 1
			}

		}
	}

	nextLeader = guardianSet[leaderIndex].Address
	for {
		workerBalance, err := wrkr.GetBalance(nextLeader)
		if err != nil {
			return db.TransactionSent{}, fmt.Errorf("SelectLeader: Unable to get balance ")
		}
		if r.node.CheckStatus(guardianSet[leaderIndex].PeerId) && nextLeader != currentLeader && workerBalance != big.NewInt(0) {
			break
		}
		leaderIndex += 1
		if leaderIndex > len(guardianSet)-1 {
			leaderIndex = 0
		}
		nextLeader = guardianSet[leaderIndex].Address
	}

	//update DB
	txUpdate := db.TransactionSent{Leader: r.node.GetPeerID().String(), TransactionStatus: db.TransactionStatusTypeInit, LeaderPublicKey: nextLeader}
	if err := r.db.CreateTxSent(&txUpdate, wrkr.ChainName); err != nil {
		return db.TransactionSent{}, fmt.Errorf("SelectLeader: %w", err)
	}
	return txUpdate, nil
}
