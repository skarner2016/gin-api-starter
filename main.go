package main

import (
	"skarner2016/gin-api-starter/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r = InitRouter(r)

	r.Run()
}

func InitRouter(r *gin.Engine) *gin.Engine {
	testController := controller.NewTestController()
	r.GET("api/test", testController.Index)
	r.GET("api/test/redirect", testController.Redirect)

	return r
}
