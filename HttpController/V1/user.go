package V1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gsc/Function/User"
	"net/http"
)

type LoginData struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
type TokenData struct {
	Token string `form:"token" json:"token" binding:"required"`
}

func CheckToken(ctx *gin.Context) {
	var token TokenData
	if err := ctx.ShouldBind(&token); err == nil {
		fmt.Printf("token info:%#v\n", token)
		//gin.H{
		//			"username":     login.Username,
		//			"password": login.Password,
		//		}
		ok, code, err := User.CheckToken(token.Token)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, ReturnError(nil))
			return
		}
		if ok {
			ctx.Next()
			return
		}
		ctx.JSON(http.StatusOK, ReturnJson(code, nil, User.StatusText(code)))
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.Abort()
}

func Login(c *gin.Context) {
	var login LoginData

	if err := c.ShouldBind(&login); err == nil {
		fmt.Printf("login info:%#v\n", login)
		//gin.H{
		//			"username":     login.Username,
		//			"password": login.Password,
		//		}
		token, code, err := User.Login(login.Username, login.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ReturnError(nil))
			return
		}
		c.JSON(http.StatusOK, ReturnJson(code, gin.H{token: token}, User.StatusText(code)))
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
