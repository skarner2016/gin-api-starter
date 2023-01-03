package redis_test

import (
	"skarner2016/gin-api-starter/packages/config"
	"skarner2016/gin-api-starter/packages/database/redis"
	"testing"
)

func TestSetup(t *testing.T) {
	config.Setup()

	redis.Setup()

	redis.GetRedis(redis.InstanceDefault)
}
