package V1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	v = "/api/v1"
)

func Routers(r *gin.Engine) {

	//r.Use(CheckToken)
	r.Any(v+"/user", func(context *gin.Context) {
		switch context.Request.Method {
		case http.MethodGet:

		case http.MethodPost:
			CreateUser(context)
		}
	})

	r.POST(v+"/login", Login)

}
