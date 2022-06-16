package main

import (
	"ginEssential/common"
	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化数据库
	common.InitDB()

	// 获取默认引擎
	r := gin.Default()

	// 初始化路由
	r = CollectRouter(r)

	// 监听运行
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}
