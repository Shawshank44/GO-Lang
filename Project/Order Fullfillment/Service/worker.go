package service

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Task defines a unit of work to be processed by the worker pool.
type Task func()

// WorkerPool manages a pool of worker goroutines to process tasks concurrently.
type WorkerPool struct {
	tasks      chan Task
	wg         sync.WaitGroup
	shutdownCh chan struct{}
	numWorkers int
	running    int32
}

// NewWorkerPool creates a new WorkerPool with the specified number of workers and task queue size.
func NewWorkerPool(numWorkers int, taskQueueSize int) *WorkerPool {
	return &WorkerPool{
		tasks:      make(chan Task, taskQueueSize),
		shutdownCh: make(chan struct{}),
		numWorkers: numWorkers,
	}
}

// Start begins worker execution in the pool.
// It is safe to call Start only once; subsequent calls will have no effect.
func (wp *WorkerPool) Start(ctx context.Context) {
	if !atomic.CompareAndSwapInt32(&wp.running, 0, 1) {
		return
	}

	for i := 0; i < wp.numWorkers; i++ {
		wp.wg.Add(1)
		go func(workerID int) {
			defer wp.wg.Done()
			for {
				select {
				case task := <-wp.tasks:
					func() {
						defer func() {
							if r := recover(); r != nil {
								fmt.Printf("Worker %d recovered from panic: %v\n", workerID, r)
							}
						}()
						task()
					}()
				case <-wp.shutdownCh:
					return
				case <-ctx.Done():
					return
				}
			}
		}(i)
	}
}

// Submit sends a task to the worker pool. It blocks if the queue is full.
// If the pool is not running, the task is rejected.
func (wp *WorkerPool) Submit(task Task) {
	if !wp.Running() {
		return
	}
	wp.tasks <- task
}

// TrySubmit tries to send a task without blocking.
// Returns false if the queue is full or the pool is stopped.
func (wp *WorkerPool) TrySubmit(task Task) bool {
	if !wp.Running() {
		return false
	}
	select {
	case wp.tasks <- task:
		return true
	default:
		return false
	}
}

// SubmitWithTimeout tries to send a task within a timeout duration.
// Returns false if the timeout is reached or the pool is stopped.
func (wp *WorkerPool) SubmitWithTimeout(task Task, timeout time.Duration) bool {
	if !wp.Running() {
		return false
	}
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case wp.tasks <- task:
		return true
	case <-timer.C:
		return false
	}
}

// Stop gracefully shuts down the worker pool and waits for all workers to finish.
func (wp *WorkerPool) Stop() {
	if !atomic.CompareAndSwapInt32(&wp.running, 1, 0) {
		return
	}
	close(wp.shutdownCh)
	wp.wg.Wait()
}

// StopWithTimeout attempts to shut down the pool within the given timeout.
// Logs a message if shutdown takes too long.
func (wp *WorkerPool) StopWithTimeout(timeout time.Duration) {
	if !atomic.CompareAndSwapInt32(&wp.running, 1, 0) {
		return
	}
	close(wp.shutdownCh)

	done := make(chan struct{})
	go func() {
		wp.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(timeout):
		fmt.Println("Worker pool shutdown timed out")
	}
}

// Running returns true if the worker pool is currently active.
func (wp *WorkerPool) Running() bool {
	return atomic.LoadInt32(&wp.running) == 1
}
