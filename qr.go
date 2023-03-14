package gotp

import "fmt"

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
	url += "&algorithm=" + Config.AlgorithmName + "&digits=" + Config.Digits + "&period=" + fmt.Sprint(Config.Period)
	return "https://chart.googleapis.com/chart?chs=200x200&chld=M|0&cht=qr&chl=" + url
}
