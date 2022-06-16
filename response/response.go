package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 用于对相应的一个封装
func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})

}

// Success 成功的响应的封装
func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

// Fail 失败的响应的封装
func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)

}
