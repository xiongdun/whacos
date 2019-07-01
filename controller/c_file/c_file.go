package c_file

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"whacos/pkg/app"
	"whacos/pkg/e"
)

// @Tags APP相关
// @Resource Name
// @Summary 上传文件
// @Accept json
// @Produce json
// @Param file query file true "file"
// @Param token query string true "token"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /app/file/upload [POST]
func UploadFile(c *gin.Context) {

	formFile, err := c.FormFile("file")
	appGin := app.Gin{C: c}
	if err != nil {
		appGin.ResponseInvalidParams(err.Error())
		return
	}

	if err := c.SaveUploadedFile(formFile, "this.jpg"); err != nil {
		appGin.ResponseFail(nil)
	} else {
		appGin.ResponseSuccess(true)
	}
}

// @Tags 文件相关
// @Resource Name
// @Summary 按ID 获取文件详情
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/file/get/{id} [GET]
func GetFile(c *gin.Context) {

}

// @Tags 文件相关
// @Resource Name
// @Summary 获取文件列表
// @Accept json
// @Produce json
// @Param listRole body c_role.ListRoleForm true "listRole"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/file/list/ [GET]
func ListFile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": e.Success,
		"msg":  "请求成功",
	})
}

// @Tags 文件相关
// @Resource Name
// @Summary 新增文件信息
// @Accept json
// @Produce json
// @Param listRole body c_role.ListRoleForm true "listRole"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/file/add [POST]
func AddFile(c *gin.Context) {

}

// @Tags 文件相关
// @Resource Name
// @Summary 修改文件信息
// @Accept json
// @Produce json
// @Param listRole body c_role.ListRoleForm true "listRole"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/file/edit [PUT]
func EditFile(c *gin.Context) {

}

// @Tags 文件相关
// @Resource Name
// @Summary 按ID 删除文件
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /sys/file/remove/{id} [DELETE]
func RemoveFile(c *gin.Context) {

}
