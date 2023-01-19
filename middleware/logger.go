package middleware

import (
	"skarner2016/gin-api-starter/packages/log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		cost := time.Since(start)

		raw := c.Request.URL.RawQuery
		path := c.Request.URL.Path
		if raw != "" {
			path = path + "?" + raw
		}

		log.GetLogger(log.InstanceGin).Infof(
			"[GIN] %v | %3d | %13v | %15s | %-7s  %#v\n%s",
			start.Format("2006/01/02 - 15:04:05"),
			c.Writer.Status(),
			cost,
			c.ClientIP(),
			c.Request.Method,
			path,
			c.Errors.ByType(gin.ErrorTypePrivate).String(),
		)
	}
}
