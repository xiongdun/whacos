package c_dict

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"whacos/pkg/e"
)

// @Tags 数据字典相关
// @Resource Name
// @Summary 按ID 获取数据字典详情
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/dict/get/{id} [GET]
func GetDict(c *gin.Context) {

}

// @Tags 数据字典相关
// @Resource Name
// @Summary 获取数据字典列表
// @Accept json
// @Produce json
// @Param listRole body c_role.ListRoleForm true "listRole"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/dict/list/ [GET]
func ListDict(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": e.Success,
		"msg":  "请求成功",
	})
}

// @Tags 数据字典相关
// @Resource Name
// @Summary 新增数据字典信息
// @Accept json
// @Produce json
// @Param listRole body c_role.ListRoleForm true "listRole"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/dict/add [POST]
func AddDict(c *gin.Context) {

}

// @Tags 数据字典相关
// @Resource Name
// @Summary 修改数据字典信息
// @Accept json
// @Produce json
// @Param listRole body c_role.ListRoleForm true "listRole"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/dict/edit [PUT]
func EditDict(c *gin.Context) {

}

// @Tags 数据字典相关
// @Resource Name
// @Summary 按ID 删除角色
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/dict/remove/{id} [DELETE]
func RemoveDict(c *gin.Context) {

}
