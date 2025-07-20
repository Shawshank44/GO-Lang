package main

import (
	"fmt"
	"time"
)

func PeriodicTask() { // scheduling logging for periodic updates
	fmt.Println("Performing Periods task at : ", time.Now())
}

func main() {
	// Basic ticker
	// ticker := time.NewTicker(time.Second)
	// defer ticker.Stop()
	// // for tick := range ticker.C {
	// // 	fmt.Println("Tick at : ", tick)
	// // }

	// i := 0
	// for range 5 {
	// 	i++
	// 	fmt.Println(i)
	// }

	// scheduling logging for periodic updates :
	// ticker := time.NewTicker(time.Second)
	// defer ticker.Stop()
	// for {
	// 	select {
	// 	case <-ticker.C:
	// 		PeriodicTask()
	// 	}
	// }

	// ticker := time.NewTicker(time.Second)
	// stop := time.After(5 * time.Second)
	// defer ticker.Stop()

	// for {
	// 	select {
	// 	case tick := <-ticker.C:
	// 		fmt.Println("Tick at : ", tick)
	// 	case <-stop:
	// 		fmt.Println("Stopping ticker.")
	// 		return
	// 	}
	// }

	// Multiple ticker :
	ticker1 := time.NewTicker(time.Second)
	ticker2 := time.NewTicker(time.Second)
	stop := time.After(10 * time.Second)

	defer ticker1.Stop()
	defer ticker2.Stop()

	done := make(chan struct{})

	// Goroutine for ticker1
	go func() {
		for {
			select {
			case tick := <-ticker1.C:
				fmt.Println("Tick:", tick)
			case <-done:
				return
			}
		}
	}()

	// Goroutine for ticker2
	go func() {
		i := 0
		for {
			select {
			case <-ticker2.C:
				i++
				fmt.Println(i)
			case <-done:
				return
			}
		}
	}()

	// Stop after 10 seconds
	<-stop
	close(done)
	fmt.Println("Multi ticker stopped")

}
