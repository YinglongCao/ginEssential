package middleware

import (
	"ginEssential/common"
	"ginEssential/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	// 自定义中间件, 保护路由

	return func(ctx *gin.Context) {
		// 获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		// 验证token格式
		if len(tokenString) == 0 || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			// Abort阻止调用挂起的处理程序。
			//注意，这不会停止当前处理程序。
			//假设您有一个验证当前请求是否被授权的授权中间件。
			//如果授权失败(例如:密码不匹配)，
			//调用Abort以确保没有调用此请求的其余处理程序。
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		// 证明token通过了验证 获取claims中的userID
		userID := claims.UserID
		db := common.GetDB()
		var user model.User

		db.First(&user, userID)

		// 验证数据库
		// 查找不到用户
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		// 用户存在 将user信息写入上下文
		ctx.Set("user", user)

		// Next应该只在中间件内部使用。它执行调用处理程序内部链中的挂起处理程序。参见GitHub中的示例。
		ctx.Next()

	}

}
