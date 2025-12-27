package mongodb

import (
	"context"
	"fmt"
	"gRPC_school_api/pkg/utils"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoClient() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, utils.ErrorHandler(err, "failed to establish connection with mongo DB")
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err)
		return nil, utils.ErrorHandler(err, "ping failed")
	}

	log.Println("Connected to MongoDB")
	return client, nil
}
