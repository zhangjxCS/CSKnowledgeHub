package main

import (
	"fmt"
	"sync"
	"time"
)

type Clock func() time.Time

type RateLimiter struct {
	mu             sync.Mutex
	capacity       int
	tokens         int
	refillAmount   int
	refillInterval time.Duration
	clock          Clock
	lastRefillTime time.Time
}

func NewRateLimiter(capacity int, refillAmount int, refillInterval time.Duration, clock Clock) (*RateLimiter, error) {
	if capacity <= 0 {
		return nil, fmt.Errorf("capacity must be a positive integer")
	}
	if refillAmount <= 0 {
		return nil, fmt.Errorf("refillAmount must be a positive integer")
	}
	if refillInterval <= 0 {
		return nil, fmt.Errorf("refillInterval must be a positive duration")
	}
	if clock == nil {
		return nil, fmt.Errorf("clock must not be nil")
	}

	return &RateLimiter{
		capacity:       capacity,
		tokens:         capacity,
		refillAmount:   refillAmount,
		refillInterval: refillInterval,
		clock:          clock,
		lastRefillTime: clock(),
	}, nil
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := rl.clock()
	elapsed := now.Sub(rl.lastRefillTime)

	if elapsed >= rl.refillInterval {
		refills := int(elapsed / rl.refillInterval)
		rl.tokens += refills * rl.refillAmount
		if rl.tokens > rl.capacity {
			rl.tokens = rl.capacity
		}
		rl.lastRefillTime = rl.lastRefillTime.Add(time.Duration(refills) * rl.refillInterval)
	}

	if rl.tokens > 0 {
		rl.tokens--
		return true
	}
	return false
}
