package hdkeychain

import (
	"errors"
	"fmt"
	"math/big"
	"testing"

	"github.com/duminghui/go-bip32/d/chaincfg"
)

type testVector struct {
	seed       string
	chainSlice []chain
}

type chain struct {
	path string
	pub  string
	prv  string
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

func testVectorMethod(testVector *testVector) {
	seed, _ := BytesFromHexStrFixZeroPrefix(testVector.seed)
	fmt.Printf("seed: %x\n", seed)
	mstKey, err := NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		panic(err)
	}
	for idx, c := range testVector.chainSlice {
		if idx == 0 {
			printVectorKey(mstKey, c.path, c.pub, c.prv)
		} else {
			childKey, err := mstKey.DerivePath(c.path)
			if err != nil {
				panic(err)
			}
			printVectorKey(childKey, c.path, c.pub, c.prv)
		}
	}
}

func printVectorKey(key *ExtendedKey, path, pub, prv string) {
	fmt.Println("------------------------------")
	fmt.Printf("%-*s%s\n", 17, "Derivation Path:", path)
	prvStr := key.String()
	fmt.Printf("%-*s%s [%v]\n", 17, "Extended Key:", prvStr, prvStr == prv)
	pubKey, err := key.Neuter()
	if err != nil {
		fmt.Println("key.Neuter:", err)
	}
	pubKeyStr := pubKey.String()
	fmt.Printf("%-*s%s [%v]\n", 17, "Extended PubKey:", pubKeyStr, pubKeyStr == pub)
	address, _ := key.Address(&chaincfg.MainNetParams)
	fmt.Printf("%-*s%s(%x)\n", 17, "Address:", address.EncodeAddress(), address.Hash160())
}

var testVector1 = &testVector{
	seed: "000102030405060708090a0b0c0d0e0f",
	chainSlice: []chain{
		{
			path: "m",
			pub:  "xpub661MyMwAqRbcFtXgS5sYJABqqG9YLmC4Q1Rdap9gSE8NqtwybGhePY2gZ29ESFjqJoCu1Rupje8YtGqsefD265TMg7usUDFdp6W1EGMcet8",
			prv:  "xprv9s21ZrQH143K3QTDL4LXw2F7HEK3wJUD2nW2nRk4stbPy6cq3jPPqjiChkVvvNKmPGJxWUtg6LnF5kejMRNNU3TGtRBeJgk33yuGBxrMPHi",
		},
		{
			path: "m/0'",
			pub:  "xpub68Gmy5EdvgibQVfPdqkBBCHxA5htiqg55crXYuXoQRKfDBFA1WEjWgP6LHhwBZeNK1VTsfTFUHCdrfp1bgwQ9xv5ski8PX9rL2dZXvgGDnw",
			prv:  "xprv9uHRZZhk6KAJC1avXpDAp4MDc3sQKNxDiPvvkX8Br5ngLNv1TxvUxt4cV1rGL5hj6KCesnDYUhd7oWgT11eZG7XnxHrnYeSvkzY7d2bhkJ7",
		},
		{
			path: "m/0'/1",
			pub:  "xpub6ASuArnXKPbfEwhqN6e3mwBcDTgzisQN1wXN9BJcM47sSikHjJf3UFHKkNAWbWMiGj7Wf5uMash7SyYq527Hqck2AxYysAA7xmALppuCkwQ",
			prv:  "xprv9wTYmMFdV23N2TdNG573QoEsfRrWKQgWeibmLntzniatZvR9BmLnvSxqu53Kw1UmYPxLgboyZQaXwTCg8MSY3H2EU4pWcQDnRnrVA1xe8fs",
		},
		{
			path: "m/0'/1/2'",
			pub:  "xpub6D4BDPcP2GT577Vvch3R8wDkScZWzQzMMUm3PWbmWvVJrZwQY4VUNgqFJPMM3No2dFDFGTsxxpG5uJh7n7epu4trkrX7x7DogT5Uv6fcLW5",
			prv:  "xprv9z4pot5VBttmtdRTWfWQmoH1taj2axGVzFqSb8C9xaxKymcFzXBDptWmT7FwuEzG3ryjH4ktypQSAewRiNMjANTtpgP4mLTj34bhnZX7UiM",
		},
		{
			path: "m/0'/1/2'/2",
			pub:  "xpub6FHa3pjLCk84BayeJxFW2SP4XRrFd1JYnxeLeU8EqN3vDfZmbqBqaGJAyiLjTAwm6ZLRQUMv1ZACTj37sR62cfN7fe5JnJ7dh8zL4fiyLHV",
			prv:  "xprvA2JDeKCSNNZky6uBCviVfJSKyQ1mDYahRjijr5idH2WwLsEd4Hsb2Tyh8RfQMuPh7f7RtyzTtdrbdqqsunu5Mm3wDvUAKRHSC34sJ7in334",
		},
		{
			path: "m/0'/1/2'/2/1000000000",
			pub:  "xpub6H1LXWLaKsWFhvm6RVpEL9P4KfRZSW7abD2ttkWP3SSQvnyA8FSVqNTEcYFgJS2UaFcxupHiYkro49S8yGasTvXEYBVPamhGW6cFJodrTHy",
			prv:  "xprvA41z7zogVVwxVSgdKUHDy1SKmdb533PjDz7J6N6mV6uS3ze1ai8FHa8kmHScGpWmj4WggLyQjgPie1rFSruoUihUZREPSL39UNdE3BBDu76",
		},
	},
}

// https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#test-vector-1
func TestVector1(t *testing.T) {
	fmt.Println("====================TestVector1=======================")
	testVectorMethod(testVector1)
}

//https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#test-vector-2
var testVector2 = &testVector{
	seed: "fffcf9f6f3f0edeae7e4e1dedbd8d5d2cfccc9c6c3c0bdbab7b4b1aeaba8a5a29f9c999693908d8a8784817e7b7875726f6c696663605d5a5754514e4b484542",
	chainSlice: []chain{
		{
			path: "m",
			pub:  "xpub661MyMwAqRbcFW31YEwpkMuc5THy2PSt5bDMsktWQcFF8syAmRUapSCGu8ED9W6oDMSgv6Zz8idoc4a6mr8BDzTJY47LJhkJ8UB7WEGuduB",
			prv:  "xprv9s21ZrQH143K31xYSDQpPDxsXRTUcvj2iNHm5NUtrGiGG5e2DtALGdso3pGz6ssrdK4PFmM8NSpSBHNqPqm55Qn3LqFtT2emdEXVYsCzC2U",
		},
		{
			path: "m/0",
			pub:  "xpub69H7F5d8KSRgmmdJg2KhpAK8SR3DjMwAdkxj3ZuxV27CprR9LgpeyGmXUbC6wb7ERfvrnKZjXoUmmDznezpbZb7ap6r1D3tgFxHmwMkQTPH",
			prv:  "xprv9vHkqa6EV4sPZHYqZznhT2NPtPCjKuDKGY38FBWLvgaDx45zo9WQRUT3dKYnjwih2yJD9mkrocEZXo1ex8G81dwSM1fwqWpWkeS3v86pgKt",
		},
		{
			path: "m/0/2147483647'",
			pub:  "xpub6ASAVgeehLbnwdqV6UKMHVzgqAG8Gr6riv3Fxxpj8ksbH9ebxaEyBLZ85ySDhKiLDBrQSARLq1uNRts8RuJiHjaDMBU4Zn9h8LZNnBC5y4a",
			prv:  "xprv9wSp6B7kry3Vj9m1zSnLvN3xH8RdsPP1Mh7fAaR7aRLcQMKTR2vidYEeEg2mUCTAwCd6vnxVrcjfy2kRgVsFawNzmjuHc2YmYRmagcEPdU9",
		},
		{
			path: "m/0/2147483647'/1",
			pub:  "xpub6DF8uhdarytz3FWdA8TvFSvvAh8dP3283MY7p2V4SeE2wyWmG5mg5EwVvmdMVCQcoNJxGoWaU9DCWh89LojfZ537wTfunKau47EL2dhHKon",
			prv:  "xprv9zFnWC6h2cLgpmSA46vutJzBcfJ8yaJGg8cX1e5StJh45BBciYTRXSd25UEPVuesF9yog62tGAQtHjXajPPdbRCHuWS6T8XA2ECKADdw4Ef",
		},
		{
			path: "m/0/2147483647'/1/2147483646'",
			pub:  "xpub6ERApfZwUNrhLCkDtcHTcxd75RbzS1ed54G1LkBUHQVHQKqhMkhgbmJbZRkrgZw4koxb5JaHWkY4ALHY2grBGRjaDMzQLcgJvLJuZZvRcEL",
			prv:  "xprvA1RpRA33e1JQ7ifknakTFpgNXPmW2YvmhqLQYMmrj4xJXXWYpDPS3xz7iAxn8L39njGVyuoseXzU6rcxFLJ8HFsTjSyQbLYnMpCqE2VbFWc",
		},
		{
			path: "m/0/2147483647'/1/2147483646'/2",
			pub:  "xpub6FnCn6nSzZAw5Tw7cgR9bi15UV96gLZhjDstkXXxvCLsUXBGXPdSnLFbdpq8p9HmGsApME5hQTZ3emM2rnY5agb9rXpVGyy3bdW6EEgAtqt",
			prv:  "xprvA2nrNbFZABcdryreWet9Ea4LvTJcGsqrMzxHx98MMrotbir7yrKCEXw7nadnHM8Dq38EGfSh6dqA9QWTyefMLEcBYJUuekgW4BYPJcr9E7j",
		},
	},
}

//
func TestVector2(t *testing.T) {
	fmt.Println("=========================Vector2========================")
	testVectorMethod(testVector2)
}

var testVector3 = &testVector{
	seed: "4b381541583be4423346c643850da4b320e46a87ae3d2a4e6da11eba819cd4acba45d239319ac14f863b8d5ab5a0d0c64d2e8a1e7d1457df2e5a3c51c73235be",
	chainSlice: []chain{
		{
			path: "m",
			pub:  "xpub661MyMwAqRbcEZVB4dScxMAdx6d4nFc9nvyvH3v4gJL378CSRZiYmhRoP7mBy6gSPSCYk6SzXPTf3ND1cZAceL7SfJ1Z3GC8vBgp2epUt13",
			prv:  "xprv9s21ZrQH143K25QhxbucbDDuQ4naNntJRi4KUfWT7xo4EKsHt2QJDu7KXp1A3u7Bi1j8ph3EGsZ9Xvz9dGuVrtHHs7pXeTzjuxBrCmmhgC6",
		},
		{
			path: "m/0'",
			pub:  "xpub68NZiKmJWnxxS6aaHmn81bvJeTESw724CRDs6HbuccFQN9Ku14VQrADWgqbhhTHBaohPX4CjNLf9fq9MYo6oDaPPLPxSb7gwQN3ih19Zm4Y",
			prv:  "xprv9uPDJpEQgRQfDcW7BkF7eTya6RPxXeJCqCJGHuCJ4GiRVLzkTXBAJMu2qaMWPrS7AANYqdq6vcBcBUdJCVVFceUvJFjaPdGZ2y9WACViL4L",
		},
	},
}

func TestVector3(t *testing.T) {
	fmt.Println("=========================Vector3========================")
	testVectorMethod(testVector3)
}
