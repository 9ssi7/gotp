package gotp

import (
	"crypto/hmac"
	"encoding/base32"
	"fmt"
)

type config struct {
	Digits     Digits
	Algorithm  Algorithm
	SecretSize int
	Period     int64
	ValidChars string
}

func getCode(secret string, time int64) (string, error) {
	secretBytes, err := base32Decode(secret)
	if err != nil {
		return "", err
	}
	timeBytes := make([]byte, 8)
	for i := 0; i < 8; i++ {
		timeBytes[i] = byte(time >> uint(8*(7-i)) & 0xff)
	}
	hash := createHash(secretBytes, timeBytes)
	offset := hash[len(hash)-1] & 0xf
	truncatedHash := hash[offset : offset+4]
	truncatedHash[0] = truncatedHash[0] & 0x7f
	code := int(truncatedHash[0])<<24 | int(truncatedHash[1])<<16 | int(truncatedHash[2])<<8 | int(truncatedHash[3])
	code = code % int(1e6)
	return fmt.Sprintf("%06d", code), nil
}

func base32Decode(secret string) ([]byte, error) {
	secret = secret[:len(secret)-len(secret)%8]
	return base32.StdEncoding.DecodeString(secret)
}

func createHash(key, message []byte) []byte {
	mac := hmac.New(Config.Algorithm.Hash(), key)
	mac.Write(message)
	return mac.Sum(nil)
}
