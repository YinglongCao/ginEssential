package util

import (
	"ginEssential/model"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func GenerateName(length int) string {
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
