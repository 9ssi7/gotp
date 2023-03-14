package gotp

import (
	"fmt"
	"math/rand"
)

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

type QrConfig struct {
	Secret      string
	Issuer      string
	AccountName string
}

func GetQRCode(config QrConfig) string {
	url := "otpauth://totp/" + config.AccountName + "?secret=" + config.Secret
	if config.Issuer != "" {
		url += "&issuer=" + config.Issuer
	}
	url += "&algorithm=" + Algorithm + "&digits=" + Digits + "&period=" + fmt.Sprint(Period)
	return "https://chart.googleapis.com/chart?chs=200x200&chld=M|0&cht=qr&chl=" + url
}
