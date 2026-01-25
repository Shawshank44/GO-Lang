package main

import (
	"embed"
	"fmt"
	"gRPC_school_api/internals/api/handlers"
	"gRPC_school_api/internals/api/interceptors"
	"gRPC_school_api/internals/repositories/mongodb"
	"gRPC_school_api/pkg/utils"
	pb "gRPC_school_api/proto/gen"
	"log"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

//go:embed .env
var envFile embed.FS

func LoadEnvFromEmbeddedFile() {
	content, err := envFile.ReadFile((".env"))
	if err != nil {
		log.Fatal(err.Error())
	}

	tempFile, err := os.CreateTemp("", ".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer os.Remove(tempFile.Name())

	_, err = tempFile.Write(content)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = tempFile.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = godotenv.Load(tempFile.Name())
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	mongodb.CreateMongoClient()

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("error in loading the .env file", err)
	// }

	LoadEnvFromEmbeddedFile()

	cert := os.Getenv("CERT_FILE")
	key := os.Getenv("KEY_FILE")

	creds, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatal("failed to load TLS from file")
	}

	server := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptors.NewRateLimiter(5, time.Minute).RateLimiterInterceptor, interceptors.ResponseTimeInterceptors, interceptors.AuthenticationInterceptor), grpc.Creds(creds))
	pb.RegisterTeachersServiceServer(server, &handlers.Server{})
	pb.RegisterStudentsServiceServer(server, &handlers.Server{})
	pb.RegisterExecsServiceServer(server, &handlers.Server{})

	reflection.Register(server)

	port := os.Getenv("SERVER_PORT")
	fmt.Println("GRPC Server is running on port ", port)

	go utils.JwtStore.CleanUpExpiredToken()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Error in lauching the server : ", err)
	}

	err = server.Serve(lis)
	if err != nil {
		log.Fatal("Failed to SERVE", err)
	}
}
