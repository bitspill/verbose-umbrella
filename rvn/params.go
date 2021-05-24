package rvn

import (
	"github.com/bitspill/flod/chaincfg"
	"github.com/bitspill/flod/wire"
)

// MainNetParams defines the network parameters for the main Raven network.
var MainNetParams = chaincfg.Params{
	Net: wire.BitcoinNet(0x4e564152),

	// Address encoding magics
	PubKeyHashAddrID: 60,  // starts with ?
	ScriptHashAddrID: 122, // starts with ?
	PrivateKeyID:     128, // starts with ? (uncompressed) or ? (compressed)
}

// TestNetParams defines the network parameters for the main Raven network.
var TestNetParams = chaincfg.Params{
	Net: wire.BitcoinNet(0x544e5652),

	// Address encoding magics
	PubKeyHashAddrID: 111, // starts with ?
	ScriptHashAddrID: 196, // starts with ?
	PrivateKeyID:     239, // starts with ? (uncompressed) or ? (compressed)
}

func init() {
	err := chaincfg.Register(&MainNetParams)
	if err != nil {
		panic("failed to register rvn mainnet: " + err.Error())
	}
	err = chaincfg.Register(&TestNetParams)
	if err != nil {
		panic("failed to register rvn testnet: " + err.Error())
	}
}
