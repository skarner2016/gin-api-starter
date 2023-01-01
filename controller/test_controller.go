package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestController struct {
}

func NewTestController() *TestController {
	return &TestController{}
}

func (c *TestController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "test",
		"data": "index",
	})
}

func (c *TestController) Redirect(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, "http://www.baidu.com")
}
