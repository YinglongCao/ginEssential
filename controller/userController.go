package controller

import (
	"fmt"
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {

	db := common.GetDB()

	fmt.Println(db)

	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	// 判断用户名长度
	if len(name) > 12 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户名不能大于6位"})
		return
	}

	// 用户名为空则随机生成
	if len(name) == 0 {
		// 为空则生成8位随机用户名
		name = util.GenerateName(8)
	}

	// 密码长度必须在8-16之间
	if len(password) < 8 || len(password) > 16 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码必须为8到16位"})
		return
	}

	// 电话非11位返回422
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "请输入正确的电话号码"})
		return
	}

	// 电话已存在返回422
	if util.IsTelephoneExist(db, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号已被注册"})
		return
	}

	// 创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	db.Create(&newUser)

	// 返回结果
	log.Println("创建用户")
	log.Println(name, telephone, password)
	ctx.JSON(200, gin.H{
		"msg":       "register success",
		"name":      name,
		"telephone": telephone,
		"password":  password,
	})

}
