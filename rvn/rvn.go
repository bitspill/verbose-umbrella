package rvn

import (
	"github.com/bitspill/flosig"
	"github.com/bitspill/floutil"
	"github.com/pkg/errors"

	"github.com/oipwg/oip/config"
)

func CheckAddress(address string) (bool, error) {
	var err error
	if config.IsTestnet() {
		_, err = floutil.DecodeAddress(address, &TestNetParams)
	} else {
		_, err = floutil.DecodeAddress(address, &MainNetParams)
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func CheckSignature(address, signature, message string) (bool, error) {
	var ok bool
	var err error
	if config.IsTestnet() {
		ok, err = flosig.CheckSignature(address, signature, message, "Raven", &TestNetParams)
	} else {
		ok, err = flosig.CheckSignature(address, signature, message, "Raven", &MainNetParams)
	}
	if !ok && err == nil {
		err = errors.New("bad signature")
	}

	return ok, err
}
