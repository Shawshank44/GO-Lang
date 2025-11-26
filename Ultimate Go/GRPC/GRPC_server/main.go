package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "simplegrpcserve/proto/gen"
	farewellpb "simplegrpcserve/proto/gen/farewell"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedCalculateServer
	pb.UnimplementedGreeterServer
	// farewellpb.UnimplementedAufwiedersehenServer
	pb.UnimplementedBidFarewellServer
}

type Greeter struct {
	pb.UnimplementedGreeterServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	sum := req.A + req.B
	log.Println("Sum : ", sum)
	return &pb.AddResponse{
		Sum: req.A + req.B,
	}, nil
}

func (s *Greeter) Add(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello %s, Nice to Recieve request from you.", req.Name),
	}, nil
}

func (s *Greeter) Greet(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello %s, Nice to Recieve request from you.", req.Name),
	}, nil
}

func (s *server) BidGoodBye(ctx context.Context, req *farewellpb.GoodByeRequest) (*farewellpb.GoodByeResponse, error) {
	return &farewellpb.GoodByeResponse{
		Message: fmt.Sprintf("Thankyou %s, Have a nice day", req.Name),
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
	pb.RegisterGreeterServer(GRPC, &Greeter{})
	pb.RegisterBidFarewellServer(GRPC, &server{})
	// farewellpb.RegisterAufwiedersehenServer(GRPC, &server{})

	log.Println("Server is running on port ", port)
	err = GRPC.Serve(lis)
	if err != nil {
		log.Fatal("Failed to Serve", err)
		return
	}

}
