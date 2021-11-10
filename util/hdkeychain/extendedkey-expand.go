package hdkeychain

import (
	"errors"
	"strconv"
	"strings"
)

var ErrKeyPathFormat = errors.New("wallet path error")
var ErrParentKey = errors.New("key must master")

var mapKey = make(map[string]*ExtendedKey)

// DerivePath return key by path : m/0'/1/2' etc...
func (k *ExtendedKey) DerivePath(pathStr string) (*ExtendedKey, error) {
	//fmt.Println("###########", len(mapKey))
	if k.depth > 0 || k.childNum > 0 {
		return nil, ErrParentKey
	}
	keyTmp, ok := mapKey[pathStr]
	if ok {
		return keyTmp, nil
	}
	path := strings.Split(pathStr, "/")
	err := validPath(path)
	if err != nil {
		return nil, err
	}
	var tmpPath []string
	var tmpPathStr string
	var tmpParentKey *ExtendedKey
	for _, childNumStr := range path {
		tmpPath = append(tmpPath, childNumStr)
		tmpPathStr = strings.Join(tmpPath, "/")
		keyTmp, ok = mapKey[tmpPathStr]
		if !ok {
			if tmpPathStr == "m" {
				keyTmp = k
			} else {
				isHardenedChild := false
				if strings.HasSuffix(childNumStr, "'") {
					childNumStr = strings.Replace(childNumStr, "'", "", -1)
					isHardenedChild = true
				}
				childNum, _ := strconv.Atoi(childNumStr)
				var err error
				if isHardenedChild {
					childNum = HardenedKeyStart + childNum
					keyTmp, err = tmpParentKey.Derive(uint32(childNum))
				} else {
					keyTmp, err = tmpParentKey.Derive(uint32(childNum))
				}
				if err != nil {
					return nil, err
				}
			}
			mapKey[tmpPathStr] = keyTmp
		}
		tmpParentKey = keyTmp
	}
	return keyTmp, nil
}

func validPath(path []string) error {
	if path[0] != "m" {
		return ErrKeyPathFormat
	}
	for i := 1; i < len(path); i++ {
		childNumStr := path[i]
		if strings.HasSuffix(childNumStr, "'") {
			childNumStr = strings.Replace(childNumStr, "'", "", -1)
		}
		childNum, err := strconv.Atoi(childNumStr)
		if err != nil {
			return ErrKeyPathFormat
		}
		if childNum >= HardenedKeyStart || childNum < 0 {
			return ErrKeyPathFormat
		}
	}
	return nil
}
