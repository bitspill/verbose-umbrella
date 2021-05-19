package oip5

import (
	"fmt"
	"testing"

	"github.com/bitspill/flod/chaincfg/chainhash"
	"github.com/oipwg/proto/go/pb_oip"
	"github.com/oipwg/proto/go/pb_oip5/pb_templates/livenet"
)

func TestAsset(t *testing.T) {
	scs := &livenet.SimpleCoinSale{
		Destination: "ALBAWABA2021",
		Scale:       1,
		Amount:      1,
		Coin:        pb_oip.TxidFromString("0000000000000000000000000000000000000000000000000000000000000002"),
		Duration:    0,
		Indefinite:  true,
	}

	addr := "RR9cRYdfY8idD1vEw52wzTYSsJV1a5ZfXG"

	paid, err := checkRvnAsset(scs, addr)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(paid)
}

// Random TX f341eda1dc0070fc2799f8f29552865e9bfe6812c6b944647ff6689407697346
func TestTx(t *testing.T) {
	scs := &livenet.SimpleCoinSale{
		Destination: "RAq9yKBkEgAPRhUeNS3XcusgvsBf9gsSmW",
		Scale:       1,
		Amount:      100,
		Coin:        pb_oip.TxidFromString("0000000000000000000000000000000000000000000000000000000000000001"),
		Duration:    0,
		Indefinite:  true,
	}

	addr := "RCc4qqhSWgSSfxBSWtZkNkV3X5nJR1DVtz"

	txid := &chainhash.Hash{}
	_ = chainhash.Decode(txid, "f341eda1dc0070fc2799f8f29552865e9bfe6812c6b944647ff6689407697346")

	paid, err := checkRvnPayment(txid, scs, addr)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(paid)
}
