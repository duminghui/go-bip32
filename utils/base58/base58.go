package base58

import (
	"github.com/duminghui/go-bip32/utils/basen"
)

var base58Encoding = basen.NewEncoding("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

//Encode to base58
func Encode(data []byte) string {
	return base58Encoding.EncodeToString(data)
}
