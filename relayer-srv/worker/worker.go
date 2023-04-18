package worker

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"

	database "github.com/volmexfinance/relayers/relayer-srv/db"
	"github.com/volmexfinance/relayers/relayer-srv/utils"
	"github.com/volmexfinance/relayers/relayer-srv/worker/abi/gnosis"
	IndexPriceOracle "github.com/volmexfinance/relayers/relayer-srv/worker/abi/index_price_oracle"
	MarkPriceOracle "github.com/volmexfinance/relayers/relayer-srv/worker/abi/mark_price_oracle"
	Positioning "github.com/volmexfinance/relayers/relayer-srv/worker/abi/positioning"
)

// WorkerConfig ...
type WorkerConfig struct {
	ChainName           string         `json:"chain_name"`
	Provider            string         `json:"provider"`
	MatchingEngine      common.Address `json:"matching_engine_contract"`
	GnosisContract      common.Address `json:"gnosis_contract"`
	PositioningContract common.Address `json:"positioning_contract"`
	PeripheryContract   common.Address `json:"periphery_contract"`
	IndexPriceOracle    common.Address `json:"index_price_oracle"`
	MarkPriceOracle     common.Address `json:"mark_price_oracle"`
	GasLimit            int64          `json:"gas_limit"`
	WorkerAddr          common.Address `json:"worker_addr"`
	StartBlockHeight    *big.Int       `json:"from_block"`
}

// Worker creates an instance and store its information
type Worker struct {
	provider            string
	ChainName           string
	chainID             int64
	Logger              *logrus.Entry // Logger
	config              WorkerConfig
	client              *ethclient.Client
	gnosisContract      common.Address
	positioningContract common.Address
	indexPriceOracle    common.Address
	markPriceOracle     common.Address
	peripheryContract   common.Address
	DB                  *database.DataBase
	PrivateKey          string
	Threshold           int64
}

// MatchingEngineAbiOrdersFilled represents a OrdersFilled event raised by the MatchingEngineAbi contract.
type MatchingEngineAbiOrdersFilled struct {
	Traders [2]common.Address
	Salts   [2]*big.Int
	Fills   [2]*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

type MatchingEngineMatched struct {
	Traders      [2]common.Address
	Deadline     [2]uint64
	Salt         [2]*big.Int
	NewLeftFill  *big.Int
	NewRightFill *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// MatchingEngineAbiCanceled represents a Canceled event raised by the MatchingEngineAbi contract.
type MatchingEngineAbiCanceled struct {
	Hash      [32]byte
	Trader    common.Address
	BaseToken common.Address
	Amount    *big.Int
	Salt      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// MatchingEngineAbiCanceledAll represents a CanceledAll event raised by the MatchingEngineAbi contract.
type MatchingEngineAbiCanceledAll struct {
	Trader  common.Address
	MinSalt *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

type GnosisAddedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// GnosisRemovedOwner represents a RemovedOwner event raised by the Gnosis contract.
type GnosisRemovedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// GnosisChangedThreshold represents a ChangedThreshold event raised by the Gnosis contract.
type GnosisChangedThreshold struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// SafeTransaction used in creating data load for executing gnosis transaction
type SafeTransaction struct {
	To             common.Address
	Value          *big.Int
	Data           []byte
	Operation      uint8
	SafeTxGas      *big.Int
	BaseGas        *big.Int
	GasPrice       *big.Int
	GasToken       common.Address
	RefundReceiver common.Address
	Signature      []byte
}

// NewWorker initialises worker (used for tx on any chain)
func NewWorker(logger *logrus.Logger, cfg WorkerConfig, privateKey string, db *database.DataBase) *Worker {
	client, err := ethclient.Dial(cfg.Provider)
	if err != nil {
		logger.Panicf("rpc error for %s: %v", cfg.ChainName, err)
	}

	privKey, err := utils.GetPrivateKey(privateKey)
	if err != nil {
		logger.Panicf("generate private key error: %v", err)
	}

	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.Panicf("get public key error: %v", err)
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	if !bytes.Equal(cfg.WorkerAddr.Bytes(), fromAddress.Bytes()) {
		logger.Panicf(
			"relayer address supplied in config (%s) does not match mnemonic (%s)",
			cfg.WorkerAddr, fromAddress,
		)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		logger.Panicf("rpc not returning chain id")
	}

	return &Worker{
		ChainName:           cfg.ChainName,
		chainID:             chainId.Int64(),
		Logger:              logger.WithField("worker", cfg.ChainName),
		provider:            cfg.Provider,
		config:              cfg,
		client:              client,
		gnosisContract:      cfg.GnosisContract,
		peripheryContract:   cfg.PeripheryContract,
		positioningContract: cfg.PositioningContract,
		indexPriceOracle:    cfg.IndexPriceOracle,
		markPriceOracle:     cfg.MarkPriceOracle,
		DB:                  db,
		PrivateKey:          privateKey,
	}
}

// GetWorkerAddress loads worker address from config
func (w *Worker) GetWorkerAddress() string {
	return w.config.WorkerAddr.String()
}

// GetLatestBlock returns latest block
func (w *Worker) GetLatestBlock() (*big.Int, error) {
	latestBlock, err := w.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return big.NewInt(0), err
	}
	return latestBlock.Number, nil
}

// GetBalance returns balance of given address
func (w *Worker) GetBalance(address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := w.client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, fmt.Errorf("GetBalance: %w", err)
	}
	return balance, nil
}

func (w *Worker) GetCurrentTimestamp() (uint64, error) {
	latestBlock, err := w.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}
	return latestBlock.Time, nil
}

func (w *Worker) GetGnosisOwners() ([]common.Address, error) {
	GnosisContract, err := gnosis.NewGnosis(w.gnosisContract, w.client)
	if err != nil {
		return []common.Address{}, fmt.Errorf("GetGnosisOwners: %w", err)
	}

	orders, err := GnosisContract.GetOwners(getCallOpts())
	if err != nil {
		return []common.Address{}, fmt.Errorf("GetGnosisOwners: %w", err)
	}
	return orders, nil
}

func (w *Worker) GetThreshold() (*big.Int, error) {
	GnosisContract, err := gnosis.NewGnosis(w.gnosisContract, w.client)
	if err != nil {
		return nil, fmt.Errorf("GetThreshold: %w", err)
	}
	threshold, err := GnosisContract.GetThreshold(getCallOpts())
	if err != nil {
		return nil, fmt.Errorf("GetThreshold: %w", err)
	}
	return threshold, nil
}

// Signs a hash with ecdsa priv key
func SignHash(hash []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	sig, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return nil, err
	}

	if sig[64] < 2 {
		sig[64] += 27 // Transform V from 0/1 to 27/28 according to the yellow paper
	}

	return sig, nil
}

func (w *Worker) GetNonce() (*big.Int, error) {
	GnosisContract, err := gnosis.NewGnosis(w.gnosisContract, w.client)
	if err != nil {
		return nil, fmt.Errorf("getNonce: %w", err)
	}

	nonce, err := GnosisContract.Nonce(getCallOpts())
	if err != nil {
		return nil, fmt.Errorf("getNonce: %w", err)
	}
	return nonce, nil
}

// CreateGnosisTxAndHash creates encoded gnosis tx with relayer signature
func (w *Worker) CreateGnosisTxAndHash(buyOrder, sellOrder []*database.Order) (SafeTransaction, []byte, error) {
	GnosisContract, err := gnosis.NewGnosis(w.gnosisContract, w.client)
	if err != nil {
		return SafeTransaction{}, nil, fmt.Errorf("CreateGnosisTx: %w", err)
	}

	nonce, err := GnosisContract.Nonce(getCallOpts())
	if err != nil {
		return SafeTransaction{}, nil, fmt.Errorf("CreateGnosisTx: %w", err)
	}

	data, err := utils.EncodePeripheryContractData(buyOrder, sellOrder, w.peripheryContract.Hex())
	if err != nil {
		return SafeTransaction{}, nil, fmt.Errorf("CreateGnosisTx: %w", err)
	}

	safeTx := SafeTransaction{
		To:             w.peripheryContract,
		Value:          big.NewInt(0),    // value
		Data:           []byte(data),     // data, the contract call user wants
		Operation:      0,                // operation
		SafeTxGas:      new(big.Int),     // safeTxGas
		BaseGas:        new(big.Int),     // baseGas
		GasPrice:       new(big.Int),     // gasPrice
		GasToken:       common.Address{}, // gasToken
		RefundReceiver: common.Address{}, // refundReceiver
	}

	rawSafeTx, err := GnosisContract.EncodeTransactionData(
		getCallOpts(),
		safeTx.To,
		safeTx.Value,
		safeTx.Data,
		safeTx.Operation,
		safeTx.SafeTxGas,
		safeTx.BaseGas,
		safeTx.GasPrice,
		safeTx.GasToken,
		safeTx.RefundReceiver,
		nonce,
	)
	if err != nil {
		return SafeTransaction{}, nil, fmt.Errorf("CreateGnosisTx: %w", err)
	}
	safeTxHashBytes := crypto.Keccak256(rawSafeTx)
	return safeTx, safeTxHashBytes, nil
}

func (w *Worker) OrderValidation(order database.Order) (bool, error) {
	PositioningContract, err := Positioning.NewPositioning(w.positioningContract, w.client)
	if err != nil {
		return false, fmt.Errorf("OrderValidation: %w", err)
	}

	makevalue := order.MakeAsset().ValueAsBigInt()
	takevalue := order.TakeAsset().ValueAsBigInt()

	salt := order.OrderSalt()
	triggerPrice := order.OrderTriggerPrice()

	posOrder := Positioning.LibOrderOrder{
		OrderType:              utils.StringToBytes4(order.OrderType),
		Trader:                 common.HexToAddress(order.Trader),
		Deadline:               order.Deadline,
		IsShort:                order.IsShort,
		MakeAsset:              Positioning.LibAssetAsset{VirtualToken: common.HexToAddress(order.MakeAsset().VirtualToken), Value: makevalue},
		TakeAsset:              Positioning.LibAssetAsset{VirtualToken: common.HexToAddress(order.TakeAsset().VirtualToken), Value: takevalue},
		LimitOrderTriggerPrice: triggerPrice,
		Salt:                   salt,
	}

	return PositioningContract.GetOrderValidate(getCallOpts(), posOrder)
}

func EncodeGnosisContractData(safeTx SafeTransaction, signature []byte) ([]byte, error) {
	gnosisAbi, newerror := abi.JSON(strings.NewReader(gnosis.GnosisABI))
	if newerror != nil {
		return nil, fmt.Errorf("EncodeGnosisContractData: %w", newerror)
	}

	method := gnosisAbi.Methods["execTransaction"]

	packed, newerror := method.Inputs.Pack(safeTx.To, safeTx.Value, safeTx.Data, safeTx.Operation, safeTx.SafeTxGas, safeTx.BaseGas, safeTx.GasPrice, safeTx.GasToken, safeTx.RefundReceiver, signature)
	if newerror != nil {
		return nil, fmt.Errorf("EncodeGnosisContractData: %w", newerror)
	}

	bytesJoined := bytes.Join([][]byte{method.ID, packed}, nil)
	return bytesJoined, nil
}

func (w *Worker) RedialClient() error {
	w.client.Close()

	ethC, err := ethclient.Dial(w.config.Provider)
	if err != nil {
		return fmt.Errorf("ExecuteGnosisTx: Unable to dial RPC %w", err)
	}

	w.client = ethC
	return nil
}

// Executes openPosition through gnosis
func (w *Worker) ExecuteGnosisTx(orderLeft, orderRight []*database.Order, signatures [][]byte) (string, error) {

	newSafeTx, _, err := w.CreateGnosisTxAndHash(orderLeft, orderRight)

	if err != nil {
		return "", fmt.Errorf("ExecuteGnosisTx: error in creating gnosis txn %w", err)
	}
	// w.client.Close()

	// ethC, err := ethclient.Dial(w.config.Provider)
	// if err != nil {
	// 	return "", fmt.Errorf("ExecuteGnosisTx: Unable to dial RPC %w", err)
	// }

	// w.client = ethC
	estimatedGas, err := w.client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  w.config.GnosisContract,
		To:    &newSafeTx.To,
		Value: newSafeTx.Value,
		Data:  newSafeTx.Data,
	})
	if err != nil {
		return "", fmt.Errorf("ExecuteGnosisTx: Unable to estimate gas: %w", err)
	}
	auth, err := w.getTransactor(uint64(estimatedGas), w.client)
	if err != nil {
		return "", fmt.Errorf("ExecuteGnosisTx: error in get auth %w", err)
	}

	GnosisContract, err := gnosis.NewGnosis(w.gnosisContract, w.client)
	if err != nil {
		return "", fmt.Errorf("ExecuteGnosisTx: error in getting gnosis %w", err)
	}

	var sigToSend string
	for i := len(signatures) - 1; i >= 0; i-- {
		if signatures[i][64] < 2 {
			signatures[i][64] += 27 // Transform V from 0/1 to 27/28 according to the yellow paper
		}
		// newSafeTx.Signature = append(newSafeTx.Signature, signatures[i]...)
		sigToSend += common.Bytes2Hex(signatures[i])
	}

	newSafeTx.Signature = bytes.Join(signatures, nil)

	w.Logger.Info("Estimated gas for txn: ", estimatedGas)
	// Relay the signed transaction
	relayTx, err := GnosisContract.ExecTransaction(
		auth,
		newSafeTx.To,              // to
		newSafeTx.Value,           // value
		newSafeTx.Data,            // data
		newSafeTx.Operation,       // operation
		newSafeTx.SafeTxGas,       // safeTxGas
		newSafeTx.BaseGas,         // baseGas
		newSafeTx.GasPrice,        // gasPrice
		newSafeTx.GasToken,        // gasToken
		newSafeTx.RefundReceiver,  // refundReceiver
		common.FromHex(sigToSend), // signature
	)
	if err != nil {
		return "", fmt.Errorf("ExecuteGnosisTx: %w", err)
	}
	return relayTx.Hash().String(), nil
}

// // calls openPosition function on periphery contract
// func (w *Worker) OpenPosition(order1, order2 database.Order, key1, key2 string) (string, error) {
// 	privkey1, err := utils.GetPrivateKey(key1)
// 	if err != nil {
// 		return "", fmt.Errorf("OpenPosition: %w", err)
// 	}
// 	privkey2, err := utils.GetPrivateKey(key2)
// 	if err != nil {
// 		return "", fmt.Errorf("OpenPosition: %w", err)
// 	}

// 	hash1, err := utils.EncodeOrderStruct(order1, w.chainID, w.peripheryContract.String())
// 	if err != nil {
// 		return "", fmt.Errorf("OpenPosition: %w", err)
// 	}

// 	hash2, err := utils.EncodeOrderStruct(order2, w.chainID, w.peripheryContract.String())
// 	if err != nil {
// 		return "", fmt.Errorf("OpenPosition: %w", err)
// 	}

// 	signature1, err := SignHash(hash1, privkey1)
// 	if err != nil {
// 		return "", fmt.Errorf("OpenPosition: %w", err)
// 	}

// 	signature2, er := SignHash(hash2, privkey2)
// 	if er != nil {
// 		return "", fmt.Errorf("OpenPosition: %w", err)
// 	}

// 	order1.Sign = hexutil.Encode(signature1)
// 	order2.Sign = hexutil.Encode(signature2)

// 	PeripheryAddress := w.peripheryContract
// 	PeripheryContract, err := Periphery.NewPeriphery(PeripheryAddress, w.client)
// 	if err != nil {
// 		return "", fmt.Errorf("OpenPosition: %w", err)
// 	}

// 	transactor, err := w.getTransactor(0)
// 	if err != nil {
// 		return "", fmt.Errorf("OpenPosition: %w", err)
// 	}

// 	tx, err := PeripheryContract.OpenPosition(transactor, big.NewInt(1), utils.OrderToLibOrder(&order1), signature1, utils.OrderToLibOrder(&order2), signature2, []byte{0})
// 	if err != nil {
// 		return "", fmt.Errorf("OpenPosition: %w", err)
// 	}
// 	return tx.Hash().Hex(), nil
// }

// GetTransactor ...
func (w *Worker) getTransactor(gasLimit uint64, ethC *ethclient.Client) (auth *bind.TransactOpts, err error) {
	// gasLimit = 10000000000 + gasLimit
	privateKey, err := utils.GetPrivateKey(w.PrivateKey)
	if err != nil {
		return nil, err
	}

	nonce, err := ethC.PendingNonceAt(context.Background(), w.config.WorkerAddr)
	if err != nil {
		return nil, err
	}

	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(w.chainID))
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)                // in wei
	auth.GasLimit = uint64(w.config.GasLimit) // in units

	// gasTipCap, err := ethC.SuggestGasTipCap(context.Background())
	// if err != nil {
	// 	return nil, err
	// }

	// gasPrice, err := ethC.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	return nil, err
	// }
	// gasfee := gasPrice.Mul(gasPrice, big.NewInt(int64(gasLimit)))

	// auth.GasFeeCap = gasfee.Add(gasfee, big.NewInt(100000))
	// auth.GasTipCap = gasTipCap.Add(gasTipCap, big.NewInt(100000))
	// auth.GasPrice = w.config.GasPrice

	return auth, nil
}

// getCallOpts....
func getCallOpts() *bind.CallOpts {
	return &bind.CallOpts{
		Pending: true,
	}
}

func (w *Worker) GetChainID() int64 {
	return w.chainID
}

func (w *Worker) GetPositioningContract() common.Address {
	return w.positioningContract
}

func (w *Worker) GetPeripheryContract() common.Address {
	return w.peripheryContract
}

func (w *Worker) SubscribeToLogs(logs chan types.Log) (ethereum.Subscription, error) {
	var StartBlockHeight *big.Int
	lastTxnLog, err := w.DB.FindLatestTxnLog(w.ChainName)
	if err != nil {
		StartBlockHeight = w.config.StartBlockHeight
		if StartBlockHeight.Cmp(big.NewInt(0)) == 0 {
			StartBlockHeight, err = w.GetLatestBlock()
			if err != nil {
				return nil, fmt.Errorf("SubscribeToLogs:%w", err)
			}
		}
	} else {
		StartBlockHeight = big.NewInt(int64(lastTxnLog.BlockHeight))
	}

	// TODO : recheck this
	query := ethereum.FilterQuery{
		Addresses: []common.Address{w.config.MatchingEngine, w.config.GnosisContract}, FromBlock: StartBlockHeight,
	}

	sub, err := w.client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return nil, fmt.Errorf("SubscribeToLogs:%w", err)
	}

	return sub, nil
}

// TODO: LOST tx status check not implemented correctly here
func (w *Worker) GetTransactionReceipt(TransactionHash string) (*types.Receipt, error) {
	waitCount := 5              // TODO add to config
	waitTime := time.Second * 5 //  TODO add to config

	txHash := common.HexToHash(TransactionHash)
	txReceipt, er := w.client.TransactionReceipt(context.Background(), txHash)
	if er != nil {
		for i := 0; i < waitCount; i++ {
			if errors.Is(er, ethereum.NotFound) {
				_, isPending, er := w.client.TransactionByHash(context.Background(), txHash)
				if isPending && (i < waitCount-1) {
					//TODO: return pending txn status here
					return nil, fmt.Errorf("transaction pending") // Pending
				} else if errors.Is(er, ethereum.NotFound) && (i < waitCount-1) {
					return nil, er // Not found
				} else if er != nil {
					return nil, er // Other error in byHash func
				}
				time.Sleep(waitTime)
			} else {
				return nil, er // Other error in transactionReceipt func
			}
		}
	}
	return txReceipt, er
}

func (w *Worker) GetMarkPriceOracle() common.Address {
	return w.markPriceOracle
}

func (w *Worker) GetMarkPrice(index *big.Int) (*big.Int, error) {
	twInterval, _ := new(big.Int).SetString("28800", 10)
	markPriceOracle, _ := MarkPriceOracle.NewMarkPriceOracle(w.markPriceOracle, w.client)
	markTwap, err := markPriceOracle.GetMarkSma(getCallOpts(), twInterval, index)
	return markTwap, err
}

func (w *Worker) GetLastPrice(index *big.Int) (*big.Int, error) {
	markPriceOracle, _ := MarkPriceOracle.NewMarkPriceOracle(w.markPriceOracle, w.client)
	markTwap, err := markPriceOracle.GetLastPrice(getCallOpts(), index)
	return markTwap, err
}

func (w *Worker) GetIndexPrice(index *big.Int) (*big.Int, error) {
	indexPriceOracle, _ := IndexPriceOracle.NewIndexPriceOracle(w.indexPriceOracle, w.client)
	indexTwap, err := indexPriceOracle.GetLastPrice(getCallOpts(), index)
	return indexTwap, err
}

func (w *Worker) GetIndexByBaseToken(baseToken common.Address) (*big.Int, error) {
	markPriceOracle, _ := MarkPriceOracle.NewMarkPriceOracle(w.markPriceOracle, w.client)
	index, err := markPriceOracle.IndexByBaseToken(getCallOpts(), baseToken)
	return index, err
}

func (w *Worker) ValidateOrder(order *database.Order) (bool, error) {
	if order.OrderType == database.ORDER {
		return true, nil
	}
	var baseToken string
	if order.IsShort {
		baseToken = order.MakeAsset().VirtualToken
	} else {
		baseToken = order.TakeAsset().VirtualToken
	}
	index, err := w.GetIndexByBaseToken(common.HexToAddress(baseToken))
	if err != nil {
		return false, err
	}

	var triggeredPrice *big.Int
	switch order.OrderType {
	case database.STOP_LOSS_INDEX_PRICE, database.TAKE_PROFIT_INDEX_PRICE:
		triggeredPrice, err = w.GetIndexPrice(index)
		if err != nil {
			return false, err
		}
	case database.STOP_LOSS_MARK_PRICE, database.TAKE_PROFIT_MARK_PRICE:
		triggeredPrice, err = w.GetMarkPrice(index)
		if err != nil {
			return false, err
		}
	case database.STOP_LOSS_LAST_PRICE, database.TAKE_PROFIT_LAST_PRICE:
		triggeredPrice, err = w.GetLastPrice(index)
		if err != nil {
			return false, err
		}
	default:
		return false, nil
	}

	triggerPrice, _ := new(big.Int).SetString(order.TriggerPrice, 10)

	return validateOrderTriggerPrice(order, triggeredPrice, triggerPrice), nil
}

func validateOrderTriggerPrice(order *database.Order, triggeredPrice, triggerPrice *big.Int) bool {
	isStopLoss := _checkLimitOrderType(order.OrderType, true)
	result := triggeredPrice.Cmp(triggerPrice)

	if order.IsShort == isStopLoss {
		// Matching trigger price for a stop-loss order or not matching trigger price for a take-profit order
		return result <= 0
	} else {
		// Matching trigger price for a take-profit order or not matching trigger price for a stop-loss order
		return result >= 0
	}
}

func _checkLimitOrderType(orderType string, isStopLoss bool) bool {
	if isStopLoss {
		return orderType == database.STOP_LOSS_INDEX_PRICE || orderType == database.STOP_LOSS_LAST_PRICE || orderType == database.STOP_LOSS_MARK_PRICE
	}
	return orderType == database.TAKE_PROFIT_INDEX_PRICE || orderType == database.TAKE_PROFIT_LAST_PRICE || orderType == database.TAKE_PROFIT_MARK_PRICE
}
