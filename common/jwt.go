package common

import (
	"ginEssential/model"
	"github.com/golang-jwt/jwt"
	"time"
)

// 设置token加密密钥
var jwtKey = []byte("jwt_secret_key")

// Claims 定义token的claim结构体
type Claims struct {
	UserID uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	// 发放token

	// 获取token时间信息
	issue := time.Now()
	expiration := issue.Add(24 * time.Hour)

	// 设置claim
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expiration.Unix(),
			// 发放时间
			IssuedAt: issue.Unix(),
			// 发放人
			Issuer: "ylcao.top",
			// 主题
			Subject: "user token",
		},
	}

	// 设置token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成获取token
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil

}
