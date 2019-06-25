package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"whacos/pkg/err"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	e := c.Bind(form)
	if e != nil {
		return http.StatusBadRequest, err.Success
	}

	valid := validation.Validation{}
	check, e := valid.Valid(form)
	if e != nil {
		return http.StatusInternalServerError, err.Error
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, err.InvalidParams
	}

	return http.StatusOK, err.Success
}
