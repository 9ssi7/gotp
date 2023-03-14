package gotp

import "math/rand"

var (
	CodeLength = 6
	SecretSize = 16
	Period     = 30
	Algorithm  = "SHA1"
	Digits     = "DECIMAL"
	ValidChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567="
)

func CreateSecret() string {
	validChars := []byte(ValidChars)
	secret := make([]byte, SecretSize)
	for i := 0; i < SecretSize; i++ {
		secret[i] = validChars[rand.Intn(len(validChars))]
	}
	return string(secret)
}

func GetCode() {}

func Verify() {}

func GetQRCode() {}
