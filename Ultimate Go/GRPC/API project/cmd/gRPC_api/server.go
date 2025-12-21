package main

import (
	"fmt"
	"gRPC_school_api/internals/api/handlers"
	"gRPC_school_api/internals/repositories/mongodb"
	pb "gRPC_school_api/proto/gen"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	mongodb.CreateMongoClient()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error in loading the .env file", err)
	}

	server := grpc.NewServer()
	pb.RegisterTeachersServiceServer(server, &handlers.Server{})
	pb.RegisterStudentsServiceServer(server, &handlers.Server{})
	pb.RegisterExecsServiceServer(server, &handlers.Server{})

	reflection.Register(server)

	port := os.Getenv("SERVER_PORT")
	fmt.Println("GRPC Server is running on port ", port)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Error in lauching the server : ", err)
	}

	err = server.Serve(lis)
	if err != nil {
		log.Fatal("Failed to SERVE", err)
	}
}
