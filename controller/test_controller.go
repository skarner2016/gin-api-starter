package controller

import (
	"net/http"
	"skarner2016/gin-api-starter/models"
	"skarner2016/gin-api-starter/packages/database/mysql"
	"skarner2016/gin-api-starter/packages/response"

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

func (c *TestController) User(ctx *gin.Context) {
	user := &models.User{}

	db := mysql.GetDB()
	err := db.Limit(1).Find(&user).Error
	if err != nil {
		response.Fail(ctx, 500, nil)
	}

	res := gin.H{
		"user": user,
	}

	response.Success(ctx, res, "success")
}
