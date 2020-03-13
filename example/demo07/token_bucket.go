package demo07

import (
	"time"
	"sync"
)

// 令牌桶算法(Token Bucket)
// 步骤如下
/*
   1. 通过计算时间差，发放令牌
   2. 根据访问量 从池子中取出相应的令牌， 返回成功取得令牌的数量和未取得令牌的数量
 */
type TokenBucket struct {
	Capacity              int64 // 桶的总大小
	LeaksIntervalInMillis int64 // 每消耗一个，需要多少毫秒
	Used                  int64 // 已使用的大小
	LastLeakTimestamp     int64 // 最后更新时间
	mu sync.Mutex // 锁
}

func (self *TokenBucket) Init(capacity int64, leaksIntervalInMillis int64) {
	self.Capacity = capacity
	self.LeaksIntervalInMillis = leaksIntervalInMillis
	self.Used = 0
	self.LastLeakTimestamp = time.Now().UnixNano() / 1e6
}


func (self *TokenBucket) TryConsume(drop int64) (succNum, failNum int64){
	self.mu.Lock()
	defer self.mu.Unlock()
	self.leak()
	if (self.Used + drop <= self.Capacity) {
		self.Used = self.Used + drop
		succNum = drop
		failNum = 0
	} else {
		succNum = self.Capacity - self.Used
		failNum = self.Used + drop - self.Capacity
		self.Used = self.Capacity
	}
	return succNum, failNum
}


func (self *TokenBucket) leak() {
	now := 	time.Now().UnixNano() / 1e6
	if now > self.LastLeakTimestamp {
		millisSinceLastLeak := now - self.LastLeakTimestamp
		leaks := millisSinceLastLeak / self.LeaksIntervalInMillis
		if (leaks > 0) {
			if (self.Used <= leaks) {
				self.Used = 0
			} else {
				self.Used = self.Used - leaks
			}
			self.LastLeakTimestamp = leaks * self.LeaksIntervalInMillis + self.LastLeakTimestamp
		}
	}
}