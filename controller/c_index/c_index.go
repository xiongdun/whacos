package c_index

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"whacos/pkg/app"
)

type LoginForm struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Tags 登录校验
// @Resource Name
// @Summary 跳转到首页
// @Accept json
// @Produce json
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /index [GET]
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}

func ToLogin(c *gin.Context) {

}

// @Tags 登录校验
// @Resource Name
// @Summary 用户登录
// @Accept json
// @Produce json
// @Param login body c_index.LoginForm true "loginForm"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /login [POST]
func Login(c *gin.Context) {
	appGin := app.Gin{C: c}
	user := LoginForm{}
	if c.BindJSON(&user) == nil {
		appGin.ResponseInvalidParams(nil)
		return
	}
}
