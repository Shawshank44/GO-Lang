package main

// send data
// func processnum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("Processing number", num)
// 		time.Sleep(time.Second)
// 	}
// }

// receive data :
// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult
// }

// func Task(done chan bool) {
// 	defer func() {
// 		done <- true
// 	}()
// 	fmt.Println("Processing...")

// }

// func emailSender(emailChan chan string, done chan bool) { // you can also make channels type safe ex: emailChan <-chan string, done chan<- bool
// 	defer func() { done <- true }()
// 	for email := range emailChan {
// 		fmt.Println("Sending email to", email)
// 		time.Sleep(time.Second)
// 	}
// }

func main() {
	// blocking channels :
	// messageChan := make(chan string) // creating a channel

	// messageChan <- "ping" // channels are blocking

	// msg := <-messageChan

	// fmt.Println(msg)

	// sending the data :
	// numChan := make(chan int)

	// go processnum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// recieving data
	// result := make(chan int)

	// go sum(result, 4, 5)

	// res := <-result
	// fmt.Println(res)

	// done := make(chan bool)
	// go Task(done)
	// <-done

	// Buffered channels :
	// emailChan := make(chan string, 100)
	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i := 0; i < 5; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }
	// fmt.Println("Done sending")
	// close(emailChan) // closing channels
	// <-done

	// Multiple channels :
	// chan1 := make(chan int)
	// chan2 := make(chan string)
	// go func() {
	// 	chan1 <- 10
	// }()
	// go func() {
	// 	chan2 <- "Pong"
	// }()

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case chan1Val := <-chan1:
	// 		fmt.Println("Recieving data from chan1", chan1Val)
	// 	case chan2Val := <-chan2:
	// 		fmt.Println("Recieving data from chan2", chan2Val)
	// 	}
	// }

}
