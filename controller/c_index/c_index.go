package c_index

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Tags DEFAULT
// @Resource Name
// @Summary 跳转到首页
// @Accept json
// @Produce json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /index [GET]
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}

func ToLogin(c *gin.Context) {

}

// @Tags DEFAULT
// @Resource Name
// @Summary 用户登录
// @Accept json
// @Produce json
// @Param loginForm body string true "loginForm"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /login [POST]
func Login(c *gin.Context) {
	user := LoginForm{}
	if c.BindJSON(&user) == nil {
		fmt.Println(user.Password)
		fmt.Println(user.Username)
	}
}
