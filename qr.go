package gotp

import (
	"net/url"
	"strconv"
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
	v := url.Values{}
	v.Set("secret", config.Secret)
	if config.Issuer != "" {
		v.Set("issuer", config.Issuer)
	}
	v.Set("period", strconv.FormatUint(uint64(Config.Period), 10))
	v.Set("algorithm", Config.Algorithm.String())
	v.Set("digits", Config.Digits.String())
	u := url.URL{
		Scheme:   "otpauth",
		Host:     "totp",
		Path:     config.AccountName,
		RawQuery: v.Encode(),
	}
	baseUrl := "https://chart.googleapis.com/chart?chs=200x200&chld=M|0&cht=qr&chl="
	return baseUrl + u.String()
}
