package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 响应回复

type Response struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Msg string `json:"msg"`
}

const (
	ERROR = 7
	SUCCESS = 0
)

// 响应结果
func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

// 成功响应
func OkWithMessage(message string, c *gin.Context)  {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

// 失败响应
func FailWithMessage(message string, c *gin.Context)  {
	Result(ERROR, map[string]interface{}{}, message, c)
}

// 成功数据返回
func OkWithData(data interface{}, c *gin.Context)  {
	Result(SUCCESS, data, "操作成功", c)
}
