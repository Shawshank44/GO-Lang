package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	pid := os.Getpid()

	fmt.Println(pid)

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// Notify channel on interrupt or terminate signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP) // for SIGTERM command is (taskkill /PID {Generated pid} /F)
	// In windows we cannot directly use the SIGTERM and SIGINT command.
	go func() {
		sig := <-sigs
		fmt.Println("Received signal", sig)
		done <- true
	}()
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Stopping work due to signal")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(time.Second)
			}
		}
		// sig := <-sigs
		// switch sig {
		// case syscall.SIGINT:
		// 	fmt.Println("Received SIGINT (INTERRUPT)")
		// case syscall.SIGTERM:
		// 	fmt.Println("Received SIGTERM (TERMINATE)")
		// case syscall.SIGHUP:
		// 	fmt.Println("Received SIGHUP (HUNGUP)")
		// }
		// fmt.Println("Graceful Exit")
		// os.Exit(0)
	}()

	// Simulate some work
	// fmt.Println("Working...")
	for {
		time.Sleep(time.Second)
	}
}
