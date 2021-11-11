package bip32

import (
	"testing"
)

func TestNewAddressWitnessPubKeyHash(t *testing.T) {
	//mnemonic := "expand photo unfold meat drive summer wool auto good crystal deposit quick"
	//seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	//if err != nil {
	//	fmt.Println(1, err)
	//	return
	//}
	//masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	//fmt.Printf("%x\n", masterKey.Version())
	//if err != nil {
	//	fmt.Println(2, err)
	//	return
	//}
	//childKey, err := masterKey.DerivePath("m/84'/0'/0'/0/0")
	//if err != nil {
	//	fmt.Println(3, err)
	//	return
	//}
	//pubKey, err := childKey.ECPubKey()
	//if err != nil {
	//	fmt.Println(4, err)
	//	return
	//}
	//
	//pubKeyHash := util.Hash160(pubKey.SerializeCompressed())
	//fmt.Printf("%x\n", pubKeyHash)
	//
	//p2wkhAddr, err := util.NewAddressWitnessPubKeyHash(pubKeyHash, &chaincfg.MainNetParams)
	//if err != nil {
	//	fmt.Println(5, err)
	//	return
	//}
	//fmt.Println(p2wkhAddr.EncodeAddress())
	//
	//childKey2, err := masterKey.DerivePath("m/49'/0'/0'/0/0")
	//if err != nil {
	//	fmt.Println(6, err)
	//	return
	//}
	//pubKey2, err := childKey2.ECPubKey()
	//if err != nil {
	//	fmt.Println(7, err)
	//	return
	//}
	//
	//witnessAddr, err := util.NewAddressWitnessPubKeyHash(util.Hash160(pubKey2.SerializeCompressed()), &chaincfg.MainNetParams)
	//data := witnessAddr.ScriptAddress()
	//dataLen := len(data)
	//fmt.Println(data)
	//fmt.Println(len(data))
	//witnessProg := []byte{0x00, byte(dataLen)}
	//witnessProg = append(witnessProg, data...)
	//
	//addr, err := util.NewAddressScriptHash(witnessProg, &chaincfg.MainNetParams)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(addr.EncodeAddress())

}
