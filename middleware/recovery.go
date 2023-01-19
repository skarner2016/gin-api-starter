package middleware

import (
	"skarner2016/gin-api-starter/packages/log"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.GetLogger(log.InstanceDefault).Error(err)
			}
		}()

		ctx.Next()
	}
}
