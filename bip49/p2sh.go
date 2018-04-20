package bip49

import (
	"errors"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
)

var (
	// possible prefixes for a P2SH key (Used for key validation)
	p2shHumanPre = map[string]bool{"ypub": true, "yprv": true, "upub": true, "uprv": true}
)

func GetExtPrvForP2SHAccount(seed []byte, accountIndex uint32, network Network) (string, error) {

	m, err := GetExtendedMasterPrivateKeyFromSeedBytes(seed, network)
	if err != nil {
		return "", err
	}

	return GetBIP49AccountKey(m, network, accountIndex, true)
}

func GetExtPubForP2SHAccount(seed []byte, accountIndex uint32, network Network) (string, error) {

	m, err := GetExtendedMasterPrivateKeyFromSeedBytes(seed, network)
	if err != nil {
		return "", err
	}

	return GetBIP49AccountKey(m, network, accountIndex, false)
}

func GetP2SHAddressForIndex(accountKey string, isChange bool, addressIndex uint32, network Network) (string, error) {

	validPre := p2shHumanPre[accountKey[:4]]
	if !validPre {
		return "", errors.New("Key does not start with a P2SH prefix")
	}

	addressType := ExternalAddress
	if isChange {
		addressType = ChangeAddress
	}

	k, err := GetAccountAddressKey(accountKey, addressType, addressIndex)
	if err != nil {
		return "", err
	}

	pk, err := k.ECPubKey()
	if err != nil {
		return "", err
	}

	keyHash := btcutil.Hash160(pk.SerializeCompressed())
	scriptSig, err := txscript.NewScriptBuilder().AddOp(txscript.OP_0).AddData(keyHash).Script()
	if err != nil {
		return "", err
	}

	segAddr, err := btcutil.NewAddressScriptHash(scriptSig, network.ChainConfigParams())
	if err != nil {
		return "", err
	}
	return segAddr.EncodeAddress(), nil
}
