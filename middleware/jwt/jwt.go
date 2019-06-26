package jwt

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"whacos/pkg/e"
	"whacos/pkg/utils"
)

// 获取token 接口
func ValidJWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code int
		var data interface{}

		code = e.Success
		token := context.Query("token")

		if token == "" {
			code = e.InvalidParams
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}

		if code != e.Success {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			context.Abort()
			return
		}

		context.Next()
	}
}
