// Package bytes provides ...
package bytes

import (
	"errors"
	"math/big"
)

// PaddedBytes padding byte array to size length
func PaddedBytes(size int, src []byte) []byte {
	offset := size - len(src)
	tmp := src
	if offset > 0 {
		tmp = make([]byte, size)
		copy(tmp[offset:], src)
	}
	return tmp
}

// BytesFromHexStrFixZeroPrefix return fix Zero start strings
// like 00010203040506
func BytesFromHexStrFixZeroPrefix(str string) ([]byte, error) {
	strNum, ok := new(big.Int).SetString(str, 16)
	if !ok {
		return nil, errors.New("string error")
	}
	bytes := strNum.Bytes()
	bytes = PaddedBytes(len(str)/2, bytes)
	return bytes, nil
}
