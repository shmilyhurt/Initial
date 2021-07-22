package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseCode int

type Response struct {
	Code ResponseCode `json:"code"`
	Data interface{}  `json:"data"`
	Msg  string       `json:"msg"`
}

func Result(c *gin.Context, code ResponseCode, data interface{}, msg string) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func SuccessResponse(c *gin.Context) {
	Result(c, http.StatusOK, map[string]interface{}{}, "success")
}

func SuccessResponseWithMessage(c *gin.Context, message string) {
	Result(c, http.StatusOK, map[string]interface{}{}, message)
}

func SuccessResponseWithData(c *gin.Context, data interface{}) {
	Result(c, http.StatusOK, data, "success")
}

func SuccessResponseDetailed(c *gin.Context, data interface{}, message string) {
	Result(c, http.StatusOK, data, message)
}

func FailResponse(c *gin.Context, code ResponseCode, message string) {
	Result(c, code, map[string]interface{}{}, message)
}

func FailWithDetailed( c *gin.Context, code ResponseCode, data interface{}, message string,) {
	Result(c, code, data, message)
}