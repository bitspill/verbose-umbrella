package rvn

import (
	"github.com/bitspill/flod/chaincfg"
)

// RvnMainNetParams defines the network parameters for the main Raven network.
var RvnMainNetParams = chaincfg.Params{
	// Address encoding magics
	PubKeyHashAddrID: 60,  // starts with L
	ScriptHashAddrID: 122, // starts with M
	PrivateKeyID:     128, // starts with 6 (uncompressed) or T (compressed)
}

// RvnTestNetParams defines the network parameters for the main Raven network.
var RvnTestNetParams = chaincfg.Params{
	// Address encoding magics
	PubKeyHashAddrID: 111, // starts with L
	ScriptHashAddrID: 196, // starts with M
	PrivateKeyID:     239, // starts with 6 (uncompressed) or T (compressed)
}
