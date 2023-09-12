package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type bucket struct {
	maxToken, token int
	interval        time.Duration
	lastUsed        time.Time
}

func newBucket(limit int, interval time.Duration) *bucket {
	bucket := &bucket{
		maxToken: limit,
		token:    limit,
		interval: interval,
		lastUsed: time.Now(),
	}
	go bucket.refill()
	return bucket
}

func (b *bucket) take() {
	b.lastUsed = time.Now()
	b.token--
}

func (b *bucket) allow() bool {
	return b.token > 0
}

func (b *bucket) refill() bool {
	for {
		now := time.Now()
		if now.Sub(b.lastUsed) > b.interval {
			b.token = b.maxToken
		}
	}
}

type rateLimiter struct {
	limit    int
	interval time.Duration
	ipMap    sync.Map
}

func NewLimiter(limit int, interval time.Duration) *rateLimiter {
	return &rateLimiter{
		limit:    limit,
		interval: interval,
	}
}

func (rl *rateLimiter) RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		value, _ := rl.ipMap.LoadOrStore(ip, newBucket(rl.limit, rl.interval))

		data, ok := value.(*bucket)
		if !ok {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if !data.allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
			})
			return
		}
		data.take()
		c.Next()
	}
}
