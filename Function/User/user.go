package User

import (
	"gsc/Base"
	"gsc/Function"
	"gsc/Model/User"
	"time"
)

const (
	userNotFound  = 100
	passwordError = 101
	tokenError    = 102
	tokenExpire   = 103
	ok            = 200
	userExists    = 201
	notOk         = 500
)

var statusText = map[int]string{
	userNotFound:  "用户不存在",
	userExists:    "用户名已被占用",
	passwordError: "密码错误",
	tokenError:    "token错误",
	tokenExpire:   "会话超时",
	ok:            "ok",
	notOk:         "系统错误",
}

func CheckToken(token string) (ret bool, code int, err error) {
	ret = false
	err = nil
	code = ok
	return
}

func Login(name, pass string) (token string, code int, err error) {
	code = ok
	u := &User.User{
		Name: name,
		Pass: Function.Md5V(pass),
	}
	c, err := u.Login()
	if err != nil {
		return
	}
	if c == false {
		code = passwordError
		return
	}

	token, expire := getToken(name)
	err = u.SetToken(token, expire)
	return
}

func getToken(name string) (token string, expire int64) {
	expire = time.Now().Unix() + 3600
	token = Function.Md5V(string(expire) + name + Base.GConfig.Secret)
	return
}

func CreateUser(name, pass string) (code int, err error) {
	code = ok
	u := &User.User{
		Name: name,
		Pass: Function.Md5V(pass),
	}
	c, err := u.NameExists()
	if err != nil {
		return
	}
	if c == true {
		code = userExists
		return
	}
	//判断用户名是否存在
	err = u.Add()
	return
}

func StatusText(code int) string {
	return statusText[code]
}
