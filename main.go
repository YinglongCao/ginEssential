package main

import (
	"ginEssential/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {

	// 获取配置
	InitConfig()

	// 初始化数据库
	common.InitDB()

	// 获取默认引擎并初始化路由
	r := CollectRouter(gin.Default())

	// 设置监听端口
	port := viper.GetString("server.port")

	if port != "" {
		panic(r.Run(":" + port))
	}

	// 监听运行
	panic(r.Run(":8080")) // listen and serve on 0.0.0.0:8080

}

// InitConfig 初始化配置文件
func InitConfig() {

	// 获取运行路径
	workDir, _ := os.Getwd()

	// 设置配置文件名称
	viper.SetConfigName("application")

	// 设置配置文件类型
	viper.SetConfigType("yml")

	// 设置配置文件路径
	viper.AddConfigPath(workDir + "/config")

	err := viper.ReadInConfig()

	if err != nil {
		panic("InitConfig() 配置文件读取错误")
	}

}
