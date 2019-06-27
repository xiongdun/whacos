package c_menu

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"whacos/pkg/e"
)

type ListMenuForm struct {
	MenuName string `json:"menuName"`
}

// @Tags 系统菜单相关
// @Resource Name
// @Summary 按ID 获取菜单
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/menu/get/{id} [GET]
func GetMenu(c *gin.Context) {

}

// @Tags 系统菜单相关
// @Resource Name
// @Summary 查询菜单信息列表
// @Accept json
// @Produce json
// @Param listMenu body c_menu.ListMenuForm true "listMenu"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/menu/list [GET]
func ListMenu(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": e.Success,
		"msg":  "请求成功",
	})
}

// @Tags 系统菜单相关
// @Resource Name
// @Summary 新增菜单信息
// @Accept json
// @Produce json
// @Param listMenu body c_menu.ListMenuForm true "listMenu"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/menu/add [POST]
func AddMenu(c *gin.Context) {

}

// @Tags 系统菜单相关
// @Resource Name
// @Summary 修改菜单信息
// @Accept json
// @Produce json
// @Param listMenu body c_menu.ListMenuForm true "listMenu"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/menu/edit [PUT]
func EditMenu(c *gin.Context) {

}

// @Tags 系统菜单相关
// @Resource Name
// @Summary 按ID 删除菜单
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/menu/remove/{id} [DELETE]
func RemoveMenu(c *gin.Context) {

}
