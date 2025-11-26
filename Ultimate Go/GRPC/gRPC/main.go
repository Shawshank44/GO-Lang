package main

import (
	"context"
	"log"
	"net"
	pb "simplegrpcserve/Proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedCalculateServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	sum := req.A + req.B
	log.Println("Sum : ", sum)
	return &pb.AddResponse{
		Sum: req.A + req.B,
	}, nil
}

func main() {
	cert := "cert.pem"
	key := "key.pem"

	port := ":50051"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen", err)
		return
	}

	creds, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatal("Failed to load key and cert", err)
		return
	}

	GRPC := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterCalculateServer(GRPC, &server{})

	log.Println("Server is running on port ", port)
	err = GRPC.Serve(lis)
	if err != nil {
		log.Fatal("Failed to Serve", err)
		return
	}
}
