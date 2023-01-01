package main

import (
	"io"
	"os"
	"skarner2016/gin-api-starter/packages/config"
	"skarner2016/gin-api-starter/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.SetupConfig()

	mode := config.APPConfig.GetString("mode")
	gin.SetMode(mode)

	f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := router.SetupRouter()

	addr := config.APPConfig.GetString("addr")
	router.Run(addr)
}
