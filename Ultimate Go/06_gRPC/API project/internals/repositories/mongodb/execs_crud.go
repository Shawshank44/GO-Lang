package mongodb

import (
	"context"
	"gRPC_school_api/internals/models"
	"gRPC_school_api/pkg/utils"
	pb "gRPC_school_api/proto/gen"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddExecsToDB(ctx context.Context, execsFromReq []*pb.Exec) ([]*pb.Exec, error) {
	client, err := CreateMongoClient()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal error")
	}
	defer client.Disconnect(ctx)

	newExecs := make([]*models.Exec, len(execsFromReq))
	for i, pbExec := range execsFromReq {
		newExecs[i] = MapPbExecToModelExec(pbExec)
		hashedPassword, err := utils.HashPassword(newExecs[i].Password)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Internal error")
		}
		newExecs[i].Password = hashedPassword
		currentTime := time.Now().Format(time.RFC3339)
		newExecs[i].UserCreatedAt = currentTime
		newExecs[i].InactiveStatus = false
	}

	var addedExecs []*pb.Exec

	for _, exec := range newExecs {
		result, err := client.Database("School").Collection("execs").InsertOne(ctx, exec)
		if err != nil {
			return nil, utils.ErrorHandler(err, "error adding value to database")
		}

		objectid, ok := result.InsertedID.(primitive.ObjectID)
		if ok {
			exec.Id = objectid.Hex()
		}

		pbExec := MapModelExecToPb(*exec)
		addedExecs = append(addedExecs, pbExec)
	}
	return addedExecs, nil
}

func GetExecsfromDB(ctx context.Context, sortOptions primitive.D, filter primitive.M, pageNumber, pageSize uint32) ([]*pb.Exec, error) {
	client, err := CreateMongoClient()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal Error")
	}
	defer client.Disconnect(ctx)
	collection := client.Database("School").Collection("execs")

	findOptions := options.Find()
	findOptions.SetSkip(int64((pageNumber - 1) * pageSize))
	findOptions.SetLimit(int64(pageSize))

	if len(sortOptions) > 0 {
		findOptions.SetSort(sortOptions)
	}
	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal Error")
	}
	defer cursor.Close(ctx)

	execs, err := DecodeEntities(ctx, cursor, pbModelExec, newModelExec)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal error")
	}
	return execs, nil
}

func pbModelExec() *pb.Exec { return &pb.Exec{} }

func newModelExec() *models.Exec { return &models.Exec{} }
