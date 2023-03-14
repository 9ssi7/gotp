package gotp

type VerifyConfig struct {
	Secret string
	Code   string
	Window int
	Time   int64
}

func Verify(config VerifyConfig) bool {
	if config.Time == 0 {
		config.Time = GetTime()
	}
	for i := -config.Window; i <= config.Window; i++ {
		code, _ := getCode(config.Secret, config.Time+int64(i))
		if code == config.Code {
			return true
		}
	}
	return false
}
