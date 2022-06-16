package controller

import (
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {
	// 注册监听

	db := common.GetDB()

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
	if IsTelephoneExist(db, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号已被注册"})
		return
	}

	log.Println("创建用户:")

	// 密码使用哈希保存
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "内部数据错误"})
		log.Printf("Registor() hashedPassword error: %v", err)
		return
	}

	// 创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword),
	}
	db.Create(&newUser)

	log.Println(name, telephone, hashedPassword)

	// 返回结果
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"name":      name,
			"telephone": telephone,
			"password":  hashedPassword,
		},
		"msg": "register success",
	})

}

func Login(ctx *gin.Context) {
	// 登录监听

	db := common.GetDB()

	// 获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "请输入正确的电话号码"})
		return
	}

	// 密码长度必须在8-16之间
	if len(password) < 8 || len(password) > 16 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码必须为8到16位"})
		return
	}

	// 判断手机号是否存在，已存在返回422
	var user model.User
	// 根据电话查询用户结构
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		// 不存在
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "该用户不存在"})
		return
	}

	// 判断密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// 验证失败会产生err，成功返回nil
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "账号和密码错误"})
		return
	}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "内部数据错误"})
		log.Printf("Login() token generate error: %v", err)
		return
	}

	// 返回结果
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登录成功",
	})

}

func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	// 判断电话是否已存在

	// 创建User结构对象
	var user model.User
	// 根据电话查询用户结构
	db.Where("telephone = ?", telephone).First(&user)
	// 判断查询结构
	if user.ID != 0 {
		// 查询到ID非0则存在
		return true
	}
	// 否则不存在
	return false
}
