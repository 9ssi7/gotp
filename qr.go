package gotp

import (
	"fmt"
	"net/url"
)

// QrConfig is the configuration for the GetQRCode function.
type QrConfig struct {
	// Secret is the secret key.
	Secret string

	// Issuer is the name of the issuer. This is optional. If set, it will be displayed on the authenticator app. For example, "Google".
	Issuer string

	// AccountName is the name of the account. This is required. For example, "John Doe".
	// This will be displayed on the authenticator app.
	// The final display will be "Issuer (AccountName)" or "AccountName" if Issuer is not set.
	AccountName string
}

// GetQRCode returns the URL of the QR code for the secret key.
// The QR code can be used to add the secret key to an authenticator app.
// The QR code is a Google Chart.
// The QR code is a PNG image.
func GetQRCode(config QrConfig) string {
	u := "otpauth://totp/" + config.AccountName
	query := "?secret=" + url.QueryEscape(config.Secret)
	if config.Issuer != "" {
		query = query + url.QueryEscape("&issuer="+config.Issuer)
	}
	query = query + url.QueryEscape("&algorithm="+Config.Algorithm.String()+"&digits="+Config.Digits.String()+"&period="+fmt.Sprint(Config.Period))
	u = u + query
	baseUrl := "https://chart.googleapis.com/chart?chs=200x200&chld=M|0&cht=qr&chl="
	return baseUrl + u
}
