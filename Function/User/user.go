package User

const (
	userNotFound  = 100
	passwordError = 101
	tokenError    = 102
	tokenExpire   = 103
	ok            = 200
)

var statusText = map[int]string{
	userNotFound:  "用户不存在",
	passwordError: "密码错误",
	tokenError:    "token错误",
	tokenExpire:   "会话超时",
	ok:            "ok",
}

func CheckToken(token string) (ret bool, code int, err error) {
	ret = false
	err = nil
	code = ok
	return
}

func Login(username, password string) (token string, code int, err error) {
	//err=errors.New("test")
	token = ""
	code = ok
	err = nil
	return
}

func StatusText(code int) string {
	return statusText[code]
}
