package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	code 错误码
	message 错误信息
	data 数据
 */


type Response struct {
	Code int	`json:"code"`
	Message interface{}		`json:"message"`
	Data interface{}	`json:"data"`
}
type Code int64

const(
	CodeSuccess	Code = 1000 + iota
	CodeInvalidParams
	CodeError
	CodeUserError

)
var codeMsgMap = map[Code]string{
	CodeSuccess:"请求成功",
	CodeInvalidParams:"请求参数错误",
	CodeError:"服务器繁忙",
	CodeUserError:"用户名或密码不正确",
}

func ResponseValid(c *gin.Context,code int)  {
	rd := &Response{
		Code: code,
		Message: codeMsgMap[CodeInvalidParams],
		Data:nil,
	}
	c.JSON(http.StatusOK,rd)
}
func ResponseError(c *gin.Context,code int)  {
	rd := &Response{
		Code: code,
		Message: codeMsgMap[CodeError],
		Data:nil,
	}
	c.JSON(http.StatusOK,rd)
}

func ResponseSuccess(c *gin.Context,data interface{})  {
	rd := &Response{
		Code: 200,
		Message: codeMsgMap[CodeSuccess],
		Data:data,
	}
	c.JSON(http.StatusOK,rd)
}

func ResponseUserError(c *gin.Context,code int)  {
	rd := &Response{
		Code: code,
		Message: codeMsgMap[CodeUserError],
	}
	c.JSON(http.StatusOK,rd)
}

func ResponseString(c *gin.Context,code int,msg string)  {
	rd := &Response{
		Code: code,
		Message: msg,
	}
	c.JSON(http.StatusOK,rd)
}