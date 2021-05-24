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
	adr := "RR9cRYdfY8idD1vEw52wzTYSsJV1a5ZfXG"
	sig := "H3STxLqULkya+nRYAWkg8OIwBMrsL68UJ+ivLYGSaFy/Cw7AxzKVDMEmcrDwOyLPjiz5iLdku2A2SxlgyEwpOB0="
	msg := "452523db2e55bc27e63872f40a8c6e6562ab72816dd7039a7e244ccbb04d0f69-3993842283-1-1621808213"
	valid, err := CheckSignature(adr, sig, msg)
	if err != nil {
		t.Error(err)
	}
	if !valid {
		t.Fail()
	}

	// TestNet
	config.SetTestnet(true)
	adr = "mpLQjfK79b7CCV4VMJWEWAj5Mpx8Up5zxB"
	sig = "H5vCbG+WhOeOPJ3jf6oux/1oSjkuIGZigCw4NW+A0/fSDlgdO4fMq0SWSfx7gUMB9kuG+t/0BQxtXaTCr7v9fGM="
	msg = "This is just a test message"
	valid, err = CheckSignature(adr, sig, msg)
	if err != nil {
		t.Error(err)
	}
	if !valid {
		t.Fail()
	}

	// restore pre-test setting
	config.SetTestnet(testnet)
}
