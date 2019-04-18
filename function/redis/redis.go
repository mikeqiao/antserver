package redis

import r "github.com/mikeqiao/common/redis"

func Init() {
	r.InitRedis()
}

func Close() {
	r.OnClose()
}
