package controller

import (
	"net/http"
	"skarner2016/gin-api-starter/packages/log"
	"skarner2016/gin-api-starter/packages/response"
	"skarner2016/gin-api-starter/packages/validate"

	"github.com/gin-gonic/gin"
)

type TestController struct {
}

func NewTestController() *TestController {
	return &TestController{}
}

func (c *TestController) Index(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"hello": "gin",
	})
}

func (c *TestController) Redirect(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, "http://www.baidu.com")
}

type GetUserParam struct {
	// ID   int64  `form:"id" binding:"required,gte=1,lte=999"`
	// Name string `form:"name" binding:"required"`

	ID   int64  `form:"id" validate:"required,gte=1,lte=999"`
	Name string `form:"name" validate:"required"`
}

func (con *TestController) GetUser(c *gin.Context) {
	form := &GetUserParam{}
	if err := c.ShouldBind(&form); err != nil {
		response.Fail(c, 400, nil)
		return
	}

	msg, err := validate.GetValidateError(form)
	if err != nil {
		log.GetLogger(log.InstanceApp).Error(msg)
		response.FailValidateeMsg(c, msg)
		return
	}

	res := gin.H{
		"form": form,
	}

	response.Success(c, res)

	// user := &models.User{}
	// db := mysql.GetDB(mysql.InstanceDefault)
	// err := db.Limit(1).Find(&user).Error
	// if err != nil {
	// 	response.Fail(ctx, 500, nil)
	// }

	// res := gin.H{
	// 	"user": user,
	// }

	// response.Success(ctx, res)
}

type CreateUserParam struct {
	// ID   int64  `form:"id" binding:"required,gte=1,lte=999"`
	// Name string `form:"name" binding:"required"`

	ID   int64  `form:"id" validate:"required,gte=1,lte=999"`
	Name string `form:"name" validate:"required"`
}

func (con *TestController) CreateUser(c *gin.Context) {
	form := &CreateUserParam{}
	if err := c.ShouldBind(&form); err != nil {
		response.Fail(c, 400, nil)
		return
	}

	msg, err := validate.GetValidateError(form)
	if err != nil {
		log.GetLogger(log.InstanceApp).Error(msg)
		response.FailValidateeMsg(c, msg)
		return
	}

	res := gin.H{
		"form": form,
	}

	response.Success(c, res)
}
