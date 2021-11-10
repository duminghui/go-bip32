// Package util provides ...
package util

import (
	"errors"

	"github.com/duminghui/go-bip32/d/chaincfg"
	"github.com/duminghui/go-bip32/util/base58"
	"golang.org/x/crypto/ripemd160"
)

// encodeAddress returns a human-readable payment address given a ripemd160 hash
// and netID which encodes the bitcoin network and address type.  It is used
// in both pay-to-pubkey-hash (P2PKH) and pay-to-script-hash (P2SH) address
// encoding.
func encodeAddress(hash160 []byte, netID byte) string {
	// Format is 1 byte for a network and address class (i.e. P2PKH vs
	// P2SH), 20 bytes for a RIPEMD160 hash, and 4 bytes of checksum.
	return base58.CheckEncode(hash160[:ripemd160.Size], netID)
}

//// P2PKH P2SH address encoding
//func encodeAddress(hash160 []byte, version byte) string {
//	input := make([]byte, 21)
//	input[0] = version
//	copy(input[1:], hash160)
//	return base58.CheckEncode(input)
//}

// AddressPubKeyHash is an Address for a pay-to-pubkey-hash (P2PKH)
// transaction.
type AddressPubKeyHash struct {
	hash  [ripemd160.Size]byte
	netID byte
}

// NewAddressPubKeyHash returns a new AddressPubKeyHash.  pkHash mustbe 20
// bytes.
func NewAddressPubKeyHash(pkHash []byte, net *chaincfg.Params) (*AddressPubKeyHash, error) {
	return newAddressPubKeyHash(pkHash, net.PubKeyHashAddrID)
}

// newAddressPubKeyHash is the internal API to create a pubkey hash address
// with a known leading identifier byte for a network, rather than looking
// it up through its parameters.  This is useful when creating a new address
// structure from a string encoding where the identifer byte is already
// known.
func newAddressPubKeyHash(pkHash []byte, netID byte) (*AddressPubKeyHash, error) {
	if len(pkHash) != ripemd160.Size {
		return nil, errors.New("pkHash must be 20 bytes")
	}
	addr := &AddressPubKeyHash{
		netID: netID,
	}
	copy(addr.hash[:], pkHash)
	return addr, nil
}

// EncodeAddress returns the string encoding of a pay-to-pubkey-hash
// address.  Part of the Address interface.
func (a *AddressPubKeyHash) EncodeAddress() string {
	return encodeAddress(a.hash[:], a.netID)
}

// Hash160 return hash160
func (a *AddressPubKeyHash) Hash160() []byte {
	return a.hash[:]
}
