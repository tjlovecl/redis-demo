package demo05

import (
	"github.com/garyburd/redigo/redis"
	"learn/redis/pool"
)

// 布隆过滤器
// 安装

/*
wget https://github.com/RedisBloom/RedisBloom/archive/master.zip
unzip master.zip
cd RedisBloom-master/
make
mv redisbloom.so ../
/root/redis-5.0.5/src/redis-server  /root/redis-5.0.5/redis.conf --loadmodule /root/redis-5.0.5/modules/redisbloom.so
 */


// 添加数据
func BatchAdd(key string, values []interface{}) bool {
	fun := func (conn redis.Conn) (interface{}, error) {
		params := []interface{}{key}
		params = append(params, values...)
		return conn.Do("BF.MADD", params...)
	}

	_, err := pool.Execute(fun)
	if err != nil {
		panic(err)
	}
	return true
}

// 验证是否存在
func BatchExists(key string, values []interface{}) []interface{} {
	fun := func (conn redis.Conn) (interface{}, error) {
		params := []interface{}{key}
		params = append(params, values...)
		return conn.Do("BF.MEXISTS", params...)
	}

	result, err := pool.Execute(fun)
	if err != nil {
		panic(err)
	}

	return result.([]interface{})
}
