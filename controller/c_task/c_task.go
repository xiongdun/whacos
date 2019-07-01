package c_task

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"whacos/pkg/e"
)

// @Tags 定时任务相关
// @Resource Name
// @Summary 按ID 获取定时任务详情
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/task/get/{id} [GET]
func GetTask(c *gin.Context) {

}

// @Tags 定时任务相关
// @Resource Name
// @Summary 获取定时任务列表
// @Accept json
// @Produce json
// @Param listRole body c_role.ListRoleForm true "listRole"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/task/list/ [GET]
func ListTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": e.Success,
		"msg":  "请求成功",
	})
}

// @Tags 定时任务相关
// @Resource Name
// @Summary 新增定时任务信息
// @Accept json
// @Produce json
// @Param listRole body c_role.ListRoleForm true "listRole"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/task/add [POST]
func AddTask(c *gin.Context) {

}

// @Tags 定时任务相关
// @Resource Name
// @Summary 修改定时任务信息
// @Accept json
// @Produce json
// @Param listRole body c_role.ListRoleForm true "listRole"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/task/edit [PUT]
func EditTask(c *gin.Context) {

}

// @Tags 定时任务相关
// @Resource Name
// @Summary 按ID 删除定时任务
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/task/remove/{id} [DELETE]
func RemoveTask(c *gin.Context) {

}
