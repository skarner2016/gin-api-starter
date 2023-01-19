package main

import (
	"skarner2016/gin-api-starter/packages/config"
	"skarner2016/gin-api-starter/packages/database/mysql"
	"skarner2016/gin-api-starter/packages/database/redis"
	"skarner2016/gin-api-starter/packages/log"
	"skarner2016/gin-api-starter/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// config
	config.Setup()

	// log
	log.Setup()

	mode := config.APPConfig.GetString("mode")
	gin.SetMode(mode)

	// mysql
	mysql.Setup()

	// redis
	redis.Setup()

	router := router.Setup()

	addr := config.APPConfig.GetString("addr")
	router.Run(addr)
}
