package controller

import (
	"dyk/common"
	"dyk/model"
	"dyk/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 检查是否存在相同手机号的用户
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func RegisterController(ctx *gin.Context) {
	DB := common.InitDB()
	//获取参数
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	telephone := ctx.PostForm("telephone")
	// 数据验证  手机
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须是11位"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能小于6位"})
		return
	}
	// 没有输入名称自动随机创建10位用户名
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, password, telephone)
	// 确认电话号码是存在
	if isTelephoneExist(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
		return
	}
	// 创建用户
	newuser := model.User{
		Name:      name,
		Password:  password,
		Telephone: telephone,
	}
	DB.Create(&newuser)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200, "msg": "注册成功",
	})
}
