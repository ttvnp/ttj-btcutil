package bip39

import (
	"testing"

	"encoding/hex"

	"github.com/tyler-smith/go-bip39"
)

func TestNewSeedWithErrorChecking(t *testing.T) {

	var actual interface{}
	var expected interface{}

	var mnemonic string
	var password string
	mnemonic = "task another nasty today believe cattle creek apple thought amused cabbage exact clog hurt few radio skate hurt dutch wasp lunch answer snake save"
	password = ""
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, password)
	if err != nil {
		t.Error(err)
	}

	expected = "adf82588662959446c460e159b97e39cf43559d89122a49960d33ca0fbcb1f8369230c3da4977284498cc5463141b85c0e3661d3779d6e4ad3b116be6e2bfeb4"
	actual = hex.EncodeToString(seed)
	if actual != expected {
		t.Errorf("got %v where I want %v", actual, expected)
	}
}
