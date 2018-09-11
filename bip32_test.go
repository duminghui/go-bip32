package bip32

import (
	"fmt"
	"testing"

	"github.com/duminghui/go-bip32/utils/bytes"
	"github.com/duminghui/go-bip32/wif"
	"github.com/duminghui/go-bip39"
)

func TestNewMasterKey(t *testing.T) {
	t.SkipNow()
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		panic(err)
	}
	fmt.Printf("entropy:%x\n", entropy)
	mnemonic, err := bip39.NewMnemonic(entropy)
	mnemonic = "lemon speak wish predict athlete fringe ritual bonus ensure clay school goat"
	fmt.Printf("mnemonic:%s\n", mnemonic)
	seed, err := bip39.NewSeedWithValidMnemonic(mnemonic, "")
	fmt.Printf("seed:%x, len:%d\n", seed, len(seed))
	// seedNum, _ := new(big.Int).SetString("000102030405060708090a0b0c0d0e0f", 16)
	// seed = seedNum.Bytes()
	// seed = bytes.PaddedBytes(16, seed)
	// fmt.Printf("seed2: %x\n", seed)
	masterKey, err := NewMasterKey(seed, &BTCMainNetParams)
	fmt.Println(err, len(seed))
	// fmt.Printf("MK version:%x\n", masterKey.version)
	// fmt.Printf("MK childNum:%x\n", masterKey.childNum)
	// fmt.Printf("MK key:%x\n", masterKey.key)
	// fmt.Printf("MK chainCode:%x\n", masterKey.chainCode)
	// fmt.Printf("MK fingerprint:%x\n", masterKey.parentFP)
	// fmt.Printf("MK depth:%x\n", masterKey.depth)
	// fmt.Printf("MK isPrivate:%v\n", masterKey.isPrivate)
	b58 := masterKey.B58Serialize()
	fmt.Println("MK b58:", b58)
	b58Vaild := "xprv9s21ZrQH143K3dRqV7LTj7uFvjMFiai3HV7Ytoj7Jg3mEArCCk2ujozgMSoip9FTWH4YwjTnjayPooxVesrtHst8qo8xZCnG1MazdLfBJBR"
	fmt.Println(b58 == b58Vaild)
	masterPubKey := masterKey.Neuter()
	b58pubKey := masterPubKey.B58Serialize()
	fmt.Println("masterPubKey:", b58pubKey)
	b58pubValid := "xpub661MyMwAqRbcG7WJb8sU6FqzUmBk83Rtei39hC8is1ak6yBLkHMAHcKACgmyxr6KLGhLSM4Qr74R8meHNcttv4bmPMDC8EwB9vJ3HRGaC53"
	fmt.Println(b58pubKey == b58pubValid)
	fmt.Println(len(masterKey.key))
	fmt.Println(len("1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd"))
	// bigNum := new(big.Int).SetBytes(masterKey.key)
	fmt.Println((fmt.Sprintf("%x", masterKey.key)))
	fmt.Println((fmt.Sprintf("%x", masterKey.pubKeyBytes())))
}

type testVector struct {
	seed  string
	chain map[string]chain
}

type chain struct {
	pub string
	prv string
}

func testVectorMethod(testVector *testVector) {
	seed, _ := bytes.BytesFromHexStrFixZeroPrefix(testVector.seed)
	fmt.Printf("seed: %x\n", seed)
	masterKey, err := NewMasterKey(seed, &BTCMainNetParams)
	if err != nil {
		panic(err)
	}
	printVectorKey(masterKey, testVector)
	for path := range testVector.chain {
		childKey, err := masterKey.DerivePath(path)
		if err != nil {
			panic(err)
		}
		printVectorKey(childKey, testVector)
	}
}

var testVector1 = &testVector{
	seed: "000102030405060708090a0b0c0d0e0f",
	chain: map[string]chain{
		"m": chain{
			pub: "xpub661MyMwAqRbcFtXgS5sYJABqqG9YLmC4Q1Rdap9gSE8NqtwybGhePY2gZ29ESFjqJoCu1Rupje8YtGqsefD265TMg7usUDFdp6W1EGMcet8",
			prv: "xprv9s21ZrQH143K3QTDL4LXw2F7HEK3wJUD2nW2nRk4stbPy6cq3jPPqjiChkVvvNKmPGJxWUtg6LnF5kejMRNNU3TGtRBeJgk33yuGBxrMPHi",
		},
		"m/0'": chain{
			pub: "xpub68Gmy5EdvgibQVfPdqkBBCHxA5htiqg55crXYuXoQRKfDBFA1WEjWgP6LHhwBZeNK1VTsfTFUHCdrfp1bgwQ9xv5ski8PX9rL2dZXvgGDnw",
			prv: "xprv9uHRZZhk6KAJC1avXpDAp4MDc3sQKNxDiPvvkX8Br5ngLNv1TxvUxt4cV1rGL5hj6KCesnDYUhd7oWgT11eZG7XnxHrnYeSvkzY7d2bhkJ7",
		},
		"m/0'/1": chain{
			pub: "xpub6ASuArnXKPbfEwhqN6e3mwBcDTgzisQN1wXN9BJcM47sSikHjJf3UFHKkNAWbWMiGj7Wf5uMash7SyYq527Hqck2AxYysAA7xmALppuCkwQ",
			prv: "xprv9wTYmMFdV23N2TdNG573QoEsfRrWKQgWeibmLntzniatZvR9BmLnvSxqu53Kw1UmYPxLgboyZQaXwTCg8MSY3H2EU4pWcQDnRnrVA1xe8fs",
		},
		"m/0'/1/2'": chain{
			pub: "xpub6D4BDPcP2GT577Vvch3R8wDkScZWzQzMMUm3PWbmWvVJrZwQY4VUNgqFJPMM3No2dFDFGTsxxpG5uJh7n7epu4trkrX7x7DogT5Uv6fcLW5",
			prv: "xprv9z4pot5VBttmtdRTWfWQmoH1taj2axGVzFqSb8C9xaxKymcFzXBDptWmT7FwuEzG3ryjH4ktypQSAewRiNMjANTtpgP4mLTj34bhnZX7UiM",
		},
		"m/0'/1/2'/2": chain{
			pub: "xpub6FHa3pjLCk84BayeJxFW2SP4XRrFd1JYnxeLeU8EqN3vDfZmbqBqaGJAyiLjTAwm6ZLRQUMv1ZACTj37sR62cfN7fe5JnJ7dh8zL4fiyLHV",
			prv: "xprvA2JDeKCSNNZky6uBCviVfJSKyQ1mDYahRjijr5idH2WwLsEd4Hsb2Tyh8RfQMuPh7f7RtyzTtdrbdqqsunu5Mm3wDvUAKRHSC34sJ7in334",
		},
		"m/0'/1/2'/2/1000000000": chain{
			pub: "xpub6H1LXWLaKsWFhvm6RVpEL9P4KfRZSW7abD2ttkWP3SSQvnyA8FSVqNTEcYFgJS2UaFcxupHiYkro49S8yGasTvXEYBVPamhGW6cFJodrTHy",
			prv: "xprvA41z7zogVVwxVSgdKUHDy1SKmdb533PjDz7J6N6mV6uS3ze1ai8FHa8kmHScGpWmj4WggLyQjgPie1rFSruoUihUZREPSL39UNdE3BBDu76",
		},
	},
}

// https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#test-vector-1
func TestVector1(t *testing.T) {
	t.SkipNow()
	fmt.Println("====================TestVector1=======================")
	testVectorMethod(testVector1)
}

func printVectorKey(key *ExtendedKey, testVector *testVector) {
	vaildValue := testVector.chain[key.DerivationPath]
	fmt.Println("------------------------------")
	fmt.Printf("%-*s%s\n", 17, "Derivation Path:", key.DerivationPath)
	b58Key := key.B58Serialize()
	fmt.Printf("%-*s%s [%v]\n", 17, "Extended Key:", b58Key, b58Key == vaildValue.prv)
	pubKey := key.Neuter()
	b58PubKey := pubKey.B58Serialize()
	fmt.Printf("%-*s%s [%v]\n", 17, "Extended PubKey:", b58PubKey, b58PubKey == vaildValue.pub)
	address, _ := key.Address()
	fmt.Printf("%-*s%s(%x)\n", 17, "Address:", address.EncodeAddress(), address.Hash160())
}

//https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#test-vector-2
var testVector2 = &testVector{
	seed: "fffcf9f6f3f0edeae7e4e1dedbd8d5d2cfccc9c6c3c0bdbab7b4b1aeaba8a5a29f9c999693908d8a8784817e7b7875726f6c696663605d5a5754514e4b484542",
	chain: map[string]chain{
		"m": chain{
			pub: "xpub661MyMwAqRbcFW31YEwpkMuc5THy2PSt5bDMsktWQcFF8syAmRUapSCGu8ED9W6oDMSgv6Zz8idoc4a6mr8BDzTJY47LJhkJ8UB7WEGuduB",
			prv: "xprv9s21ZrQH143K31xYSDQpPDxsXRTUcvj2iNHm5NUtrGiGG5e2DtALGdso3pGz6ssrdK4PFmM8NSpSBHNqPqm55Qn3LqFtT2emdEXVYsCzC2U",
		},
		"m/0": chain{
			pub: "xpub69H7F5d8KSRgmmdJg2KhpAK8SR3DjMwAdkxj3ZuxV27CprR9LgpeyGmXUbC6wb7ERfvrnKZjXoUmmDznezpbZb7ap6r1D3tgFxHmwMkQTPH",
			prv: "xprv9vHkqa6EV4sPZHYqZznhT2NPtPCjKuDKGY38FBWLvgaDx45zo9WQRUT3dKYnjwih2yJD9mkrocEZXo1ex8G81dwSM1fwqWpWkeS3v86pgKt",
		},
		"m/0/2147483647'": chain{
			pub: "xpub6ASAVgeehLbnwdqV6UKMHVzgqAG8Gr6riv3Fxxpj8ksbH9ebxaEyBLZ85ySDhKiLDBrQSARLq1uNRts8RuJiHjaDMBU4Zn9h8LZNnBC5y4a",
			prv: "xprv9wSp6B7kry3Vj9m1zSnLvN3xH8RdsPP1Mh7fAaR7aRLcQMKTR2vidYEeEg2mUCTAwCd6vnxVrcjfy2kRgVsFawNzmjuHc2YmYRmagcEPdU9",
		},
		"m/0/2147483647'/1": chain{
			pub: "xpub6DF8uhdarytz3FWdA8TvFSvvAh8dP3283MY7p2V4SeE2wyWmG5mg5EwVvmdMVCQcoNJxGoWaU9DCWh89LojfZ537wTfunKau47EL2dhHKon",
			prv: "xprv9zFnWC6h2cLgpmSA46vutJzBcfJ8yaJGg8cX1e5StJh45BBciYTRXSd25UEPVuesF9yog62tGAQtHjXajPPdbRCHuWS6T8XA2ECKADdw4Ef",
		},
		"m/0/2147483647'/1/2147483646'": chain{
			pub: "xpub6ERApfZwUNrhLCkDtcHTcxd75RbzS1ed54G1LkBUHQVHQKqhMkhgbmJbZRkrgZw4koxb5JaHWkY4ALHY2grBGRjaDMzQLcgJvLJuZZvRcEL",
			prv: "xprvA1RpRA33e1JQ7ifknakTFpgNXPmW2YvmhqLQYMmrj4xJXXWYpDPS3xz7iAxn8L39njGVyuoseXzU6rcxFLJ8HFsTjSyQbLYnMpCqE2VbFWc",
		},
		"m/0/2147483647'/1/2147483646'/2": chain{
			pub: "xpub6FnCn6nSzZAw5Tw7cgR9bi15UV96gLZhjDstkXXxvCLsUXBGXPdSnLFbdpq8p9HmGsApME5hQTZ3emM2rnY5agb9rXpVGyy3bdW6EEgAtqt",
			prv: "xprvA2nrNbFZABcdryreWet9Ea4LvTJcGsqrMzxHx98MMrotbir7yrKCEXw7nadnHM8Dq38EGfSh6dqA9QWTyefMLEcBYJUuekgW4BYPJcr9E7j",
		},
	},
}

func TestVector2(t *testing.T) {
	t.SkipNow()
	fmt.Println("=========================Vector2========================")
	testVectorMethod(testVector2)
}

var testVector3 = &testVector{
	seed: "4b381541583be4423346c643850da4b320e46a87ae3d2a4e6da11eba819cd4acba45d239319ac14f863b8d5ab5a0d0c64d2e8a1e7d1457df2e5a3c51c73235be",
	chain: map[string]chain{
		"m": {
			pub: "xpub661MyMwAqRbcEZVB4dScxMAdx6d4nFc9nvyvH3v4gJL378CSRZiYmhRoP7mBy6gSPSCYk6SzXPTf3ND1cZAceL7SfJ1Z3GC8vBgp2epUt13",
			prv: "xprv9s21ZrQH143K25QhxbucbDDuQ4naNntJRi4KUfWT7xo4EKsHt2QJDu7KXp1A3u7Bi1j8ph3EGsZ9Xvz9dGuVrtHHs7pXeTzjuxBrCmmhgC6",
		},
		"m/0'": {
			pub: "xpub68NZiKmJWnxxS6aaHmn81bvJeTESw724CRDs6HbuccFQN9Ku14VQrADWgqbhhTHBaohPX4CjNLf9fq9MYo6oDaPPLPxSb7gwQN3ih19Zm4Y",
			prv: "xprv9uPDJpEQgRQfDcW7BkF7eTya6RPxXeJCqCJGHuCJ4GiRVLzkTXBAJMu2qaMWPrS7AANYqdq6vcBcBUdJCVVFceUvJFjaPdGZ2y9WACViL4L",
		},
	},
}

func TestVector3(t *testing.T) {
	t.SkipNow()
	fmt.Println("=========================Vector3========================")
	testVectorMethod(testVector3)
}

// https://github.com/iancoleman/bip39/issues/58
// 17rxURoF96VhmkcEGCj5LNQkmN9HVhWb7F
func TestVector3_2(t *testing.T) {
	t.SkipNow()
	mnemnic := "fruit wave dwarf banana earth journey tattoo true farm silk olive fence"
	seed, _ := bip39.NewSeedWithValidMnemonic(mnemnic, "banana")
	key, _ := NewMasterKey(seed, &BTCMainNetParams)
	fmt.Println(key.B58Serialize())
	childKey, _ := key.DerivePath("m/44'/0'/0'/0/0")
	address, _ := childKey.Address()
	fmt.Println(address.EncodeAddress())
	fmt.Printf("%x\n", childKey.pubKeyBytes())
}

func TestNeute(t *testing.T) {
	// t.SkipNow()
	mnemnic := "fruit wave dwarf banana earth journey tattoo true farm silk olive fence"
	seed, _ := bip39.NewSeedWithValidMnemonic(mnemnic, "banana")
	masterKey, _ := NewMasterKey(seed, &BTCMainNetParams)
	masterPubKey := masterKey.Neuter()
	fmt.Println("masterKey:", masterKey.B58Serialize())
	// Normal N(CKDpriv((k,c),i))
	childKey1, _ := masterKey.Child(0)
	fmt.Println("childKey1:   ", childKey1.B58Serialize())
	fmt.Println("childPubKey1:", childKey1.Neuter().B58Serialize())
	// Noremal CKDpub(N(k,c),i)
	childKey1, _ = masterPubKey.Child(0)
	fmt.Println("childKey1:   ", childKey1.B58Serialize())
	fmt.Println("childPubKey1:", childKey1.Neuter().B58Serialize())

	// hardened N(CKDpriv((k,c),i))
	childKey2, _ := masterKey.HardenedChild(0)
	fmt.Println("childKey2:   ", childKey2.B58Serialize())
	fmt.Println("childPubKey2:", childKey2.Neuter().B58Serialize())
	// hardened CKDpub(N(k,c),i)
	_, err := masterPubKey.HardenedChild(0)
	fmt.Println("childPubKey2:", err)
	fmt.Println("-----------------------------------")
	childKey2ChildKey, _ := childKey2.Child(0)
	fmt.Println(childKey2ChildKey.DerivationPath)
	fmt.Println("CCKey", childKey2ChildKey.B58Serialize())
	privKey, _ := childKey2ChildKey.ECPrivKey()
	fmt.Println("PubkKey", privKey.PubKey.SerializeCompressed())
	fmt.Printf("PubKey:%x\n", privKey.PubKey.SerializeCompressed())
	wif := wif.NewWIF(privKey, BTCMainNetParams.PrivateKeyID)
	fmt.Println("WIF", wif.EncodeCompressed())
	address, _ := childKey2ChildKey.Address()
	fmt.Println("Address", address.EncodeAddress())
}
