package password

import (
	"bytes"
	"crypto/sha512"
	"encoding/hex"
	"errors"
)

func ParamToByte(hexStr string) ([]byte, error) {
	decoded, err := hex.DecodeString(hexStr)
	if err != nil {
		return []byte{}, err
	}

	if len(decoded) != 64 {
		return []byte{}, errors.New("length of hash is invalid")
	}

	return decoded, nil
}

func StringToByte(str string) []byte {
	hash := sha512.Sum512([]byte(str))
	return hash[:]
}

func IsPasswordValid(src []byte, dst []byte) (bool, error) {
	if len(src) != 64 || len(dst) != 64 {
		return false, errors.New("length of hash is invalid")
	}
	return bytes.Equal(src, dst), nil
}
