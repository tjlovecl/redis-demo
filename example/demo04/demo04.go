package demo04

import (
	"github.com/garyburd/redigo/redis"
	"learn/redis/pool"
)

// 批量插入
func BatchAdd(start, end int) bool {
	key := getKey()
	fun := func (conn redis.Conn) (interface{}, error){
		for i:= start;i<=end; i++ {
			conn.Do("PFADD", key, i)
		}
		return nil, nil
	}
	pool.Execute(fun)
	return true
}


// 获取统计数据
func Count() int64 {
	key := getKey()
	fun := func (conn redis.Conn) (interface{}, error){
		return conn.Do("PFCOUNT", key)
	}
	count, err := pool.Execute(fun)
	if err != nil {
		panic(err)
	}
	return count.(int64)
}


func getKey() string {
	return "hyper_key"
}