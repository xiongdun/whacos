package jwt

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"whacos/pkg/err"
	"whacos/pkg/utils"
)

// 获取token 接口
func ValidJWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code int
		var data interface{}

		code = err.Success
		token := context.Query("token")

		if token == "" {
			code = err.InvalidParams
		} else {
			claims, e := utils.ParseToken(token)
			if e != nil {
				code = err.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = err.ErrorAuthCheckTokenTimeout
			}
		}

		if code != err.Success {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  err.GetMsg(code),
				"data": data,
			})
			context.Abort()
			return
		}

		context.Next()
	}
}
