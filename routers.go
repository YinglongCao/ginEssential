package main

import (
	"ginEssential/controller"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	// 为gin引擎添加路由监听

	// ping请求
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Hello Gin!",
		})
	})

	// 注册请求
	r.POST("/api/auth/register", controller.Register)

	return r

}
