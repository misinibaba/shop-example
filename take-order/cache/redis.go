package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/hashicorp/go-uuid"
	"os"
	"time"
)
// 获取键和删除键
const LuaScript = `
        if redis.call('get',KEYS[1]) == ARGV[1] then
        	return redis.call('del',KEYS[1]) else return 0 end
`
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

// 解锁
func RedisUnlock(key string, value string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	lua := redis.NewScript(1, LuaScript)
	result, err := redis.Int(lua.Do(conn, key, value))
	if err != nil {
		return false
	}
	return result != 0
}

// 分布式锁
func RedisLock() string {
	conn := RedisConn.Get()
	defer conn.Close()

	// 需要原子设置key并且设置过期时间，否则宕机后没有过期时间就完蛋了
	// value需要一个唯一的值，因为如果阻塞后回来（此时锁已经过期，另一个进程上了锁）释放锁，结果释放掉了别人的锁；所以需要验证一下value是不是当时设置的那个唯一value再释放锁
	// 这种分布式锁的话也会有问题，就是要是master挂掉，但锁还没来得及同步到slave,此时这个slave变成master，就可以让另一个客户端也拿到锁
	redisKey := "order:" + "1"
	redisValue, _ := uuid.GenerateUUID()
	for {
		if reply, _ := conn.Do("SET", redisKey, redisValue, "NX", "EX", 60); reply != nil {
			break
		}
		fmt.Println("sleep")
		time.Sleep(time.Duration(2) * time.Second)
	}
	return redisValue
}