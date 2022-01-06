package main

import (
	"dyk/route"

	"github.com/gin-gonic/gin"
)

func main() {
	// DB := common.InitDB()
	r := gin.Default()
	r = route.UserRoute(r)

	// r.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "hello word")
	// })
	// r.POST("/api/auth/register", controller.RegisterController)

	//监听端口默认为8080
	panic(r.Run())

}
