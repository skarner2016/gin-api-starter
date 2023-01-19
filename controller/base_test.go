package controller_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"skarner2016/gin-api-starter/packages/config"
	"skarner2016/gin-api-starter/packages/database/mysql"
	"skarner2016/gin-api-starter/packages/database/redis"
	"skarner2016/gin-api-starter/packages/log"
	"skarner2016/gin-api-starter/router"
	"strings"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
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

	return router.Setup()
}

func SendHttpRequest(method, uri string, body url.Values) *httptest.ResponseRecorder {
	router := Setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, uri, strings.NewReader(body.Encode()))

	switch method {
	case http.MethodGet:
		req.Header.Add("Content-type", "application/json")
	default:
		req.Header.Add("Content-type", "application/x-www-form-urlencoded")
	}
	// req.Header.Add("x-token", "123")

	router.ServeHTTP(w, req)

	return w
}
