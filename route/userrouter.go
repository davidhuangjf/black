package route

import (
	"dyk/controller"
	"dyk/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) *gin.Engine {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word")
	})

	r.POST("/api/auth/register", controller.RegisterController)
	r.POST("/api/auth/login", controller.LoginController)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.InfoController)
	// r.GET("/api/auth/info", controller.InfoController)

	return r
}
