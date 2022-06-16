package main

import (
	"ginEssential/controller"
	"ginEssential/middleware"
	"github.com/gin-gonic/gin"
)

// CollectRouter 为gin引擎添加路由监听
func CollectRouter(r *gin.Engine) *gin.Engine {

	// ping请求
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Hello Gin!",
		})
	})

	// 注册请求
	r.POST("/api/auth/register", controller.Register)

	// 登录路由
	r.POST("api/auth/login", controller.Login)

	// 查询用户信息路由
	r.GET("api/auth/info", middleware.AuthMiddleware(), controller.Info)

	return r

}
