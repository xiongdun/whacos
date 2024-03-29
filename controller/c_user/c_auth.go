package c_user

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"whacos/pkg/e"
	"whacos/pkg/logging"
	"whacos/pkg/utils"
	"whacos/service/s_user"
)

type AuthForm struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Tags 登录校验
// @Resource Name
// @Summary 校验用户并获取token 值
// @Description 通过 username 和 password 获取 token
// @Accept json
// @Produce json
// @Param auth body c_user.AuthForm true "authForm"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [POST]
func GetAuth(c *gin.Context) {

	a := AuthForm{}
	y := c.BindJSON(&a)
	if y != nil {
		fmt.Println(y)
	}
	valid := validation.Validation{}

	ok, _ := valid.Valid(&a)
	data := make(map[string]interface{})
	code := e.InvalidParams
	if ok {
		isExist := s_user.CheckAuth(a.Username, a.Password)
		if isExist {
			token, err := utils.GenerateToken(a.Username, a.Password)
			if err != nil {
				code = e.ErrorAuthToken
			} else {
				data["token"] = token

				code = e.Success
			}
		} else {
			code = e.ErrorAuth
		}
	} else {
		for _, err := range valid.Errors {
			//log.Println(e.Key, e.Message)

			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
