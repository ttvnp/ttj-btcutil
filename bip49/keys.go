package bip49

import (
	"encoding/hex"

	"github.com/ttvnp/ttj-btcutil/hdkeychain"
)

// AddressType used in BIP44
type AddressType uint32

// HardenedKeyZeroIndex is the index where hardended keys start
const (
	HardenedKeyZeroIndex = 0x80000000 // 2^31

	// BIP49Purpose P2WPKH-nested-in-P2SH purpose
	// see: https://github.com/bitcoin/bips/blob/master/bip-0049.mediawiki
	Purpose uint32 = 49

	ExternalAddress AddressType = 0
	ChangeAddress   AddressType = 1
)

func GetExtendedMasterPrivateKeyFromSeedHex(seed string, network Network) (privateKey *hdkeychain.ExtendedKey, err error) {
	pk, err := hex.DecodeString(seed)
	if err != nil {
		return nil, err
	}
	return GetExtendedMasterPrivateKeyFromSeedBytes(pk, network)
}

func GetExtendedMasterPrivateKeyFromSeedBytes(seed []byte, network Network) (privateKey *hdkeychain.ExtendedKey, err error) {
	return hdkeychain.NewMaster(seed, network.ChainConfigParams().HDPrivateKeyID)
}
func GetExtendedKeyFromString(xKey string) (key *hdkeychain.ExtendedKey, err error) {
	return hdkeychain.NewKeyFromString(xKey)
}

func GetBIP49AccountKey(masterKey *hdkeychain.ExtendedKey, network Network, accountIndex uint32, includePrivateKey bool) (key string, err error) {
	return getAccountKeyWithPurpose(masterKey, Purpose, network, network.CoinType(), accountIndex, includePrivateKey)
}

func getAccountKeyWithPurpose(masterKey *hdkeychain.ExtendedKey, purpose uint32, network Network, coinType CoinType, accountIndex uint32, includePrivateKey bool) (key string, err error) {
	purposeK, err := masterKey.Child(HardenedKeyZeroIndex + purpose)
	if err != nil {
		return "", err
	}

	coinTypeK, err := purposeK.Child(HardenedKeyZeroIndex + coinType.Uint32())
	if err != nil {
		return "", err
	}

	accountKey, err := coinTypeK.Child(HardenedKeyZeroIndex + accountIndex)
	if err != nil {
		return "", err
	}

	if includePrivateKey {
		return accountKey.String(), nil
	}

	pub, err := accountKey.Neuter(network.ChainConfigParams().HDPublicKeyID)
	if err != nil {
		return "", err
	}

	return pub.String(), nil
}

func GetAccountAddressKey(xKey string, change AddressType, addressIndex uint32) (key *hdkeychain.ExtendedKey, err error) {
	account, err := GetExtendedKeyFromString(xKey)
	if err != nil {
		return nil, err
	}

	changeKey, err := account.Child(uint32(change))
	if err != nil {
		return nil, err
	}

	return changeKey.Child(addressIndex)
}
