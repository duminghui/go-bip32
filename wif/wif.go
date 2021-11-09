// Package wif provides ...
package wif

import (
	"go-bip32/ec"
	"go-bip32/utils/base58"
	"go-bip32/utils/hash"
)

type WIF struct {
	privKey *ec.PrivateKey
	version byte
}

func NewWIF(privKey *ec.PrivateKey, version byte) *WIF {
	return &WIF{
		privKey: privKey,
		version: version,
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
	b = append(b, wif.version)
	b = append(b, wif.privKey.Serialize()...)
	if compressPubKey {
		b = append(b, 0x01)
	}
	checksum := hash.DoubleHash256(b)[:4]
	b = append(b, checksum...)
	return base58.Encode(b)
}
