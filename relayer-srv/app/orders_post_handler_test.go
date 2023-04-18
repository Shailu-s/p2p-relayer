package app

import (
	"github.com/volmexfinance/relayers/relayer-srv/db"
	"reflect"
	"testing"
)

func TestToDBOrder(t *testing.T) {
	newOrder := reqOrder{ // TODO : fill values
		OrderType: "0xf555eb98",
		Deadline:  "87654321987654",
		Trader:    "0x401f0B1c51A7048D3dB9A8ca4E9a370e563E0Fb9",
		MakeAsset: reqAsset{
			VirtualToken: "0x5a0cB5D14c17a5Faa5655Ba39235445cAED19a90",
			Value:        "2000000000000000000",
		},
		TakeAsset: reqAsset{
			VirtualToken: "0xb866E40cA0C89c5E7feC0E102B2f371d7602bc9d",
			Value:        "2000000000000000000",
		},
		Salt:         "44",
		TriggerPrice: "0",
		IsShort:      false,
		Signature:    "",
	}

	dbOrderWant := &db.Order{}
	dbOrderGot, err := toDBOrder(newOrder, "")
	if err != nil {
		t.Errorf("toDBOrder failed with error :%q", err)
	}

	if reflect.TypeOf(dbOrderGot) != reflect.TypeOf(dbOrderWant) {
		t.Errorf("got %q, wanted db.Order", reflect.TypeOf(dbOrderGot))
	}
}
