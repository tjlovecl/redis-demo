package demo03

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"learn/redis/pool"
)

// 位图操作

// 模拟登陆
func Login(uid int, day string) bool {
	key := getKey(day)
	fun := func (conn redis.Conn) (interface{}, error){
		return conn.Do("SETBIT", key, uid, 1)
	}

	_, err := pool.Execute(fun)
	if err != nil {
		panic(err)
	}

	return true
}

// 判定用户是否登陆过
func IsLogin(uid int, day string) bool {
	key := getKey(day)
	fun := func (conn redis.Conn) (interface{}, error){
		return conn.Do("GETBIT", key, uid)
	}

	result, err := pool.Execute(fun)
	if err != nil {
		panic(err)
	}
	return result.(int64) == 1
}

// 获得当天的用户登陆总数
func GetDayLoginNum(day string) int64 {
	key := getKey(day)
	fun := func (conn redis.Conn) (interface{}, error){
		return conn.Do("BITCOUNT", key)
	}
	result, err := pool.Execute(fun)
	if err != nil {
		panic(err)
	}
	return result.(int64)
}

// 获取连续n天的用户登陆总数
func GetLoginNumByDays(days []string) int64 {
	if len(days) == 0 {
		return 0
	}

	destKey := "dest_login_count"
	keys := []interface{}{
		"AND", getKey(destKey),
	}

	for _, day := range days {
		keys = append(keys, getKey(day))
	}

	fun := func (conn redis.Conn) (interface{}, error){
		_, err := conn.Do("BITOP", keys...)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	_, err := pool.Execute(fun)
	if err != nil {
		panic(err)
	}

	return GetDayLoginNum(destKey)
}



// 获取
func getKey(day string) string {
	return fmt.Sprintf("bit_%s", day)
}