package router

import (
	"skarner2016/gin-api-starter/controller"
	"skarner2016/gin-api-starter/middleware"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.Use(
		// gin.Logger(),
		// gin.Recovery(),

		// middleware.CORSMiddleware(),
		middleware.RecoveryMiddleware(),
		middleware.LoggerMiddleware(),
	)

	r = setupFileRouter(r)

	test := r.Group("/test")
	{
		testController := controller.NewTestController()
		test.GET("", testController.Index)
		test.GET("/user", testController.GetUser)
		test.POST("/user", testController.CreateUser)
	}

	return r
}
