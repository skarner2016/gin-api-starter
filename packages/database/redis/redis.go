package redis

import (
	"context"
	"fmt"
	"skarner2016/gin-api-starter/packages/config"

	redisGo "github.com/go-redis/redis/v8"
)

var RedisClient *redisGo.Client

var InstanceMap map[Instance]*redisGo.Client

func Setup() {
	confMap := make(map[Instance]*RedisConf, 0)
	if err := config.APPConfig.UnmarshalKey("redis", &confMap); err != nil {
		panic("redis setup: parse config err:" + err.Error())
	}

	InstanceMap = make(map[Instance]*redisGo.Client, 0)
	for i, redisconf := range confMap {
		RedisClient = redisGo.NewClient(&redisGo.Options{
			Addr:     fmt.Sprintf("%s:%d", redisconf.Host, redisconf.Port),
			Password: redisconf.Pass,
			DB:       redisconf.DB,
		})

		ctx := context.Background()
		res, err := RedisClient.Ping(ctx).Result()
		if err != nil {
			panic("redis setup: redis ping error")
		}

		fmt.Println(i, res)

		InstanceMap[i] = RedisClient
	}
}

func GetRedis(instance Instance) *redisGo.Client {
	if _, ok := InstanceMap[instance]; !ok {
		Setup()
	}

	client, _ := InstanceMap[instance]

	return client
}
