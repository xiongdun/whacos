package c_log

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"whacos/pkg/e"
)

// @Tags 系统日志相关
// @Resource Name
// @Summary 按ID 获取系统日志详情
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/log/get/{id} [GET]
func GetLog(c *gin.Context) {

}

// @Tags 系统日志相关
// @Resource Name
// @Summary 获取系统日志列表
// @Accept json
// @Produce json
// @Param listRole body c_role.ListRoleForm true "listRole"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/log/list/ [GET]
func ListLog(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": e.Success,
		"msg":  "请求成功",
	})
}

// @Tags 系统日志相关
// @Resource Name
// @Summary 新增系统日志信息
// @Accept json
// @Produce json
// @Param listRole body c_role.ListRoleForm true "listRole"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/log/add [POST]
func AddLog(c *gin.Context) {

}

// @Tags 系统日志相关
// @Resource Name
// @Summary 修改系统日志信息
// @Accept json
// @Produce json
// @Param listRole body c_role.ListRoleForm true "listRole"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/log/edit [PUT]
func EditLog(c *gin.Context) {

}

// @Tags 系统日志相关
// @Resource Name
// @Summary 按ID 删除角色
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/log/remove/{id} [DELETE]
func RemoveLog(c *gin.Context) {

}
