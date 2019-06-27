package c_role

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"whacos/pkg/e"
)

type ListRoleForm struct {
	RoleName string `json:"roleName"`
}

// @Tags 系统角色相关
// @Resource Name
// @Summary 按ID 获取角色
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/role/get/{id} [GET]
func GetRole(c *gin.Context) {

}

// @Tags 系统角色相关
// @Resource Name
// @Summary 获取角色列表
// @Accept json
// @Produce json
// @Param listRole body c_role.ListRoleForm true "listRole"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/role/list/ [GET]
func ListRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": e.Success,
		"msg":  "请求成功",
	})
}

// @Tags 系统角色相关
// @Resource Name
// @Summary 新增角色信息
// @Accept json
// @Produce json
// @Param listRole body c_role.ListRoleForm true "listRole"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/role/add [POST]
func AddRole(c *gin.Context) {

}

// @Tags 系统角色相关
// @Resource Name
// @Summary 修改角色信息
// @Accept json
// @Produce json
// @Param listRole body c_role.ListRoleForm true "listRole"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/role/edit [PUT]
func EditRole(c *gin.Context) {

}

// @Tags 系统角色相关
// @Resource Name
// @Summary 按ID 删除角色
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/role/remove/{id} [DELETE]
func RemoveRole(c *gin.Context) {

}
