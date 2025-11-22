package main

import (
	"context"
	"log"
	mainapipb "simplegrpcclient/Proto/gen"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connec, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Did not connect : ", err)
		return
	}

	defer connec.Close()

	client := mainapipb.NewCalculateClient(connec)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := mainapipb.AddRequest{
		A: 10,
		B: 300,
	}
	res, err := client.Add(ctx, &req)
	if err != nil {
		log.Fatalln("Could not add", err)
		return
	}

	log.Println("Sum :", res.Sum)

}
