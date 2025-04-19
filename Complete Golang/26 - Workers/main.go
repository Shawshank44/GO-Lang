package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Amount int
}

func worker(id int, jobs <-chan Order, wg *sync.WaitGroup, mu *sync.Mutex, revenue *int) {
	for order := range jobs {
		fmt.Printf("Worker %d : Processing order %d ($%d)\n", id, order.ID, order.Amount)
		// time.Sleep(time.Second)

		mu.Lock()
		*revenue += order.Amount
		mu.Unlock()

		fmt.Printf("Worker %d : Done with order %d \n", id, order.ID)
		wg.Done()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	jobs := make(chan Order, 20)

	var wg sync.WaitGroup
	var mu sync.Mutex
	totalRevenue := 0

	numworkers := 3

	for i := 0; i < numworkers; i++ {
		go worker(i, jobs, &wg, &mu, &totalRevenue)
	}

	for i := 1; i < 10; i++ {
		jobs <- Order{ID: i, Amount: rand.Intn(400) + 100}
		wg.Add(1)
	}

	close(jobs)
	wg.Wait()

	fmt.Println("Total revenue : ", totalRevenue)

}
