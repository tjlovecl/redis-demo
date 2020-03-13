package demo01

import (
	"math/rand"
	"time"
	"learn/redis/pool"
)

// 分布式锁
func DistributedLock() (interface{}, error) {
	redisClient := pool.GetRedisClient()
	defer redisClient.Close()
	key := getKey()
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(1000)
	lockTime := getLockTime()

	result, err :=  redisClient.Do("set", key, randNum, "EX", lockTime, "NX")

	if err != nil {
		return false, err
	}
	if result == nil {
		return false, err
	}

	return true, err
}

// 获取key
func getKey() string{
	return "1234"
}

// 获取锁的时间
func getLockTime() int{
	return 5
}

