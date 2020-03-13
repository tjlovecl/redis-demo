package demo02

import (
	"sync"
	"fmt"
	"time"
	"github.com/garyburd/redigo/redis"
	"learn/redis/pool"
)

// 延迟队列 delay-queue

type UniqueId struct {
	id int
	mu sync.Mutex
}

var sUniqueId UniqueId

// 加入延迟队列
func AddDelayQueue(command string, score int64) string {
	id := getUniqueId()
	delayQueueName := getDelayQueueName()
	commandId := uniqueIdToCommandKey(id)
	fun := func(redisClient redis.Conn) (interface{}, error) {
		err := redisClient.Send("ZADD", delayQueueName, score, id)
		if err != nil {
			return nil, err
		}
		err = redisClient.Send("SET", commandId, command, "EX", 3600*24)
		if err != nil {
			return nil, err
		}

		err = redisClient.Flush()
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	_, err := pool.Execute(fun)
	if err != nil {
		panic(err)
	}
	return commandId
}

// 获取命令的key
func uniqueIdToCommandKey(uniqueId string) string {
	return fmt.Sprintf("command_id_%s", uniqueId)
}

func getDelayQueueName() string {
	return "delay_queue"
}

// 获取唯一的id
func getUniqueId() string {
	sUniqueId.mu.Lock()
	defer sUniqueId.mu.Unlock()
	sUniqueId.id = (sUniqueId.id + 1) % 1000
	return fmt.Sprintf("%d_%d", time.Now().Unix(), sUniqueId.id)
}


// 获取延迟队列的值，并把延迟队列的值放入普通的队列
func DelayQueueToQueue(score int64) string {
	delayQueueName := getDelayQueueName()
	fun := func(redisClient redis.Conn) (interface{}, error) {
		result, err := redisClient.Do("ZRANGEBYSCORE", delayQueueName, 0, score, "LIMIT", 0 ,1)
		if err != nil {
			return "", err
		}
		list := result.([]interface{})
		if len(list) == 0 {
			return "", nil
		}

		uniqueId := string(list[0].([]uint8))
		err = redisClient.Send("LPUSH", getQueueKey(), uniqueId)
		if err != nil {
			return "", err
		}

		err = redisClient.Send("ZREM", delayQueueName, uniqueId)
		if err != nil {
			return "", err
		}

		redisClient.Flush()
		if err != nil {
			return "", err
		}

		return uniqueId, err
	}
	uniqueId, err := pool.Execute(fun)
	if err != nil {
		panic(err)
	}
	return uniqueId.(string)
}

// 执行普通队列
func GetCommandFromQueue() string {
	queueKey := getQueueKey()
	fun := func(redisClient redis.Conn) (interface{}, error) {
		result, err := redisClient.Do("RPOP", queueKey)
		if err != nil {
			return "", err
		}

		if result == nil {
			return "", err
		}

		uniqueId := string(result.([]byte))
		commandId := uniqueIdToCommandKey(uniqueId)

		result, err = redisClient.Do("GET", commandId)
		if err != nil {
			return "", err
		}

		if result == nil {
			return "", err
		}

		redisClient.Do("DEL", commandId)
		return string(result.([]byte)), nil
	}

	result, err := pool.Execute(fun)
	if err != nil {
		panic(err)
	}

	return result.(string)

}


func getQueueKey() string{
	return "queue"
}


