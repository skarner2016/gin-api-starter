package main

import (
	"io"
	"os"
	"skarner2016/gin-api-starter/packages/config"
	"skarner2016/gin-api-starter/packages/database/mysql"
	"skarner2016/gin-api-starter/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Setup()

	mode := config.APPConfig.GetString("mode")
	gin.SetMode(mode)

	sqlDB := mysql.Setup()
	defer sqlDB.Close()

	f, _ := os.Create("gin.log")

	if mode == "debug" {
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	} else {
		gin.DefaultWriter = io.MultiWriter(f)
	}

	router := router.SetupRouter()

	addr := config.APPConfig.GetString("addr")
	router.Run(addr)
}
