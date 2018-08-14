package basen

var base58Encoding = NewEncoding("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

//Base58Encode to base58
func Base58Encode(data []byte) string {
	return base58Encoding.EncodeToString(data)
}
