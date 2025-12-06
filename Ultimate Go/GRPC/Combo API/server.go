package main

import (
	"context"
	"fmt"
	mainpb "grpcgatewayproject/proto/gen"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	mainpb.UnimplementedGreeterServer
}

func (s *server) Greet(ctx context.Context, req *mainpb.HelloRequest) (*mainpb.HelloRespoonse, error) {
	err := req.Validate()
	if err != nil {
		log.Printf("Validation failed : %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid request : %v", err)
	}

	return &mainpb.HelloRespoonse{
		Message: fmt.Sprintf("Hello, %s", req.GetName()),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln("Unable to start server", err)
		return
	}

	grpcServer := grpc.NewServer()

	// enable reflection :
	reflection.Register(grpcServer) // DO NOT USE THIS IN PRODUCTION.

	mainpb.RegisterGreeterServer(grpcServer, &server{})

	log.Println("Server running in port :50051")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
