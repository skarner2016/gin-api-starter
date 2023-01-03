package redis

type Instance string

type RedisConf struct {
	Host string
	Port int64
	Pass string
	DB   int
}

const InstanceDefault Instance = "default"
