package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		close(quit)
	}()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Tick")
		case <-quit:
			fmt.Println("Quitting.....")
			return
		}
	}

}
