package router

import (
	"skarner2016/gin-api-starter/controller"
	"skarner2016/gin-api-starter/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(
		// middleware.CORSMiddleware(),
		middleware.RecoveryMiddleware(),
	)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r = setupFileRouter(r)

	test := r.Group("/test")
	{
		testController := controller.NewTestController()
		test.GET("", testController.Index)
		test.GET("/user", testController.User)
	}

	return r
}
