package main

import (
	"fmt"
	"sync"
	"time"
)

// LeakyBucket defines the rate limiter structure using the leaky bucket algorithm
type LeakyBucket struct {
	capacity int           // Max number of tokens the bucket can hold
	leakRate time.Duration // Time interval between leaking (replenishing) one token
	tokens   int           // Current number of available tokens
	lastLeak time.Time     // Last time tokens were leaked
	mu       sync.Mutex    // Mutex to ensure thread safety
}

// NewLeakyBucket initializes a new LeakyBucket with a given capacity and leak rate
func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		tokens:   capacity,   // Start full
		lastLeak: time.Now(), // Initialize last leak time to now
	}
}

// Allow checks if a request can be allowed or not, based on the current tokens
func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()         // Lock for concurrent access
	defer lb.mu.Unlock() // Unlock when function exits

	now := time.Now()
	elapsedTime := now.Sub(lb.lastLeak)

	// Calculate how many tokens to add based on elapsed time
	tokensToAdd := int(elapsedTime / lb.leakRate)
	lb.tokens += tokensToAdd

	// Cap tokens at bucket capacity
	if lb.tokens > lb.capacity {
		lb.tokens = lb.capacity
	}

	// Advance the lastLeak time by how much time we added tokens for
	lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensToAdd) * lb.leakRate)
	fmt.Printf("Token added %d, Tokens subtracted %d, Total tokens %d \n", tokensToAdd, 1, lb.tokens)
	fmt.Printf("Last leak time : %v \n", lb.lastLeak)
	// If at least one token is available, allow the request and consume one token
	if lb.tokens > 0 {
		lb.tokens--
		return true
	}

	// Otherwise, reject the request
	return false
}

func main() {
	var wg sync.WaitGroup

	LB := NewLeakyBucket(5, 500*time.Millisecond) // 5 tokens max, leak (replenish) 1 token every 500ms

	for range 10 { // Attempt 10 concurrent requests
		wg.Add(1)
		go func() {
			defer wg.Done()
			if LB.Allow() {
				fmt.Println("Current time : ", time.Now())
				fmt.Println("Request Accepted")
			} else {
				fmt.Println("Request Denied")
			}
			// Optional delay to simulate processing time
			// time.Sleep(200 * time.Millisecond)
		}()
	}
	wg.Wait() // Wait for all goroutines to finish
}
