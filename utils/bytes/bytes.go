// Package bytes provides ...
package bytes

// PaddedAppend append src to dst, if less than size padding 0 at start
func PaddedAppend(size int, dst, src []byte) []byte {
	return append(dst, PaddedBytes(size, src)...)
}

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
