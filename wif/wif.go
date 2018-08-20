// Package wif provides ...
package wif

import (
	"github.com/duminghui/go-bip32/ec"
	"github.com/duminghui/go-bip32/utils/base58"
	"github.com/duminghui/go-bip32/utils/hash"
)

type WIF struct {
	PrivKey *ec.PrivateKey
}

const version = 0x80

func NewWIF(privKey *ec.PrivateKey) *WIF {
	return &WIF{
		PrivKey: privKey,
	}
}

// Encode return WIF-uncompressed
func (wif *WIF) Encode() string {
	return wif.encode(false)
}

// EncodeCompressed return WIF-compressed
func (wif *WIF) EncodeCompressed() string {
	return wif.encode(true)
}

func (wif *WIF) encode(compressPubKey bool) string {
	encodeLen := 1 + ec.PrivKeyBytesLen + 4
	if compressPubKey {
		encodeLen++
	}
	b := make([]byte, 0, encodeLen)
	b = append(b, version)
	b = append(b, wif.PrivKey.Serialize()...)
	if compressPubKey {
		b = append(b, 0x01)
	}
	checksum := hash.DoubleHash256(b)[:4]
	b = append(b, checksum...)
	return base58.Encode(b)
}
