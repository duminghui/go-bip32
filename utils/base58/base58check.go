// Package basen provides ...
package base58

import "crypto/sha256"

func checksum(input []byte) (cksum [4]byte) {
	h := sha256.Sum256(input)
	h2 := sha256.Sum256(h[:])
	copy(cksum[:], h2[:4])
	return
}

func CheckEncode(input []byte) string {
	b := make([]byte, len(input)+4)
	cksum := checksum(input)
	b = append(b, cksum[:]...)
	return Encode(b)

}
