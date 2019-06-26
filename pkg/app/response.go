package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"whacos/pkg/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: httpCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}

// Response setting gin.JSON success
func (g *Gin) ResponseSuccess(data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code: e.Success,
		Msg:  e.GetMsg(e.Success),
		Data: data,
	})
	return
}

// Response setting gin.JSON fail
func (g *Gin) ResponseFail(data interface{}) {
	g.C.JSON(http.StatusInternalServerError, Response{
		Code: e.Error,
		Msg:  e.GetMsg(e.Error),
		Data: data,
	})
	return
}

// Response setting gin.JSON invalid params
func (g *Gin) ResponseInvalidParams(data interface{}) {
	g.C.JSON(http.StatusBadRequest, Response{
		Code: e.InvalidParams,
		Msg:  e.GetMsg(e.InvalidParams),
		Data: data,
	})
	return
}
