package main

import (
	"dyk/route"

	"github.com/gin-gonic/gin"
)

func main() {
	// DB := common.InitDB()
	r := gin.Default()
	r = route.UserRoute(r)

	//监听端口默认为8080
	panic(r.Run())

}
