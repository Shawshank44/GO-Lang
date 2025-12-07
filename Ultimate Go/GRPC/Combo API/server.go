package main

import (
	"context"
	"crypto/tls"
	"fmt"
	mainpb "grpcgatewayproject/proto/gen"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
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

func runGRPCServer(certFile, keyFile string) {
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to load TLS cert : %v", err)
	}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln("Unable to start server", err)
		return
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	// enable reflection :
	reflection.Register(grpcServer) // DO NOT USE THIS IN PRODUCTION.

	mainpb.RegisterGreeterServer(grpcServer, &server{})

	log.Println("gRPC Server running in port :50051")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func loadTLSCredentials(certFile, keyFile string) tls.Certificate {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalln("failed to load the credentials")
	}
	return cert
}

// Rest functionality :
func runGatewayServer(certFile, keyFile string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	}))}
	err := mainpb.RegisterGreeterHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalln("failed to register gRPC-gateway handler : ", err)
	}

	tlsconfig := &tls.Config{
		Certificates: []tls.Certificate{loadTLSCredentials(certFile, keyFile)},
	}

	server := &http.Server{
		Addr:      ":8080",
		Handler:   mux,
		TLSConfig: tlsconfig,
	}

	log.Println("HTTP server is running on port: 8080...")

	err = server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Fatalln("failed to Listen :", err)
	}

	// Non TLS version
	// http.ListenAndServe(":8080", mux)
	// if err != nil {
	// 	log.Fatalln("failed to Listen : ", err)
	// }
}

func main() {
	cert := "cert.pem"
	key := "key.pem"
	go runGRPCServer(cert, key)
	runGatewayServer(cert, key)
}
