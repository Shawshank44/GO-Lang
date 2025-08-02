package main

import (
	"fmt"
	"time"
)

// RateLimiter limits the number of requests over time using a token bucket algorithm.
type RateLimiter struct {
	tokens     chan struct{} // Channel acts as a token bucket; each token allows one request
	refillTime time.Duration // Time interval between token refills
}

// NewRateLimiter creates and returns a new RateLimiter instance.
// rateLimit: max number of tokens (capacity of the bucket)
// refillTime: duration to wait before trying to add a new token
func NewRateLimiter(rateLimit int, refillTime time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens:     make(chan struct{}, rateLimit),
		refillTime: refillTime,
	}

	// Fill the token bucket initially with all available tokens
	for i := 0; i < rateLimit; i++ {
		rl.tokens <- struct{}{}
	}

	// Start a goroutine to refill tokens periodically
	go rl.startRefill()

	return rl
}

// startRefill adds tokens back into the bucket at fixed intervals.
// It uses a ticker to trigger refill attempts.
func (rl *RateLimiter) startRefill() {
	ticker := time.NewTicker(rl.refillTime)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Try to add a token if there's space in the bucket
			select {
			case rl.tokens <- struct{}{}:
			default:
				// Bucket is full; do nothing
			}
		}
	}
}

// allow returns true if a token is available (i.e., request is allowed).
// Otherwise, it returns false (i.e., request is denied).
func (rl *RateLimiter) allow() bool {
	select {
	case <-rl.tokens:
		return true // Token consumed, request allowed
	default:
		return false // No tokens, request denied
	}
}

// main simulates a series of requests using the rate limiter.
func main() {
	// Create a rate limiter that allows up to 5 requests and refills 1 token every second
	rateLimiter := NewRateLimiter(5, time.Second)

	// Simulate 10 requests, spaced 200 milliseconds apart
	for i := 0; i < 10; i++ {
		if rateLimiter.allow() {
			fmt.Printf("Request %d: Allowed\n", i+1)
		} else {
			fmt.Printf("Request %d: Denied\n", i+1)
		}
		time.Sleep(200 * time.Millisecond)
	}
}
