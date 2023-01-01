package router

import "github.com/gin-gonic/gin"

func setupFileRouter(r *gin.Engine) *gin.Engine {
	r.StaticFile("/robots.txt", "./resources/robots.txt")
	r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	return r
}
