package controller

import (
	"dyk/common"
	"dyk/dto"
	"dyk/model"
	"dyk/response"
	"dyk/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	// DB := common.GetDB()
	DB := common.InitDB()
	//获取参数

	// var requestUser = model.User{}
	// // json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	// gin.Bind(&requestUser)
	// name := requestUser.Name
	// password := requestUser.Password
	// telephone := requestUser.Telephone
	// =====================================================
	// json := model.User{}
	// ctx.BindJSON(&json)
	// name := json.Name
	// password := json.Password
	// telephone := json.Telephone
	// =====================================================
	var requestUser = model.User{}
	json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	fmt.Println(requestUser.Name, requestUser.Password, requestUser.Telephone)
	name := requestUser.Name
	password := requestUser.Password
	telephone := requestUser.Telephone

	// name := ctx.PostForm("name")
	// password := ctx.PostForm("password")
	// telephone := ctx.PostForm("telephone")

	// 数据验证  手机
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须是11位")
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须是11位"})
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能小于6位"})
		return
	}
	// 没有输入名称自动随机创建10位用户名
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, password, telephone)
	// 确认电话号码是存在
	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
		return
	}
	// 密码加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		// ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密错误"})
		return
	}
	// 创建用户
	newuser := model.User{
		Name:      name,
		Password:  string(hashPassword),
		Telephone: telephone,
	}
	DB.Create(&newuser)

	response.Success(ctx, nil, "注册成功")
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"code": 200, "msg": "注册成功",
	// })
}

func LoginController(ctx *gin.Context) {
	DB := common.InitDB()
	// 获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// 数据合规性验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须是11位")
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须是11位"})
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能小于6位"})
		return
	}
	// 判断手机号是否存在
	var user model.User
	DB.Where("telephone =?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		// ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	// 发放TOKEN
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统内部错误")
		// ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统内部错误"})
		log.Printf("token generate error : %v", err)
		return
	}
	// token := "11"

	// 登录成功
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"code": 200,
	// 	"data": gin.H{"token": token},
	// 	"msg":  "注册成功",
	// })
	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

func InfoController(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{"user": dto.ToUserDto(user.(model.User))}, "")
	// ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}
