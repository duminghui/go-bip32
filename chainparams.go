// Package bip32 provides ...
package bip32

type ChainParams struct {
	PubKeyHashAddrID byte
	PrivateKeyID     byte
	HDPrivateKeyID   [4]byte
	HDPublicKeyID    [4]byte
	HDCoinType       uint32
}

var BTCMainNetParams = ChainParams{
	PubKeyHashAddrID: 0x00,                            // 1
	PrivateKeyID:     0x80,                            // 5(uncompressed) or K (compressed)
	HDPrivateKeyID:   [4]byte{0x04, 0x88, 0xad, 0xe4}, //xprv
	HDPublicKeyID:    [4]byte{0x04, 0x88, 0xb2, 0x1e}, //xpub
	HDCoinType:       0,
}
