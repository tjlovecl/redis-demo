package pool

import (
	"github.com/garyburd/redigo/redis"
	"os"
)

var redisPool *redis.Pool

func getPool() *redis.Pool {
	if redisPool == nil {
		redisPool = &redis.Pool{
			MaxIdle:     10,
			MaxActive:   100,
			IdleTimeout: 60,
			Wait:        true,
			Dial: func() (redis.Conn, error) {
				return redis.DialURL(os.Getenv("REDIS_URL"))
			},
		}
	}
	return redisPool
}

//获取redis连接
func GetRedisClient() redis.Conn {
	redisClient := getPool().Get()
	if redisClient.Err() != nil {
		panic(redisClient.Err())
	}
	return redisClient
}

// 执行redis
func Execute(fun func(redis.Conn) (interface{}, error)) (interface{}, error) {
	redisClient := getPool().Get()
	defer redisClient.Close()
	return fun(redisClient)
}
