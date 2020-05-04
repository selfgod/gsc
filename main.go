/**
* main.go
 */
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gsc/Base"
	"gsc/HttpController/V1"
	"net/http"
)

func main() {
	config := Base.LoadConfig()
	fmt.Printf("config:%#v\n", config)
	Base.InitDb(config.Db)
	r := gin.Default()
	V1.Routers(r)
	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, "404 page not found")
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
