package gotp

import (
	"math/rand"
	"time"
)

// Config is the configuration for the package.
// It can be changed to customize the package.
// The default configuration is:
//
//	SecretSize:    16,
//	Period:        30,
//	Algorithm:     AlgorithmSHA1,
//	Digits:        DigitsSix,
//	ValidChars:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567",
var Config = config{
	// Digits is the number of digits in the code.
	// It can be 6 or 8.
	// The default is 6.
	Digits: DigitsSix,

	// SecretSize is the length of the secret key.
	// The default is 16.
	SecretSize: 16,

	// Period is the length of the time period.
	Period: 30,

	// Algorithm is the hash function used to generate the code.
	// It can be sha1.New, sha256.New, or sha512.New.
	Algorithm: AlgorithmSHA1,

	// ValidChars is the set of characters used in the secret key.
	ValidChars: "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567",
}

// CreateSecret returns a random secret key.
// The secret key is a string of characters from the set A-Z, 2-7.
// The default length of the secret key is 16 characters.
// This can be changed by setting Config.SecretSize.
// The default set of characters can be changed by setting Config.ValidChars.
func CreateSecret() string {
	validChars := []byte(Config.ValidChars)
	secret := make([]byte, Config.SecretSize)
	for i := 0; i < Config.SecretSize; i++ {
		secret[i] = validChars[rand.Intn(len(validChars))]
	}
	return string(secret)
}

// GetCode returns the current code for the secret key.
// The code is valid for the current period.
// The period is 30 seconds by default.
// This can be changed by setting Config.Period.
func GetCode(secret string, time int64) string {
	code, _ := getCode(secret, time)
	return code
}

// GetTime returns the current time in Unix time divided by the period.
// This is used to generate the counter for the HOTP algorithm.
// The period is 30 seconds by default.
// This can be changed by setting Config.Period.
// The Unix time is divided by the period to make it easier to test.
// This way, the counter can be set to any value.
func GetTime() int64 {
	return time.Now().Unix() / Config.Period
}
