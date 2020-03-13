package demo06

import (
	"time"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"math/rand"
	"learn/redis/pool"
)

// 简单限流 1分钟内 只能访问5次

const
(
	INTERVAL = 10
	TIMES    = 3
	KEY = "SIMPLE_LIMIT_KEY"
)

func Limit() bool {
	now := time.Now().Unix()
	start := now - INTERVAL


	fun := func (conn redis.Conn) (interface{}, error) {
		// 新增
		conn.Do("ZADD", KEY, now, fmt.Sprintf("%d_%d", now, rand.Intn(10000)))

		// 查看数量
		result, err := conn.Do("ZRANGEBYSCORE", KEY, start, now)
		if err != nil {
			panic(err)
		}

		conn.Do("ZREMRANGEBYSCORE", KEY, 0, start)
		conn.Do("EXPIRE", KEY, INTERVAL + 2)

		return result, nil
	}

	result, err := pool.Execute(fun)
	if err != nil {
		panic(err)
	}

	list := result.([]interface{})
	if len(list) <= TIMES {
		return true
	} else {
		return false
	}
}
