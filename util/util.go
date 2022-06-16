package util

import (
	"math/rand"
	"time"
)

// GenerateName 生成length位随机用户名
func GenerateName(length int) string {

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
