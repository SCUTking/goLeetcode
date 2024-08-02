package limiter

import (
	"sync"
	"time"
)

// TokenBucketLimiter 令牌桶限流器
// 与漏桶算法的字段个数与含义一模一样
type TokenBucketLimiter struct {
	capacity      int        // 容量
	currentTokens int        // 令牌数量
	rate          int        // 发放令牌速率/秒
	lastTime      time.Time  // 上次发放令牌时间
	mutex         sync.Mutex // 避免并发问题
}

func NewTokenBucketLimiter(capacity, rate int) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		capacity: capacity,
		rate:     rate,
		lastTime: time.Now(),
	}
}

// TryAcquire 请求是从里面取令牌，后台从里面生产令牌，与漏桶算法是相反的。
func (l *TokenBucketLimiter) TryAcquire() bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// 尝试发放令牌
	now := time.Now()
	// 距离上次发放令牌的时间
	//计算他们的差值
	interval := now.Sub(l.lastTime)
	//如果大于一秒，才进行生产令牌，在这次请求后面的一秒内都不会生产令牌；
	//因为小于一秒生产的话，会导致令牌的数量每次请求都会增加，达不到限流的目的。
	if interval >= time.Second {

		// 当前令牌数量+距离上次发放令牌的时间(秒)*发放令牌速率
		//取两者的最小值
		l.currentTokens = minInt(l.capacity, l.currentTokens+int(interval/time.Second)*l.rate)
		l.lastTime = now
	}

	// 如果没有令牌，请求失败
	if l.currentTokens == 0 {
		return false
	}
	// 如果有令牌，当前令牌-1，请求成功
	l.currentTokens--
	return true
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
