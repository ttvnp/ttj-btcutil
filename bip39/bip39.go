package bip39

import (
	"github.com/tyler-smith/go-bip39"
)


func NewSeedWithErrorChecking(mnemonic , password string) ([]byte, error) {
	return bip39.NewSeedWithErrorChecking(mnemonic, password)
}

func NewSeed(mnemonic, password string) []byte {
	return bip39.NewSeed(mnemonic, password)
}
