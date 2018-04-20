package bip49

import (
	"errors"

	"github.com/btcsuite/btcd/chaincfg"
)

// Network references a distributed network endpoint
type Network int16
type CoinType uint32

const (
	// BTCTestnet Bitcoin test network (3)
	BTCTestnet Network = 0

	// BTCMainnet Bitcoin main network
	BTCMainnet Network = 1

	// BTCCoinType (Full list of coin types available here: https://github.com/satoshilabs/slips/blob/master/slip-0044.md)
	BTCCoinType        CoinType = 0
	BTCTestnetCoinType CoinType = 1
)

var (
	// Version bytes of master public/private keys indicate what type of output script should be used.
	privP2WPKHinP2SHVer         = [4]byte{0x04, 0x9d, 0x78, 0x78}
	pubP2WPKHinP2SHVer          = [4]byte{0x04, 0x9d, 0x7c, 0xb2}
	privP2WPKHinP2SHVer4Testnet = [4]byte{0x04, 0x4A, 0x4E, 0x28}
	pubP2WPKHinP2SHVer4Testnet  = [4]byte{0x04, 0x4A, 0x52, 0x62}
)

func (net Network) ChainConfigParams() *chaincfg.Params {
	switch net {
	case BTCMainnet:
		ret := &chaincfg.MainNetParams
		ret.HDPrivateKeyID = privP2WPKHinP2SHVer
		ret.HDPublicKeyID = pubP2WPKHinP2SHVer
		return ret

	case BTCTestnet:
		ret := &chaincfg.TestNet3Params
		ret.HDPrivateKeyID = privP2WPKHinP2SHVer4Testnet
		ret.HDPublicKeyID = pubP2WPKHinP2SHVer4Testnet
		return ret
	default:
		panic(errors.New("Unknown network specified"))
	}
}

func (net Network) CoinType() CoinType {
	switch net {
	case BTCMainnet:
		return BTCCoinType
	case BTCTestnet:
		return BTCTestnetCoinType
	default:
		panic(errors.New("Unknown network specified"))
	}
}

func (ct CoinType) Uint32() uint32 {
	return uint32(ct)
}
