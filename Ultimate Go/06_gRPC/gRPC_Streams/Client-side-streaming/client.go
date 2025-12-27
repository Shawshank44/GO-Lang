package main

import (
	mainpb "client_stream/proto/gen"
	"context"
	"io"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connection, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer connection.Close()

	client := mainpb.NewCalculatorClient(connection)

	ctx := context.Background()

	//=================================== Server side streaming:=================================== //
	req := &mainpb.FibonacciRequest{
		N: 20,
	}

	stream, err := client.GenerateFibonacci(ctx, req)
	if err != nil {
		log.Fatalln(err)
		return
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Println("End of File")
			break
		}
		if err != nil {
			log.Fatalln(err)
			return
		}

		log.Println(res.GetNumber())
	}

	//=================================== Client side streaming: ===================================//
	stream1, err := client.SendNumbers(ctx)
	if err != nil {
		log.Fatalln(err)
		return
	}
	for num := range 9 {
		log.Println("Sending :", num)
		err := stream1.Send(&mainpb.NumberRequest{Number: int32(num)})
		if err != nil {
			log.Fatalln(err)
			return
		}
		time.Sleep(time.Second)
	}

	res, err := stream1.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("Sum", res.Sum)

	//=================================== BI Directional streaming: ===================================//
	var wg sync.WaitGroup

	chatStream, err := client.Chat(ctx)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// waitC := make(chan struct{})
	wg.Add(2)
	// Go Routine for send messages
	go func() {
		defer wg.Done()
		messages := []string{"Hello", "How are you?", "Good Bye"}
		for _, message := range messages {
			err := chatStream.Send(&mainpb.ChatMessage{
				Message: message,
			})
			if err != nil {
				log.Fatalln(err)
				return
			}
			time.Sleep(time.Second)
		}
		chatStream.CloseSend()
	}()

	// Go Routine for receiving messages
	go func() {
		defer wg.Done()
		for {
			res, err := chatStream.Recv()
			if err == io.EOF {
				log.Println("End of File")
				break
			}
			if err != nil {
				log.Fatalln(err)
				return
			}
			log.Println(res.GetMessage())
		}
		// close(waitC)
	}()
	// <-waitC
	wg.Wait()
}
