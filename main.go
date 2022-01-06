package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 生成随机名称
func RandomString(n int) string {
	letters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"type:varchar(128);not null"`
	// Password2 string `gorm:"type:varchar(128)"`
}

func InitDB() *gorm.DB {
	// username := viper.GetString("datasource.username")
	// // username := "root"
	// password := viper.GetString("datasource.password")
	// host := viper.GetString("datasource.host")
	// port := viper.GetString("datasource.port")
	// database := viper.GetString("datasource.database")
	// charset := viper.GetString("datasource.charset")

	username := "root"
	password := "Admwork0620"
	host := "localhost"
	port := "3306"
	database := "gin"
	charset := "utf8mb4"
	// driverName := "mysql"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	// 自动创建数据表
	db.AutoMigrate(&User{})

	return db
}

// 检查是否存在相同手机号的用户
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

// 	DB.AutoMigrate(&models.User{})
// 	DB.AutoMigrate(&models.WorkPolicy{})
// 	return DB
// }

// func GetDB() *gorm.DB {
// 	return DB
// }

//

func main() {
	db := InitDB()
	// defer db.Close()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word")
	})

	r.POST("/api/auth/register", func(ctx *gin.Context) {
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
			name = RandomString(10)
		}
		log.Println(name, password, telephone)
		// 确认电话号码是存在
		if isTelephoneExist(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
			return
		}
		// 创建用户
		newuser := User{
			Name:      name,
			Password:  password,
			Telephone: telephone,
		}
		db.Create(&newuser)

		ctx.JSON(http.StatusOK, gin.H{
			"code": 200, "msg": "注册成功",
		})
	})

	//监听端口默认为8080
	panic(r.Run())

}
