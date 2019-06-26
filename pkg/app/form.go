package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"whacos/pkg/e"
)

type CommonForm struct {
	PageNum  int `json:"pageNum" valid:"Required MinSize(1)"`
	PageSize int `json:"pageSize" valid:"Required MinSize(1)"`
}

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e.Success
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, e.Error
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.InvalidParams
	}

	return http.StatusOK, e.Success
}
