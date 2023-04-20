package validation

import (
	"testing"

	"github.com/volmexfinance/relayers/internal/testlib"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/volmexfinance/relayers/relayer-srv/db"
	"github.com/volmexfinance/relayers/relayer-srv/utils"
	// "github.com/volmexfinance/relayers/relayer-srv/worker/abi/positioning"
)

func TestValidateMatchedOrder(t *testing.T) {
	testDb := testlib.NewTestDB(t)

	order1 := db.Order{
		OrderType:    "0xf555eb98",
		Trader:       "0x401f0B1c51A7048D3dB9A8ca4E9a370e563E0Fb9",
		Deadline:     87654321987654,
		Assets:       []db.Assets{{VirtualToken: "0xb866E40cA0C89c5E7feC0E102B2f371d7602bc9d", Value: "2000000000000000000"}, {VirtualToken: "0x5a0cB5D14c17a5Faa5655Ba39235445cAED19a90", Value: "2000000000000000000"}},
		Price:        1,
		Salt:         "44",
		TriggerPrice: "0",
		Sign:         "",
		IsShort:      false,
		Status:       db.MatchedStatusInit,
		Fills:        "0",
	}
	order2 := db.Order{
		OrderType:    "0xf555eb98",
		Trader:       "0x77a18F00CE53c004337b4A8b7A9a585AFFDEeD5e",
		Deadline:     87654321987654,
		Assets:       []db.Assets{{VirtualToken: "0x5a0cB5D14c17a5Faa5655Ba39235445cAED19a90", Value: "2000000000000000000"}, {VirtualToken: "0xb866E40cA0C89c5E7feC0E102B2f371d7602bc9d", Value: "2000000000000000000"}},
		Price:        1,
		Salt:         "44",
		TriggerPrice: "0",
		Sign:         "",
		IsShort:      true,
		Status:       db.MatchedStatusInit,
		Fills:        "0", //TODO: check why 0
	}

	if err := ValidateMatchedOrder(order1, order2, testDb, 2); err != nil {
		t.Errorf("Failed to validate order%q", err)
	}
}

func TestValidateThreshold(t *testing.T) {
	testDb := testlib.NewTestDB(t)

	order1 := db.Order{
		OrderType:    "0xf555eb98",
		Trader:       "0x401f0B1c51A7048D3dB9A8ca4E9a370e563E0Fb9",
		Deadline:     87654321987654,
		Assets:       []db.Assets{{VirtualToken: "0xb866E40cA0C89c5E7feC0E102B2f371d7602bc9d", Value: "2000000000000000000"}, {VirtualToken: "0x5a0cB5D14c17a5Faa5655Ba39235445cAED19a90", Value: "2000000000000000000"}},
		Price:        1,
		Salt:         "44",
		TriggerPrice: "0",
		Sign:         "",
		IsShort:      false,
		Status:       db.MatchedStatusInit,
	}
	order2 := db.Order{
		OrderType:    "0xf555eb98",
		Trader:       "0x77a18F00CE53c004337b4A8b7A9a585AFFDEeD5e",
		Deadline:     87654321987654,
		Assets:       []db.Assets{{VirtualToken: "0x5a0cB5D14c17a5Faa5655Ba39235445cAED19a90", Value: "2000000000000000000"}, {VirtualToken: "0xb866E40cA0C89c5E7feC0E102B2f371d7602bc9d", Value: "2000000000000000000"}},
		Price:        1,
		Salt:         "44",
		TriggerPrice: "0",
		Sign:         "",
		IsShort:      true,
		Status:       db.MatchedStatusInit,
	}

	if err := ValidateThreshold([]*db.Order{&order1, &order2}, testDb); err != nil {
		t.Errorf("Failed to validate threshold")
	}
}

func TestVerifyOrderSignature(t *testing.T) {

	positioningContract := "0x7CcF23f53B5886347D90d6Ea3DFd1B54CeD09254"
	order1 := db.Order{
		OrderType:    "0xf555eb98",
		Trader:       "0x401f0B1c51A7048D3dB9A8ca4E9a370e563E0Fb9",
		Deadline:     87654321987654,
		Assets:       []db.Assets{{VirtualToken: "0xb866E40cA0C89c5E7feC0E102B2f371d7602bc9d", Value: "2000000000000000000"}, {VirtualToken: "0x5a0cB5D14c17a5Faa5655Ba39235445cAED19a90", Value: "2000000000000000000"}},
		Price:        1,
		Salt:         "44",
		TriggerPrice: "0",
		Sign:         "",
		IsShort:      false,
		Status:       db.MatchedStatusInit,
	}

	privkey1, err := utils.GetPrivateKey("15598fdd84647189880f752f16796d9935f7bb1e58e5243375b8e61f8f1a06be")
	if err != nil {
		println("err privkey", err.Error())
	}

	hash1, err := utils.EncodeOrderStruct(order1, 1, positioningContract)
	if err != nil {
		println("err", err.Error())
	}

	signature1, er := crypto.Sign(hash1, privkey1)
	if er != nil {
		println("err in signing", er.Error())
	}

	order1.Sign = hexutil.Encode(signature1)

	// publicKey := privkey1.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	panic("error casting public key to ECDSA")
	// }

	// publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// fmt.Println("publicKeyECDSA",publicKeyECDSA)

	if err := VerifyOrderSignature(order1, 1, positioningContract); err != nil {
		t.Errorf("Failed to Verify Order Signature%q", err)
	}

}
