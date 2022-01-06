package middleware

import (
	"dyk/common"
	"dyk/model"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		// 获取 authentication header
// 		tokenString := ctx.GetHeader("Authorization")

// 		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
// 			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
// 			ctx.Abort()
// 			return
// 		}
// 		// log.println(tokenString)
// 		tokenString = tokenString[7:]

// 		token, claims, err := common.ParseToken(tokenString)
// 		if err != nil || !token.Valid {
// 			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
// 			ctx.Abort()
// 			return
// 		}
// 		// 通过验证后获取
// 		userId := claims.UserId
// 		DB := common.GetDB()
// 		var user model.User
// 		DB.First(&user, userId)
// 		// log.Printf(user.Name)
// 		if user.ID == 0 {
// 			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "用户已不存在"})
// 			ctx.Abort()
// 			return
// 		}
// 		ctx.Set("user", user)
// 		log.Printf(user.Name)

// 	}
// }

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 authentication header
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		log.Printf("=======================")
		log.Printf("=======================")
		log.Printf("=======================")

		// 通过验证后获取
		userId := claims.UserId
		DB := common.InitDB()
		var user model.User
		DB.First(&user, userId)

		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "用户已不存在"})
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
