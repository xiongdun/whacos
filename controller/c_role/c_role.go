package c_role

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"whacos/pkg/e"
)

func GetRole(c *gin.Context) {

}

func ListRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": e.Success,
		"msg":  "请求成功",
	})
}

func AddRole(c *gin.Context) {

}

func EditRole(c *gin.Context) {

}

func RemoveRole(c *gin.Context) {

}
