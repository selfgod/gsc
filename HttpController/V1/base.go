package V1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReturnJson(code int, data gin.H, msg string) (ret gin.H) {
	return gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	}
}

func ReturnError(data gin.H) (ret gin.H) {
	return gin.H{
		"code": http.StatusInternalServerError,
		"data": data,
		"msg":  http.StatusText(http.StatusInternalServerError),
	}
}
