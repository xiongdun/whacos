package c_file

import (
	"github.com/gin-gonic/gin"
	"whacos/pkg/app"
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
