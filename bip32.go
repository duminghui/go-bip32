// Package bip32 provides ...
package bip32

import (
	"crypto/hmac"
	"crypto/sha512"
	"errors"
	"math/big"

	"github.com/duminghui/go-bip32/crypto/elliptic"
)

// https://github.com/btcsuite/btcutil/blob/master/hdkeychain/extendedkey.go

const (
	//HardenedKeyStart hardended key starts.
	HardenedKeyStart = 0x80000000 // 2^31

	// version(4 bytes) || depth(1 byte) || fingerprint(4 bytes) || child number(4 bytes) || chaincode(32 bytes) || pub/pri key data(33 bytes)
	serializedKeyLen = 78

	// max child number
	maxChildNum = 0xFF

	minSeedBytes = 16 //128 bits

	maxSeedBytes = 64 // 512 bits
)

var (
	// ErrInvalidSeedLen seed Len error
	ErrInvalidSeedLen = errors.New("seed lenght must be between 128 and 512 bits")
	// ErrUnusableSeed describes an error in which the provided seed is not
	// usable due to the derived key falling outside of the valid range for
	// secp256k1 private keys.  This error indicates the caller must choose
	// another seed.
	ErrUnusableSeed = errors.New("unusable seed")
)

var (
	masterKey        = []byte("Bitcoin seed")
	keyVerMainNetPri = []byte{0x04, 0x88, 0xad, 0xe4} // starts with xprv
	keyVerMainNetPub = []byte{0x04, 0x88, 0xb2, 0x1e} // starts with xpub
)

// ExtendedKey private/public key data
type ExtendedKey struct {
	version     []byte // 4 bytes
	depth       byte   // 1 byte
	fingerprint []byte // 4 bytes
	childNum    []byte // 4 bytes
	chainCode   []byte // 32 bytes
	key         []byte // 33 bytes
	isPrivate   bool
}

//NewMasterKey create a new master key data from seed
func NewMasterKey(seed []byte) (*ExtendedKey, error) {
	if len(seed) < minSeedBytes || len(seed) > maxSeedBytes {
		return nil, ErrInvalidSeedLen
	}
	// I = HMAC-SHA512(Key = "Bitcoin seed", Data = S)
	hmac512 := hmac.New(sha512.New, masterKey)
	hmac512.Write(seed)
	lr := hmac512.Sum(nil)

	// Split "I" into two 32-byte sequences Il and Ir where:
	//   Il = master secret key
	//   Ir = master chain code
	secretKey := lr[:32]
	chainCode := lr[32:]

	secretKeyNum := new(big.Int).SetBytes(secretKey)
	if secretKeyNum.Cmp(elliptic.Secp265k1().Params().N) >= 0 || secretKeyNum.Sign() == 0 {
		return nil, ErrUnusableSeed
	}

	return &ExtendedKey{
		version:     keyVerMainNetPri,
		depth:       0,
		fingerprint: []byte{0x00, 0x00, 0x00, 0x00},
		childNum:    []byte{0x00, 0x00, 0x00, 0x00},
		chainCode:   chainCode,
		key:         secretKey,
		isPrivate:   true,
	}, nil
}
