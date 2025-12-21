package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// func checkEvenOdd(ctx context.Context, num int) string {
// 	select {
// 	case <-ctx.Done():
// 		return "operation cancelled"
// 	default:
// 		if num%2 == 0 {
// 			return fmt.Sprintf("%d is even", num)
// 		} else {
// 			return fmt.Sprintf("%d is odd", num)
// 		}
// 	}
// }

func DoWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work Cancelled", ctx.Err())
			return
		default:
			fmt.Println("Working...")

		}
		time.Sleep(500 * time.Millisecond)
	}
}

// Contextual logging function :
func logWithContext(ctx context.Context, message string) {
	requestID := ctx.Value("request ID")
	log.Printf("RequestID : %v - %v", requestID, message)
}

func main() {
	// **** Difference between Context.TODO and Context.Background
	// todoContext := context.TODO()
	// contextBKG := context.Background()

	// ctx := context.WithValue(todoContext, "name", "John")
	// fmt.Println(ctx)
	// fmt.Println(ctx.Value("name"))

	// ctx1 := context.WithValue(contextBKG, "city", "New york")
	// fmt.Println(ctx1)
	// fmt.Println(ctx1.Value("city"))

	// ****
	// ctx := context.TODO()
	// result := checkEvenOdd(ctx, 5)
	// fmt.Println("Result with context.TODO()", result)

	// ctx = context.Background()
	// ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	// defer cancel()

	// result = checkEvenOdd(ctx, 10)
	// fmt.Println("Result from timeout context", result)

	// time.Sleep(2 * time.Second)
	// result = checkEvenOdd(ctx, 15)
	// fmt.Println("Result after timeout : ", result)

	// Cancellation example with timout and manual
	ctx := context.Background()
	// with time out:
	// ctx, cancel := context.WithTimeout(ctx, 2*time.Second) // timer of the context starts here. You have this specified time duration to use this context, after this time duration, the context will send a cancelation signal
	// defer cancel()

	// With manual
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		time.Sleep(2 * time.Second) // simulating a heavy task. Consider this a heavy time consuming operation
		cancel()                    // manually cancelling only after the task is finished
	}()

	ctx = context.WithValue(ctx, "name", "John")

	go DoWork(ctx)

	time.Sleep(3 * time.Second)

	requestID := ctx.Value("name")
	if requestID != nil {
		fmt.Println("Request ID : ", requestID)
	} else {
		fmt.Println("No request ID found")
	}

	logWithContext(ctx, "This is just a log message") // contextual logging
}
