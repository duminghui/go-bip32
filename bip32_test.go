package bip32

import (
	"fmt"
	"testing"

	"github.com/duminghui/go-bip39"
)

func TestNewMasterKey(t *testing.T) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		panic(err)
	}
	fmt.Printf("entropy:%x\n", entropy)
	mnemonic, err := bip39.NewMnemonic(entropy)
	mnemonic = "lemon speak wish predict athlete fringe ritual bonus ensure clay school goat"
	fmt.Printf("mnemonic:%s\n", mnemonic)
	seed, err := bip39.NewSeedWithValidMnemonic(mnemonic, "")
	fmt.Printf("seed:%x, len:%d\n", seed, len(seed))
	// seedNum, _ := new(big.Int).SetString("000102030405060708090a0b0c0d0e0f", 16)
	// seed = seedNum.Bytes()
	// seed = bytes.PaddedBytes(16, seed)
	// fmt.Printf("seed2: %x\n", seed)
	masterKey, err := NewMasterKey(seed)
	fmt.Println(err, len(seed))
	// fmt.Printf("MK version:%x\n", masterKey.version)
	// fmt.Printf("MK childNum:%x\n", masterKey.childNum)
	// fmt.Printf("MK key:%x\n", masterKey.key)
	// fmt.Printf("MK chainCode:%x\n", masterKey.chainCode)
	// fmt.Printf("MK fingerprint:%x\n", masterKey.parentFP)
	// fmt.Printf("MK depth:%x\n", masterKey.depth)
	// fmt.Printf("MK isPrivate:%v\n", masterKey.isPrivate)
	b58 := masterKey.B58Serialize()
	fmt.Println("MK b58:", b58)
	b58Vaild := "xprv9s21ZrQH143K3dRqV7LTj7uFvjMFiai3HV7Ytoj7Jg3mEArCCk2ujozgMSoip9FTWH4YwjTnjayPooxVesrtHst8qo8xZCnG1MazdLfBJBR"
	fmt.Println(b58 == b58Vaild)
	masterPubKey := masterKey.Neuter()
	b58pubKey := masterPubKey.B58Serialize()
	fmt.Println("masterPubKey:", b58pubKey)
	b58pubValid := "xpub661MyMwAqRbcG7WJb8sU6FqzUmBk83Rtei39hC8is1ak6yBLkHMAHcKACgmyxr6KLGhLSM4Qr74R8meHNcttv4bmPMDC8EwB9vJ3HRGaC53"
	fmt.Println(b58pubKey == b58pubValid)
	fmt.Println(len(masterKey.key))
	fmt.Println(len("1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd"))
	// bigNum := new(big.Int).SetBytes(masterKey.key)
	fmt.Println((fmt.Sprintf("%x", masterKey.key)))
	fmt.Println((fmt.Sprintf("%x", masterKey.pubKeyBytes())))
}
