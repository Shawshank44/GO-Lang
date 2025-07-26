package main

import (
	"fmt"
	"sync"
	"time"
)

// func worker(id int, wg *sync.WaitGroup) { // Basic Example without using channels
// 	defer wg.Done()
// 	fmt.Printf("Worker %d starting \n", id)
// 	time.Sleep(time.Second)
// 	fmt.Printf("Worker %d finished the task \n", id)
// }

// func Worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) { // wait group with channels
// 	defer wg.Done()
// 	fmt.Printf("Worker %d starting \n", id)
// 	time.Sleep(time.Second)
// 	for task := range tasks {
// 		results <- task * 2
// 		fmt.Printf("Worker %d finished \n", id)
// 	}
// }

type Worker struct {
	ID   int
	Task string
}

// Perform a task will simulate worker:
func (w *Worker) PerformTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("WorkerID %d started %s \n", w.ID, w.Task)
	time.Sleep(time.Second)
	fmt.Printf("WorkerId %d finished %s \n", w.ID, w.Task)
}

func main() {
	// Basic Example without using channels :
	// var wg sync.WaitGroup
	// numWorkers := 3

	// wg.Add(numWorkers)

	// // Launch workers:
	// for i := range numWorkers {
	// 	go worker(i, &wg)
	// }

	// wg.Wait() // will block the thread
	// fmt.Println("All Workers finished")

	// Waitgroup with Channels :
	// var wg sync.WaitGroup
	// numworkers := 3
	// numJobs := 5
	// results := make(chan int, numJobs)
	// tasks := make(chan int, numJobs)

	// wg.Add(numworkers)

	// for i := range numworkers {
	// 	go Worker(i+1, tasks, results, &wg)
	// }

	// for i := range numJobs {
	// 	tasks <- i + 1
	// }
	// close(tasks)
	// go func() {
	// 	wg.Wait()
	// 	close(results)
	// }()

	// for result := range results {
	// 	fmt.Println("Result", result)
	// }

	// Performing a Task (Construction):
	var wg sync.WaitGroup
	task := []string{"Digging", "Laying bricks", "Painting"}

	for i, task := range task {
		worker := Worker{
			ID:   i + 1,
			Task: task,
		}
		wg.Add(1)
		go worker.PerformTask(&wg)
	}
	wg.Wait()

	// Construction is finished
	fmt.Println("Construction is finished")

}
