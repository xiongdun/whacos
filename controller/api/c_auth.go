package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"whacos/models/sys"
	"whacos/pkg/err"
	"whacos/pkg/logging"
	"whacos/pkg/utils"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary 校验用户并获取token 值
// @Produce  json
// @Param query username body string true "username"
// @Param query password body string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [POST]
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}

	ok, _ := valid.Valid(&a)
	data := make(map[string]interface{})
	code := err.InvalidParams
	if ok {
		isExist := sys.CheckAuth(username, password)
		if isExist {
			token, e := utils.GenerateToken(username, password)
			if e != nil {
				code = err.ErrorAuthToken
			} else {
				data["token"] = token

				code = err.Success
			}
		} else {
			code = err.ErrorAuth
		}
	} else {
		for _, e := range valid.Errors {
			//log.Println(e.Key, e.Message)

			logging.Info(e.Key, e.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  err.GetMsg(code),
		"data": data,
	})
}
