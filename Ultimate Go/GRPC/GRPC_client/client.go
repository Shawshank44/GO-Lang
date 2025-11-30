package main

import (
	"context"
	"log"
	mainapipb "simplegrpcclient/proto/gen"
	farewellpb "simplegrpcclient/proto/gen/farewell"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding/gzip"
)

func main() {
	cert := "cert.pem"
	// key := "key.pem"

	creds, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		log.Fatal("Failed in loading the credentials")
		return
	}

	connec, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(creds), grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	if err != nil {
		log.Fatalln("Did not connect : ", err)
		return
	}

	defer connec.Close()

	client := mainapipb.NewCalculateClient(connec)

	client1 := mainapipb.NewGreeterClient(connec)
	bfclient := mainapipb.NewBidFarewellClient(connec)

	// fwclient := farewellpb.NewAufwiedersehenClient(connec)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := mainapipb.AddRequest{
		A: 10,
		B: 300,
	}
	res, err := client.Add(ctx, &req, grpc.UseCompressor(gzip.Name)) // if compression required for only particular request.
	if err != nil {
		log.Fatalln("Could not add", err)
		return
	}

	reqGreet := &mainapipb.HelloRequest{
		Name: "john",
	}

	res1, err := client1.Greet(ctx, reqGreet)
	if err != nil {
		log.Fatal("Could not greet", err)
		return
	}

	res2, err := client1.Add(ctx, reqGreet)
	if err != nil {
		log.Fatal("Could not add-----", err)
		return
	}

	reqGoodBye := &farewellpb.GoodByeRequest{
		Name: "jane",
	}

	// resfw, err := fwclient.BidGoodBye(ctx, reqGoodBye)
	// if err != nil {
	// 	log.Fatal("Could not farewell", err)
	// 	return
	// }

	resbf, err := bfclient.BidGoodBye(ctx, reqGoodBye)
	if err != nil {
		log.Fatal("Could not farewell", err)
		return
	}

	log.Println("Sum :", res.Sum)
	log.Println("Greet :", res1.Message)
	log.Println("Greet from Add :", res2.Message)
	log.Println("farewell :", resbf.Message)
	// log.Println("Farewell : ", resfw.Message)
	// state := connec.GetState()
	// log.Println("connection state : ", state)

}
