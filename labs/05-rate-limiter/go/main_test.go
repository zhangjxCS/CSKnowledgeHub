package main

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type fakeClock struct {
	mu  sync.Mutex
	now time.Time
}

func (c *fakeClock) Now() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.now
}

func (c *fakeClock) Advance(duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.now = c.now.Add(duration)
}

func TestInitialCapacityAndExhaustion(t *testing.T) {
	clock := &fakeClock{}
	limiter, err := NewRateLimiter(3, 1, 10*time.Second, clock.Now)
	if err != nil {
		t.Fatalf("NewRateLimiter() error = %v", err)
	}

	for i := 0; i < 3; i++ {
		if !limiter.Allow() {
			t.Fatalf("Allow() call %d = false, want true", i+1)
		}
	}
	if limiter.Allow() {
		t.Fatal("Allow() after exhaustion = true, want false")
	}
}

func TestRefillAtExactBoundary(t *testing.T) {
	clock := &fakeClock{}
	limiter, _ := NewRateLimiter(1, 1, 10*time.Second, clock.Now)
	limiter.Allow()

	clock.Advance(9 * time.Second)
	if limiter.Allow() {
		t.Fatal("Allow() before boundary = true, want false")
	}

	clock.Advance(time.Second)
	if !limiter.Allow() {
		t.Fatal("Allow() at boundary = false, want true")
	}
}

func TestPartialRefillIntervalIsPreserved(t *testing.T) {
	clock := &fakeClock{}
	limiter, _ := NewRateLimiter(1, 1, 10*time.Second, clock.Now)
	limiter.Allow()

	clock.Advance(15 * time.Second)
	if !limiter.Allow() {
		t.Fatal("Allow() after one interval = false, want true")
	}

	clock.Advance(5 * time.Second)
	if !limiter.Allow() {
		t.Fatal("Allow() after completing partial interval = false, want true")
	}
}

func TestInvalidConfiguration(t *testing.T) {
	clock := &fakeClock{}
	tests := []struct {
		name           string
		capacity       int
		refillAmount   int
		refillInterval time.Duration
		clock          Clock
	}{
		{name: "capacity", capacity: 0, refillAmount: 1, refillInterval: time.Second, clock: clock.Now},
		{name: "refill amount", capacity: 1, refillAmount: 0, refillInterval: time.Second, clock: clock.Now},
		{name: "refill interval", capacity: 1, refillAmount: 1, refillInterval: 0, clock: clock.Now},
		{name: "clock", capacity: 1, refillAmount: 1, refillInterval: time.Second, clock: nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, err := NewRateLimiter(test.capacity, test.refillAmount, test.refillInterval, test.clock); err == nil {
				t.Fatal("NewRateLimiter() error = nil, want validation error")
			}
		})
	}
}

func TestConcurrentCallersDoNotShareTokens(t *testing.T) {
	const (
		capacity    = 5
		callerCount = 20
	)

	clock := &fakeClock{}
	limiter, _ := NewRateLimiter(capacity, 1, time.Second, clock.Now)
	start := make(chan struct{})
	var successes atomic.Int32
	var workers sync.WaitGroup

	workers.Add(callerCount)
	for i := 0; i < callerCount; i++ {
		go func() {
			defer workers.Done()
			<-start
			if limiter.Allow() {
				successes.Add(1)
			}
		}()
	}

	close(start)
	workers.Wait()

	if got := successes.Load(); got != capacity {
		t.Fatalf("successful calls = %d, want %d", got, capacity)
	}
}
