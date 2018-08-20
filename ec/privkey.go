// Package ec provides ...
package ec

import "github.com/duminghui/go-bip32/utils/bytes"
import "math/big"

const PrivKeyBytesLen = 32

type PrivateKey struct {
	PublicKey
	D *big.Int
}

func PrivKeyFromBytes(bytes []byte) (*PrivateKey, *PublicKey) {
	x, y := secp256k1.ScalarBaseMult(bytes)
	privKey := &PrivateKey{
		PublicKey: PublicKey{
			X: x,
			Y: y,
		},
		D: new(big.Int).SetBytes(bytes),
	}
	return privKey, &privKey.PublicKey
}

func (privKey *PrivateKey) Serialize() []byte {
	return bytes.PaddedBytes(32, privKey.D.Bytes())
}
