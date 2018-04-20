package bip49

import (
	"encoding/hex"
	"testing"
)

func TestFromPrvKey(t *testing.T) {

	var actual interface{}
	var expected interface{}

	seedStr := "337b3fba0f292fa71c0a997d13ab9f8e2b09ac901ece31c2e3f0d56347e85a3205729f513f0167d449f331cc541e701bdd5b2c4363df7da926039b1b123c3b1c"
	seed, _ := hex.DecodeString(seedStr)
	network := BTCMainnet
	accountIndex := uint32(0)
	isChange := false

	accountPrvKey, err := GetExtPrvForP2SHAccount(seed, accountIndex, network)
	if err != nil {
		t.Fatal(err)
	}

	address0, err := GetP2SHAddressForIndex(accountPrvKey, isChange, uint32(0), network)
	if err != nil {
		t.Fatal(err)
	}

	expected = "39XTm1FNWYeuPErH51XR4v9GusMuSc1HUy"
	actual = address0
	if actual != expected {
		t.Errorf("got %v where I want %v", actual, expected)
	}

	address1, err := GetP2SHAddressForIndex(accountPrvKey, isChange, uint32(1), network)
	if err != nil {
		t.Fatal(err)
	}
	expected = "3A94BMRogsAcnN5jdpyKXvo9rWyKNW7wVd"
	actual = address1
	if actual != expected {
		t.Errorf("got %v where I want %v", actual, expected)
	}
}

func TestFromPubKey(t *testing.T) {

	var actual interface{}
	var expected interface{}

	seedStr := "337b3fba0f292fa71c0a997d13ab9f8e2b09ac901ece31c2e3f0d56347e85a3205729f513f0167d449f331cc541e701bdd5b2c4363df7da926039b1b123c3b1c"
	seed, _ := hex.DecodeString(seedStr)
	network := BTCMainnet
	accountIndex := uint32(0)
	isChange := false

	accountPubKey, err := GetExtPubForP2SHAccount(seed, accountIndex, network)
	if err != nil {
		t.Fatal(err)
	}

	address0, err := GetP2SHAddressForIndex(accountPubKey, isChange, uint32(0), network)
	if err != nil {
		t.Fatal(err)
	}

	expected = "39XTm1FNWYeuPErH51XR4v9GusMuSc1HUy"
	actual = address0
	if actual != expected {
		t.Errorf("got %v where I want %v", actual, expected)
	}

	address1, err := GetP2SHAddressForIndex(accountPubKey, isChange, uint32(1), network)
	if err != nil {
		t.Fatal(err)
	}
	expected = "3A94BMRogsAcnN5jdpyKXvo9rWyKNW7wVd"
	actual = address1
	if actual != expected {
		t.Errorf("got %v where I want %v", actual, expected)
	}
}

func TestTestnet(t *testing.T) {

	var actual interface{}
	var expected interface{}

	seedStr := "337b3fba0f292fa71c0a997d13ab9f8e2b09ac901ece31c2e3f0d56347e85a3205729f513f0167d449f331cc541e701bdd5b2c4363df7da926039b1b123c3b1c"
	seed, _ := hex.DecodeString(seedStr)
	network := BTCTestnet
	accountIndex := uint32(0)
	isChange := false

	accountPubKey, err := GetExtPubForP2SHAccount(seed, accountIndex, network)
	if err != nil {
		t.Fatal(err)
	}

	address0, err := GetP2SHAddressForIndex(accountPubKey, isChange, uint32(0), network)
	if err != nil {
		t.Fatal(err)
	}

	expected = "2N7462xQhpvRc6H5s3ATN35EbULx9CSeGyP"
	actual = address0
	if actual != expected {
		t.Errorf("got %v where I want %v", actual, expected)
	}

	address1, err := GetP2SHAddressForIndex(accountPubKey, isChange, uint32(1), network)
	if err != nil {
		t.Fatal(err)
	}
	expected = "2N1H4D66eYZND3n5YaTeqUV2cznGgo8UGtH"
	actual = address1
	if actual != expected {
		t.Errorf("got %v where I want %v", actual, expected)
	}
}
