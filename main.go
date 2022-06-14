package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// User 数据库结构体
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(12);not null"`
	Telephone string `gorm:"varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

func main() {

	// 初始化数据库
	db := InitDB()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Hello Gin!",
		})
	})

	// 注册请求
	r.POST("/api/auth/register", func(ctx *gin.Context) {

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
			name = generateName(8)
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
		if isTelephoneExist(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号已被注册"})
			return
		}

		// 创建用户

		// 返回结果
		log.Println("创建用户")
		log.Println(name, telephone, password)
		ctx.JSON(200, gin.H{
			"msg":       "register success",
			"name":      name,
			"telephone": telephone,
			"password":  password,
		})

	})

	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}

func generateName(length int) string {
	// 生成length位随机用户名

	// 字母表
	var letters = []byte("abcdefghijklmnopqrstuvwxyz")
	// 创建length长度字节流
	result := make([]byte, length)

	// 以微秒作为种子，如果注册太快可能会出现相同用户名
	rand.Seed(time.Now().UnixMicro())
	// 随机选择字母表中的字母
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

func InitDB() *gorm.DB {
	// 初始化连接池

	// args
	host := "localhost"
	port := "3306"
	user := "root"
	password := "admin"
	database := "ginEssential"
	charset := "utf8mb4"

	// DSN
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", user, password, host, port, database, charset)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})

	if err != nil {
		panic("database connect err:" + err.Error())
	}

	return db

}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	// 判断电话是否已存在

	// 创建User结构对象
	var user User
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
