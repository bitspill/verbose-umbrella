package rvn

import (
	"testing"

	"github.com/oipwg/oip/config"
)

func TestCheckSignature(t *testing.T) {
	// save setting to restore post-test
	testnet := config.IsTestnet()

	// MainNet
	config.SetTestnet(false)
	// ToDo: add mainnet test case

	// TestNet
	config.SetTestnet(true)
	adr := "mpLQjfK79b7CCV4VMJWEWAj5Mpx8Up5zxB"
	sig := "H5vCbG+WhOeOPJ3jf6oux/1oSjkuIGZigCw4NW+A0/fSDlgdO4fMq0SWSfx7gUMB9kuG+t/0BQxtXaTCr7v9fGM="
	msg := "This is just a test message"
	valid, err := CheckSignature(adr, sig, msg)
	if err != nil {
		t.Error(err)
	}
	if !valid {
		t.Fail()
	}

	// restore pre-test setting
	config.SetTestnet(testnet)
}
