package route

import (
	"dyk/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) *gin.Engine {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word")
	})

	r.POST("/api/auth/register", controller.RegisterController)
	return r
}
