package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func response(c *gin.Context, httpStatus int, code int64, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func Success(c *gin.Context, data gin.H, msg string) {
	response(c, http.StatusOK, 200, data, msg)
}

func SuccessMsg(c *gin.Context, code int64) {
	// TODO:
	msg := ""

	response(c, http.StatusOK, code, nil, msg)
}

func Fail(c *gin.Context, code int64, data gin.H) {
	// TODO:
	msg := ""

	response(c, http.StatusOK, code, data, msg)
}

func FailValidateeMsg(c *gin.Context, msg string) {
	response(c, http.StatusOK, 401, nil, msg)
}
