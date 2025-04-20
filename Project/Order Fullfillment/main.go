package main

import (
	"context"
	"fmt"
	"math/rand"
	models "myapp/Models"
	service "myapp/Service"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create the worker pool
	pool := service.NewWorkerPool(4, 20)
	pool.Start(ctx)

	customers := []string{"Alice", "Bob", "Charlie", "Diana"}
	items := []string{"Laptop", "Phone", "Shoes", "Book"}

	for i := 1; i <= 12; i++ {
		order := models.Order{
			ID:       i,
			Customer: customers[rand.Intn(len(customers))],
			Item:     items[rand.Intn(len(items))],
			Quantity: rand.Intn(3) + 1,
		}

		ok := pool.TrySubmit(func() {
			order.Fulfill()
		})

		if !ok {
			fmt.Printf("⚠️ Order #%d dropped: queue full\n", order.ID)
		}
	}

	time.Sleep(10 * time.Second)
	pool.StopWithTimeout(3 * time.Second)

	fmt.Println("🎉 All available orders processed. Exiting.")
}
