package gotp

// VerifyConfig is the configuration for the Verify function.
type VerifyConfig struct {
	// Secret is the secret key.
	// It must be a string of valid characters.
	Secret string
	// Code is the code to verify.
	// It must be a string of digits.
	Code string
	// Window is the number of codes to check before and after the current code.
	// If Window is 0, only the current code is checked.
	Window int
}

// Verify verifies the code against the secret key.
// It returns true if the code is valid, false otherwise.
func Verify(config VerifyConfig) bool {
	time := GetTime()
	for i := -config.Window; i <= config.Window; i++ {
		code, _ := getCode(config.Secret, time+int64(i))
		if code == config.Code {
			return true
		}
	}
	return false
}
