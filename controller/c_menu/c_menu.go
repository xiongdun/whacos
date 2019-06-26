package c_menu

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"whacos/pkg/e"
)

func GetMenu(c *gin.Context) {

}

func ListMenu(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": e.Success,
		"msg":  "请求成功",
	})
}

func AddMenu(c *gin.Context) {

}

func EditMenu(c *gin.Context) {

}

func RemoveMenu(c *gin.Context) {

}
