package limiter

import (
	"sync"
	"time"
)

// LeakyBucketLimiter 漏桶限流器
type LeakyBucketLimiter struct {
	peakLevel       int        // 最高水位
	currentLevel    int        // 当前水位 记录当前水位
	currentVelocity int        // 水流速度/秒   以秒为单位
	lastTime        time.Time  // 上次放水时间
	mutex           sync.Mutex // 避免并发问题
}

// NewLeakyBucketLimiter 漏桶的放水速度是定四的，不能应对突发流量
func NewLeakyBucketLimiter(peakLevel, currentVelocity int) *LeakyBucketLimiter {
	return &LeakyBucketLimiter{
		peakLevel:       peakLevel,
		currentVelocity: currentVelocity,
		lastTime:        time.Now(),
	}
}

// TryAcquire 最重要的方法，请求到来时，是否对该请求放行  每次有请求到来时就调用这个方法即可
// 请求是往里面塞，放水是由时间放
func (l *LeakyBucketLimiter) TryAcquire() bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// 尝试放水
	now := time.Now()
	// 距离上次放水的时间
	interval := now.Sub(l.lastTime)
	if interval >= time.Second {
		//水位减去在这期间放的水
		// 当前水位-距离上次放水的时间(秒)*水流速度
		//减去水位后，可能为负数，要与0进行比较
		l.currentLevel = maxInt(0, l.currentLevel-int(interval/time.Second)*l.currentVelocity)
		//设置最后一次放水，用于后续请求到来时进行判断
		l.lastTime = now
	}

	// 若到达最高水位，请求失败
	if l.currentLevel >= l.peakLevel {
		return false
	}
	// 若没有到达最高水位，当前水位+1，请求成功
	l.currentLevel++
	return true
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
