package utils

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	protocols_p2p "github.com/volmexfinance/relayers/relayer-srv/chat"
	"github.com/volmexfinance/relayers/relayer-srv/db"
	Periphery "github.com/volmexfinance/relayers/relayer-srv/worker/abi/periphery"
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

// OrderToLibOrder converts database to lib order
func OrderToLibOrder(dbOrder *db.Order) Periphery.LibOrderOrder {
	makeValue := dbOrder.MakeAsset().ValueAsBigInt()
	takeValue := dbOrder.TakeAsset().ValueAsBigInt()
	salt := dbOrder.OrderSalt()
	triggerPrice := dbOrder.OrderTriggerPrice()

	return Periphery.LibOrderOrder{
		OrderType: StringToBytes4(dbOrder.OrderType),
		Deadline:  dbOrder.Deadline,
		Trader:    common.HexToAddress(dbOrder.Trader),
		MakeAsset: Periphery.LibAssetAsset{
			VirtualToken: common.HexToAddress(dbOrder.MakeAsset().VirtualToken),
			Value:        makeValue,
		},
		TakeAsset: Periphery.LibAssetAsset{
			VirtualToken: common.HexToAddress(dbOrder.TakeAsset().VirtualToken),
			Value:        takeValue,
		},
		Salt:                   salt,
		LimitOrderTriggerPrice: triggerPrice,
		IsShort:                dbOrder.IsShort,
	}
}

// P2POrderToDBOrder converts P2P order to DB order
func P2POrderToDBOrder(order *protocols_p2p.Order, chain string) (*db.Order, error) {
	dbOrder := &db.Order{
		OrderType: order.OrderType,
		Deadline:  order.Deadline,
		Trader:    order.Trader,
		Assets: []db.Assets{
			{
				VirtualToken: order.MakeAsset.VirtualToken,
				Value:        order.MakeAsset.Value,
			},
			{
				VirtualToken: order.TakeAsset.VirtualToken,
				Value:        order.TakeAsset.Value,
			},
		},
		Salt:         order.Salt,
		TriggerPrice: order.LimitOrderTriggerPrice,
		IsShort:      order.IsShort,
		Sign:         order.Sign,
		Fills:        order.Fills,
		ChainName:    chain,
	}

	dbOrder.OrderID = db.CreateOrderID(dbOrder.Trader, dbOrder.Salt, chain)

	makeValue := dbOrder.MakeAsset().ValueAsBigFloat()
	takeValue := dbOrder.TakeAsset().ValueAsBigFloat()
	dbOrder.Price = db.CalculatePrice(order.IsShort, makeValue, takeValue)

	return dbOrder, nil
}

// DBOrderToP2POrder converts DB order to P2P order format
func DBOrderToP2POrder(order *db.Order) (*protocols_p2p.Order, error) {
	return &protocols_p2p.Order{
		OrderType: order.OrderType,
		Deadline:  order.Deadline,
		Trader:    order.Trader,
		MakeAsset: &protocols_p2p.Asset{
			VirtualToken: order.MakeAsset().VirtualToken,
			Value:        order.MakeAsset().Value,
		},
		TakeAsset: &protocols_p2p.Asset{
			VirtualToken: order.TakeAsset().VirtualToken,
			Value:        order.TakeAsset().Value,
		},
		Salt:                   order.Salt,
		LimitOrderTriggerPrice: order.TriggerPrice,
		IsShort:                order.IsShort,
		Sign:                   order.Sign,
		Fills:                  order.Fills,
	}, nil
}

// EncodePeripheryContractData returns data representing openPosition call
func EncodePeripheryContractData(dbOrderLeft, dbOrderRight []*db.Order, contractAddress string) ([]byte, error) {
	libLeftOrderArray := []db.LibOrderOrder{}
	leftOrderSignatureArray := [][]byte{}

	libRightOrderArray := []db.LibOrderOrder{}
	rightOrderSignatureArray := [][]byte{}
	for i := 0; i < len(dbOrderLeft); i++ {
		makevalue1 := dbOrderLeft[i].MakeAsset().ValueAsBigInt()
		takevalue1 := dbOrderLeft[i].TakeAsset().ValueAsBigInt()

		salt1 := dbOrderLeft[i].OrderSalt()
		triggerPrice1 := dbOrderLeft[i].OrderTriggerPrice()

		order1 := &db.LibOrderOrder{
			OrderType:              StringToBytes4(dbOrderLeft[i].OrderType),
			Trader:                 common.HexToAddress(dbOrderLeft[i].Trader),
			Deadline:               dbOrderLeft[i].Deadline,
			IsShort:                dbOrderLeft[i].IsShort,
			MakeAsset:              db.LibAssetAsset{VirtualToken: common.HexToAddress(dbOrderLeft[i].MakeAsset().VirtualToken), Value: makevalue1},
			TakeAsset:              db.LibAssetAsset{VirtualToken: common.HexToAddress(dbOrderLeft[i].TakeAsset().VirtualToken), Value: takevalue1},
			LimitOrderTriggerPrice: triggerPrice1,
			Salt:                   salt1,
		}

		libLeftOrderArray = append(libLeftOrderArray, *order1)
		signature := common.FromHex(dbOrderLeft[i].Sign)
		leftOrderSignatureArray = append(leftOrderSignatureArray, signature)
	}

	for i := 0; i < len(dbOrderRight); i++ {
		makevalue := dbOrderRight[i].MakeAsset().ValueAsBigInt()
		takevalue := dbOrderRight[i].TakeAsset().ValueAsBigInt()
		salt := dbOrderRight[i].OrderSalt()
		triggerPrice := dbOrderRight[i].OrderTriggerPrice()

		order := &db.LibOrderOrder{
			OrderType:              StringToBytes4(dbOrderRight[i].OrderType),
			Trader:                 common.HexToAddress(dbOrderRight[i].Trader),
			Deadline:               dbOrderRight[i].Deadline,
			IsShort:                dbOrderRight[i].IsShort,
			MakeAsset:              db.LibAssetAsset{VirtualToken: common.HexToAddress(dbOrderRight[i].MakeAsset().VirtualToken), Value: makevalue},
			TakeAsset:              db.LibAssetAsset{VirtualToken: common.HexToAddress(dbOrderRight[i].TakeAsset().VirtualToken), Value: takevalue},
			LimitOrderTriggerPrice: triggerPrice,
			Salt:                   salt,
		}

		libRightOrderArray = append(libRightOrderArray, *order)
		signature := common.FromHex(dbOrderRight[i].Sign)
		rightOrderSignatureArray = append(rightOrderSignatureArray, signature)
	}

	posAbi, newerror := abi.JSON(strings.NewReader(Periphery.PeripheryABI))
	if newerror != nil {
		return nil, fmt.Errorf("EncodePeripheryContractData: %w", newerror)
	}

	method := posAbi.Methods["batchOpenPosition"]

	//TODO: Change index on the basis of positioning contract
	packed, newerror := method.Inputs.Pack(big.NewInt(0), libLeftOrderArray, leftOrderSignatureArray, libRightOrderArray, rightOrderSignatureArray, []byte{0})
	if newerror != nil {
		return nil, fmt.Errorf("EncodePeripheryContractData: error in packing %w", newerror)
	}

	bytesJoined := bytes.Join([][]byte{method.ID, packed}, nil)
	return bytesJoined, nil
}

// EncodeOrderStruct encodes order struct in bytes
func EncodeOrderStruct(dbOrder db.Order, chainID int64, positioningContract string) ([]byte, error) {
	makeAsset := make(map[string]interface{})
	makeAsset["virtualToken"] = dbOrder.MakeAsset().VirtualToken
	makeAsset["value"] = dbOrder.MakeAsset().Value

	takeAsset := make(map[string]interface{})
	takeAsset["virtualToken"] = dbOrder.TakeAsset().VirtualToken
	takeAsset["value"] = dbOrder.TakeAsset().Value
	typeddata := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"Asset": []apitypes.Type{
				{Name: "virtualToken", Type: "address"},
				{Name: "value", Type: "uint256"},
			},
			"Order": []apitypes.Type{
				{Name: "orderType", Type: "bytes4"},
				{Name: "deadline", Type: "uint64"},
				{Name: "trader", Type: "address"},
				{Name: "makeAsset", Type: "Asset"},
				{Name: "takeAsset", Type: "Asset"},
				{Name: "salt", Type: "uint256"},
				{Name: "limitOrderTriggerPrice", Type: "uint128"},
				{Name: "isShort", Type: "bool"},
			},
		},
		PrimaryType: "Order",
		Domain: apitypes.TypedDataDomain{
			Name:              "V_PERP",
			Version:           "1",
			ChainId:           math.NewHexOrDecimal256(chainID),
			VerifyingContract: positioningContract,
		},
		Message: apitypes.TypedDataMessage{
			"orderType":              dbOrder.OrderType,
			"deadline":               fmt.Sprintf("%d", dbOrder.Deadline),
			"trader":                 dbOrder.Trader,
			"makeAsset":              makeAsset,
			"takeAsset":              takeAsset,
			"salt":                   dbOrder.Salt,
			"limitOrderTriggerPrice": dbOrder.TriggerPrice,
			"isShort":                dbOrder.IsShort,
		},
	}

	rawData, _, err := apitypes.TypedDataAndHash(typeddata)

	if err != nil {
		return nil, err
	}
	return rawData, nil
}

// DBOrderToLibOrder converts DB order to lib order
func DBOrderToLibOrder(order *db.Order) db.LibOrderOrder {
	makevalue := order.MakeAsset().ValueAsBigInt()
	takevalue := order.TakeAsset().ValueAsBigInt()

	salt := order.OrderSalt()
	triggerPrice := order.OrderTriggerPrice()

	return db.LibOrderOrder{
		OrderType:              StringToBytes4(order.OrderType),
		Trader:                 common.HexToAddress(order.Trader),
		Deadline:               order.Deadline,
		IsShort:                order.IsShort,
		MakeAsset:              db.LibAssetAsset{VirtualToken: common.HexToAddress(order.MakeAsset().VirtualToken), Value: makevalue},
		TakeAsset:              db.LibAssetAsset{VirtualToken: common.HexToAddress(order.TakeAsset().VirtualToken), Value: takevalue},
		LimitOrderTriggerPrice: triggerPrice,
		Salt:                   salt,
	}
}

func SortStringArray(strArr []string) []string {
	sort.Slice(strArr, func(i, j int) bool {
		cd := collate.New(language.English)
		a := strings.ToLower(strArr[i])
		b := strings.ToLower(strArr[j])
		return cd.CompareString(a, b) > 0
	})
	return strArr
}

func StringToBytes4(b string) [4]byte {
	var byteArr [4]byte
	copy(byteArr[:], common.FromHex(b))
	return byteArr
}

func StringToBytes(b string) []byte {
	var byteArr []byte
	copy(byteArr[:], b)
	return byteArr
}

func BytesToBytes4(b []byte) [4]byte {
	var byteArr [4]byte
	copy(byteArr[:], b)
	return byteArr
}

func BytesToBytes32(b []byte) [32]byte {
	var byteArr [32]byte
	copy(byteArr[:], b)
	return byteArr
}

// GetPrivateKey returns ecdsa private key
func GetPrivateKey(privKey string) (*ecdsa.PrivateKey, error) {
	priKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return nil, err
	}
	return priKey, nil
}

// GetOrderIdsFromTLog created order ID from tlog
func GetOrderIdsFromTLog(tLog *db.TransactionLog, chain string) (string, string) {
	return db.CreateOrderID(tLog.Traders[0], tLog.Salt[0], chain), db.CreateOrderID(tLog.Traders[1], tLog.Salt[1], chain)
}

// TODO: error handling
// CreateTransactionMsgP2P retuns gossip message for transaction sent to be used in p2p
func CreateTransactionMsgP2P(txnSent *db.TransactionSent, newTxnSent *db.TransactionSent, chain string) (*protocols_p2p.GossipMessage, error) {
	txnMsg := &protocols_p2p.TransactionMessage{
		TransactionHash: txnSent.TransactionHash,
		TxID:            int64(txnSent.ID),
		OrderID:         txnSent.OrderID,
		NextLeader:      newTxnSent.Leader,
		NextPubKey:      newTxnSent.LeaderPublicKey,
	}
	hash, err := json.Marshal(&txnMsg)
	if err != nil {
		return nil, err
	}
	return &protocols_p2p.GossipMessage{
		Message: &protocols_p2p.GossipMessage_TransactionMessage{
			TransactionMessage: &protocols_p2p.TransactionMessageMessageData{
				MessageData: &protocols_p2p.MessageData{
					Data:  crypto.Keccak256(hash),
					Chain: chain,
				},
				TransactionMessage: txnMsg,
			},
		}}, nil
}

func CreateTransactionFailMsgP2P(orderID []string, errorMessage string, chainName string) *protocols_p2p.GossipMessage {
	TxnMsg := &protocols_p2p.TransactionMessage{
		TransactionHash: "",
		TxID:            0,
		OrderID:         orderID,
		NextLeader:      "",
		NextPubKey:      "",
		Error:           errorMessage,
	}
	hash, err := json.Marshal(&TxnMsg)
	if err != nil {
		println("err in creating transaction hash", err)
	}
	return &protocols_p2p.GossipMessage{
		Message: &protocols_p2p.GossipMessage_TransactionMessage{
			TransactionMessage: &protocols_p2p.TransactionMessageMessageData{
				MessageData: &protocols_p2p.MessageData{
					Data:  crypto.Keccak256(hash),
					Chain: chainName,
				},
				TransactionMessage: TxnMsg,
			},
		}}
}
func CreateLeaderChangeMsgP2P(orderID []string, leaderNodeID, leaderPubKey, chainName string) *protocols_p2p.GossipMessage {
	TxnMsg := &protocols_p2p.TransactionMessage{
		TransactionHash: "",
		TxID:            0,
		OrderID:         orderID,
		NextLeader:      leaderNodeID,
		NextPubKey:      leaderPubKey,
		Error:           "",
	}
	hash, err := json.Marshal(&TxnMsg)
	if err != nil {
		println("err in creating transaction hash", err)
	}
	return &protocols_p2p.GossipMessage{
		Message: &protocols_p2p.GossipMessage_TransactionMessage{
			TransactionMessage: &protocols_p2p.TransactionMessageMessageData{
				MessageData: &protocols_p2p.MessageData{
					Data:  crypto.Keccak256(hash),
					Chain: chainName,
				},
				TransactionMessage: TxnMsg,
			},
		}}
}
