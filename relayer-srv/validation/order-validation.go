package validation

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/volmexfinance/relayers/relayer-srv/db"
	"github.com/volmexfinance/relayers/relayer-srv/utils"
)

// ValidateMatchedOrder validates two matched order receieved on p2p
func ValidateMatchedOrder(order1, order2 db.Order, dbs *db.DataBase, maxFail int64) error {

	// err := matching_engine.VerifyMatchedOrder(&order1, &order2, dbs, maxFail)
	// if err != nil {
	// 	return fmt.Errorf("ValidateMatchedOrder:VerifyMatchedOrder %w", err)
	// }

	//TODO: check whether this is absolete
	if order1.Trader == order2.Trader || order1.IsShort == order2.IsShort || order1.MakeAsset().VirtualToken != order2.TakeAsset().VirtualToken || order2.TakeAsset().VirtualToken != order1.MakeAsset().VirtualToken {
		return fmt.Errorf("ValidateMatchedOrder: validation failed")
	} else {
		if err := dbs.UpdateOrderStatus(order1.OrderID, []db.MatchedStatus{db.MatchedStatusInit, db.MatchedStatusFullMatchFound, db.MatchedStatusPartialMatchFound, db.MatchedStatusPartialMatchConfirmed, db.MatchedStatusFailedConfirmed},
			[]db.MatchedStatus{db.MatchedStatusValidated, db.MatchedStatusValidationConfirmed, db.MatchedStatusSentToContract, db.MatchedStatusSentFailed, db.MatchedStatusFullMatchConfirmed},
			db.MatchedStatusValidated); err != nil {
			return fmt.Errorf("ValidateMatchedOrder: err in handling status %w", err)
		}
		if err := dbs.UpdateOrderStatus(order2.OrderID, []db.MatchedStatus{db.MatchedStatusInit, db.MatchedStatusFullMatchFound, db.MatchedStatusPartialMatchFound, db.MatchedStatusPartialMatchConfirmed, db.MatchedStatusFailedConfirmed},
			[]db.MatchedStatus{db.MatchedStatusValidated, db.MatchedStatusValidationConfirmed, db.MatchedStatusSentToContract, db.MatchedStatusSentFailed, db.MatchedStatusFullMatchConfirmed},
			db.MatchedStatusValidated); err != nil {
			return fmt.Errorf("ValidateMatchedOrder: err in handling statuses %w", err)
		}

		return nil
	}
}

// ValtoidateThreshold updates status of order in DB to MatchedStatusValidationConfirmed
func ValidateThreshold(orders []*db.Order, dbs *db.DataBase) error {
	for _, order := range orders {
		if err := dbs.UpdateOrderStatus(order.OrderID, []db.MatchedStatus{db.MatchedStatusValidated, db.MatchedStatusFullMatchFound, db.MatchedStatusPartialMatchFound},
			[]db.MatchedStatus{db.MatchedStatusInit, db.MatchedStatusValidationConfirmed, db.MatchedStatusSentToContract, db.MatchedStatusSentFailed, db.MatchedStatusFullMatchConfirmed, db.MatchedStatusPartialMatchConfirmed, db.MatchedStatusFailedConfirmed},
			db.MatchedStatusValidationConfirmed); err != nil {
			return fmt.Errorf("ValidateThreshold: %w", err)
		}
	}
	return nil
}

// TODO: to be picked with security tasks (add order validation according to perp)
// VerifyOrderSignature verfies signature of user provided in order
func VerifyOrderSignature(order db.Order, chainID int64, peripheryContract string) error {
	hash, err := utils.EncodeOrderStruct(order, chainID, peripheryContract)
	if err != nil {
		return err
	}

	sign, err := hexutil.Decode(order.Sign)
	if err != nil {
		return err
	}

	if sign[64] >= 27 {
		sign[64] -= 27 // Transform V from 0/1 to 27/28 according to the yellow paper
	}

	sigPubKeyBytes, err := crypto.Ecrecover(hash, sign)
	if err != nil {
		return err
	}

	var buf []byte

	hash0 := crypto.Keccak256Hash(sigPubKeyBytes[1:])
	buf = hash0.Bytes()
	publicAddress := hexutil.Encode(buf[12:])

	hex := strings.ToLower(publicAddress)[2:]
	checkSummedAddress := "0x"
	hashCs := crypto.Keccak256Hash([]byte(hex))

	for i, b := range hex {
		c := string(b)
		if b < '0' || b > '9' {
			if hashCs[i/2]&byte(128-i%2*120) != 0 {
				c = string(b - 32)
			}
		}
		checkSummedAddress += c
	}

	result := strings.EqualFold(checkSummedAddress, order.Trader)
	if !result {
		return fmt.Errorf("signer != trader")
	}
	return nil
}
