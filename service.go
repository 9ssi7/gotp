package gotp

import (
	"crypto/sha1"
	"math/rand"
	"time"
)

var Config = config{
	CodeLength:    6,
	SecretSize:    16,
	Period:        30,
	Algorithm:     sha1.New,
	AlgorithmName: "SHA1",
	Digits:        "DECIMAL",
	ValidChars:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567=",
}

func CreateSecret() string {
	validChars := []byte(Config.ValidChars)
	secret := make([]byte, Config.SecretSize)
	for i := 0; i < Config.SecretSize; i++ {
		secret[i] = validChars[rand.Intn(len(validChars))]
	}
	return string(secret)
}

func GetCode(secret string, time int64) string {
	code, _ := getCode(secret, time)
	return code
}

func GetTime() int64 {
	return time.Now().Unix() / Config.Period
}
