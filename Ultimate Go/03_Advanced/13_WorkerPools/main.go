package main

import (
	"fmt"
	"time"
)

// func worker(id int, tasks <-chan int, results chan<- int) { // Basic worker function
// 	for task := range tasks {
// 		fmt.Printf("Worker %d processing task %d \n", id, task)
// 		// Simulate work
// 		time.Sleep(time.Second)
// 		results <- task * 2
// 	}
// }

// Ticket simulation :
type TicketRequest struct {
	PersonID  int
	numTicket int
	cost      int
}

func TicketProcessor(requests <-chan TicketRequest, results chan<- int) {
	for req := range requests {
		fmt.Printf("Processing %d ticket(s) of personID %d with total cost %d\n", req.numTicket, req.PersonID, req.cost)
		// Simulate processing time
		time.Sleep(time.Second)
		results <- req.PersonID
	}
}

func main() {
	// Basic Worker pool pattern :
	// numWorkers := 3
	// numJobs := 10
	// tasks := make(chan int, numJobs)
	// results := make(chan int, numJobs)

	// // Creating workers
	// for i := range numWorkers {
	// 	go func() {
	// 		worker(i, tasks, results)
	// 	}()
	// }

	// // Send values to the tasks channel
	// for i := range numJobs {
	// 	tasks <- i
	// }

	// close(tasks)

	// // Collect the results
	// for range numJobs {
	// 	result := <-results
	// 	fmt.Println("Result : ", result)
	// }

	// Simulating ticket registration :
	numRequests := 5
	price := 5
	ticketRequests := make(chan TicketRequest, numRequests)
	ticketResults := make(chan int)

	// start ticket processor/worker
	for range 3 { // number of workers
		go TicketProcessor(ticketRequests, ticketResults)
	}

	// // send ticket requests :
	for i := range numRequests {
		ticketRequests <- TicketRequest{PersonID: i + 1, numTicket: (i + 1) * 2, cost: (i + 1) * price}
	}
	close(ticketRequests)

	for range numRequests {
		fmt.Printf("Ticket for personId %d processed sucessfully! \n", <-ticketResults)
	}
}
