package main

import (
	"fmt"
	"sync"
	"time"
)

// RateLimiter implements a fixed-window rate limiting mechanism
type RateLimiter struct {
	mu        sync.Mutex    // Ensures safe concurrent access
	count     int           // Current request count in the window
	limit     int           // Max requests allowed per window
	window    time.Duration // Duration of the rate-limiting window
	resetTime time.Time     // When the window resets
}

// Constructor for RateLimiter
func NewRateLimter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:  limit,
		window: window,
	}
}

// Allow returns true if a request is allowed, false if it exceeds the limit
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()

	// If current time is past the reset time, reset the window and count
	if now.After(rl.resetTime) {
		rl.resetTime = now.Add(rl.window)
		rl.count = 0
	}

	// Allow the request if under the limit
	if rl.count < rl.limit {
		rl.count++
		return true
	}

	// Deny if limit is reached
	return false
}

func main() {
	var wg sync.WaitGroup

	// Create a rate limiter allowing 3 requests per second
	rateLimiter := NewRateLimter(3, 1*time.Second)

	// Launch 10 concurrent requests
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			if rateLimiter.Allow() {
				fmt.Printf("Request #%d: Allowed\n", id)
			} else {
				fmt.Printf("Request #%d: Denied\n", id)
			}
		}(i)

		// Optional: add delay to test over time instead of all at once
		// time.Sleep(200 * time.Millisecond)
	}

	wg.Wait()
}
