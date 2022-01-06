package main

import (
	"dyk/route"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	// DB := common.InitDB()
	// defer DB.Close()
	r := gin.Default()
	r = route.UserRoute(r)

	//监听端口默认为8080
	port := viper.GetString("server.port")
	panic(r.Run(":" + port))

}

// 读取yaml配置文件
func InitConfig() {
	path, _ := os.Getwd()
	workDir := filepath.Join(path, "config")
	viper.AddConfigPath(workDir)       //设置读取的文件路径
	viper.SetConfigName("application") //设置读取的文件名
	viper.SetConfigType("yaml")        //设置文件的类型
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
