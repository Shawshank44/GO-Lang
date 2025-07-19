package main

func main() {
	// ch1 := make(chan int)
	// ch2 := make(chan int)

	// go func() {
	// 	time.Sleep(time.Second)
	// 	ch1 <- 100
	// 	ch1 <- 111
	// }()

	// go func() {
	// 	// time.Sleep(time.Second)
	// 	ch2 <- 200
	// }()

	// time.Sleep(2 * time.Second)
	// for range 3 { // loop if we want both the values
	// 	select {
	// 	case msg1 := <-ch1:
	// 		fmt.Println("Received from ch1 ", msg1)
	// 	case msg2 := <-ch2:
	// 		fmt.Println("Received from ch2 ", msg2)
	// 		// default:
	// 		// 	fmt.Println("No channels ready")
	// 	}
	// }

	// fmt.Println("End of the program")

	// Channel timeout functionality :
	// ch := make(chan int)

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	ch <- 1
	// 	close(ch)
	// }()

	// select {
	// case msg := <-ch:
	// 	fmt.Println("Received : ", msg)
	// case <-time.After(3 * time.Second):
	// 	fmt.Println("Time out")
	// }

	// error handeling in channel:
	// ch := make(chan int)

	// go func() {
	// 	ch <- 1
	// 	close(ch)
	// }()

	/*
		In Go, when receiving a value from a channel using the form value, ok := <-channel, the ok is a boolean that indicates whether the channel is open (true) or closed (false).
		This helps the program know if a value was successfully received or if the channel is closed and no more values will be sent.
	*/

	// for {
	// 	select {
	// 	case msg, ok := <-ch:
	// 		if !ok {
	// 			fmt.Println("Channel closed")
	// 			// clean up :
	// 			return
	// 		}
	// 		fmt.Println("Received:", msg)
	// 	}
	// }
}
