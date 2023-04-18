package worker

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	gnosis_abi "github.com/volmexfinance/relayers/relayer-srv/worker/abi/gnosis"
	matching_engine_abi "github.com/volmexfinance/relayers/relayer-srv/worker/abi/matching_engine"
)

// ParseEvent parses data of events found on blockchain in required structure
func ParseEvent(log *types.Log) (interface{}, string, error) {
	orderFillSig := "OrdersFilled(address[2],uint256[2],uint256[2])"

	orderFillSigHash := crypto.Keccak256Hash([]byte(orderFillSig))
	matchingEngineABI, err := abi.JSON(strings.NewReader(matching_engine_abi.MatchingEngineMetaData.ABI))
	if err != nil {
		return MatchingEngineAbiOrdersFilled{}, "", fmt.Errorf("parseEvent:%w", err)
	}

	cancelledSig := "Canceled(bytes32,address,address,uint256,uint256)"
	cancelledSigHash := crypto.Keccak256Hash([]byte(cancelledSig))

	cancelledAllSig := "CanceledAll(address,uint256)"
	cancelledAllSigHash := crypto.Keccak256Hash([]byte(cancelledAllSig))

	var ordersFilled MatchingEngineAbiOrdersFilled
	var canceled MatchingEngineAbiCanceled
	var canceledAll MatchingEngineAbiCanceledAll

	addedOwnerSig := "AddedOwner(address)"
	addedOwnerSigHash := crypto.Keccak256Hash([]byte(addedOwnerSig))

	removedOwnerSig := "RemovedOwner(address)"
	removedOwnerSigHash := crypto.Keccak256Hash([]byte(removedOwnerSig))

	changedThresholdSig := "ChangedThreshold(uint256)"
	changedThresholdSigHash := crypto.Keccak256Hash([]byte(changedThresholdSig))

	gnosisABI, err := abi.JSON(strings.NewReader(string(gnosis_abi.GnosisMetaData.ABI)))
	if err != nil {
		return MatchingEngineAbiOrdersFilled{}, "", fmt.Errorf("parseEvent:%w", err)
	}

	var addedOwner GnosisAddedOwner
	var removedOwner GnosisRemovedOwner
	var changedThreshold GnosisChangedThreshold

	switch log.Topics[0].Hex() {
	case orderFillSigHash.Hex():
		err = matchingEngineABI.UnpackIntoInterface(&ordersFilled, "OrdersFilled", log.Data)
		if err != nil {
			return MatchingEngineAbiOrdersFilled{}, "", fmt.Errorf("parseEvent:%w", err)
		}
		return ordersFilled, "OrdersFilled", nil
	case cancelledSigHash.Hex():
		err = matchingEngineABI.UnpackIntoInterface(&canceled, "Canceled", log.Data)
		if err != nil {
			return MatchingEngineAbiOrdersFilled{}, "", fmt.Errorf("parseEvent:%w", err)
		}
		return canceled, "Canceled", nil
	case cancelledAllSigHash.Hex():
		err = matchingEngineABI.UnpackIntoInterface(&canceledAll, "CanceledAll", log.Data)
		if err != nil {
			return MatchingEngineAbiOrdersFilled{}, "", fmt.Errorf("parseEvent:%w", err)
		}
		canceledAll.Trader = common.HexToAddress(log.Topics[1].Hex())
		return canceledAll, "CanceledAll", nil
	case addedOwnerSigHash.Hex():
		err = gnosisABI.UnpackIntoInterface(&addedOwner, "AddedOwner", log.Data)
		if err != nil {
			return GnosisAddedOwner{}, "", fmt.Errorf("parseEvent:%w", err)
		}
		return addedOwner, "AddedOwner", nil
	case removedOwnerSigHash.Hex():
		err = gnosisABI.UnpackIntoInterface(&removedOwner, "RemovedOwner", log.Data)
		if err != nil {
			return GnosisRemovedOwner{}, "", fmt.Errorf("parseEvent:%w", err)
		}
		return removedOwner, "RemovedOwner", nil
	case changedThresholdSigHash.Hex():
		err = gnosisABI.UnpackIntoInterface(&changedThreshold, "ChangedThreshold", log.Data)
		if err != nil {
			return GnosisChangedThreshold{}, "", fmt.Errorf("parseEvent:%w", err)
		}
		return changedThreshold, "ChangedThreshold", nil
	}

	return MatchingEngineAbiOrdersFilled{}, "", nil

}
