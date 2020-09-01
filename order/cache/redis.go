package cache

import (
	"github.com/gomodule/redigo/redis"
	"os"
)
// 获取键和删除键
var RedisConn *redis.Pool
//初始化redis连接池
func SetUp()  {
	redisAddr := os.Getenv("redis_addr")
	if redisAddr == "" {
		redisAddr = "192.168.93.10"
	}

	RedisConn = &redis.Pool{
		MaxIdle:   10000,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisAddr + ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func IsExist(key string) int {
	var isExist int
	conn := RedisConn.Get()
	defer conn.Close()
	isExist, _ = redis.Int(conn.Do("get", key))
	return isExist
}