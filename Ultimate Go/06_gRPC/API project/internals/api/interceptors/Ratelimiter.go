package interceptors

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type RateLimiter struct {
	mu        sync.Mutex
	visitors  map[string]int
	limit     int
	resetTime time.Duration
}

func NewRateLimiter(limit int, resetTime time.Duration) *RateLimiter {
	rl := &RateLimiter{
		visitors:  make(map[string]int),
		limit:     limit,
		resetTime: resetTime,
	}
	go rl.resetVisitorCount() // will start the reset routine
	return rl
}

func (rl *RateLimiter) resetVisitorCount() {
	for {
		time.Sleep(rl.resetTime)
		rl.mu.Lock()
		rl.visitors = make(map[string]int)
		rl.mu.Unlock()
	}
}

func (rl *RateLimiter) RateLimiterInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Println("Ratelimiter Middleware being returned")
	rl.mu.Lock()
	defer rl.mu.Unlock()
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unable to get client IP.")
	}

	visitorIP := p.Addr.String() // you might want to extract the IP in a more sophisticated way
	rl.visitors[visitorIP]++
	log.Printf(`
	++++++++++++ Visitor count from IP ++++++++++++
		%s : %d
	`, visitorIP, rl.visitors[visitorIP])

	if rl.visitors[visitorIP] > rl.limit {
		return nil, status.Error(codes.ResourceExhausted, "Too many requests")
	}
	fmt.Println("Ratelimiter Middleware ends...")
	return handler(ctx, req)
}
