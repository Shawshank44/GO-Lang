package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	mainpb "server_stream/proto/gen"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	mainpb.UnimplementedCalculatorServer
}

func (s *server) GenerateFibonacci(req *mainpb.FibonacciRequest, stream mainpb.Calculator_GenerateFibonacciServer) error {
	n := req.N
	a, b := 0, 1

	for i := 0; i < int(n); i++ {
		err := stream.Send(&mainpb.FibonacciResponse{
			Number: int32(a),
		})
		if err != nil {
			return err
		}
		log.Println("Sent Number :", a)
		a, b = b, a+b
		time.Sleep(time.Second)
	}
	return nil
}

func (s *server) SendNumbers(stream mainpb.Calculator_SendNumbersServer) error {
	var sum int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&mainpb.NumberResponse{Sum: sum})
		}
		if err != nil {
			return err
		}
		log.Println(req.GetNumber())
		sum += req.GetNumber()
	}
}

func (s *server) Chat(stream mainpb.Calculator_ChatServer) error {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Receiving values / messages from stream
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Received message : ", req.GetMessage())

		// Read input from terminal :
		fmt.Println("Enter response : ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		input = strings.TrimSpace(input)

		res := &mainpb.ChatMessage{
			Message: input,
		}
		// Sending values through the stream :
		err = stream.Send(res)
		if err != nil {
			return err
		}
	}
	fmt.Println("Returning control")
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln("Unable to start server", err)
		return
	}

	grpcServer := grpc.NewServer()
	mainpb.RegisterCalculatorServer(grpcServer, &server{})

	// enable reflection :
	reflection.Register(grpcServer) // DO NOT USE THIS IN PRODUCTION.

	log.Println("Server running in port :50051")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln(err)
		return
	}

}
