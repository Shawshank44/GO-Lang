package mongodb

import (
	"context"
	"errors"
	"fmt"
	"gRPC_school_api/internals/models"
	"gRPC_school_api/pkg/utils"
	pb "gRPC_school_api/proto/gen"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddTeachersToDB(ctx context.Context, teachersFromReq []*pb.Teacher) ([]*pb.Teacher, error) {
	client, err := CreateMongoClient()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal error")
	}
	defer client.Disconnect(ctx)

	newTeachers := make([]*models.Teacher, len(teachersFromReq))
	for i, pbTeacher := range teachersFromReq {
		newTeachers[i] = MapPbTeacherToModelTeacher(pbTeacher)
	}

	var addedTeachers []*pb.Teacher

	for _, teacher := range newTeachers {
		result, err := client.Database("School").Collection("teachers").InsertOne(ctx, teacher)
		if err != nil {
			return nil, utils.ErrorHandler(err, "error adding value to database")
		}

		objectid, ok := result.InsertedID.(primitive.ObjectID)
		if ok {
			teacher.Id = objectid.Hex()
		}

		pbTeacher := MapModelTeacherToPb(*teacher)
		addedTeachers = append(addedTeachers, pbTeacher)
	}
	return addedTeachers, nil
}

func GetTeachersfromDB(ctx context.Context, sortOptions primitive.D, filter primitive.M, pageNumber, pageSize uint32) ([]*pb.Teacher, error) {
	client, err := CreateMongoClient()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal Error")
	}
	defer client.Disconnect(ctx)
	collection := client.Database("School").Collection("teachers")

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

	teachers, err := DecodeEntities(ctx, cursor, pbModel, newModel)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal error")
	}
	return teachers, nil
}

func pbModel() *pb.Teacher { return &pb.Teacher{} }

func newModel() *models.Teacher { return &models.Teacher{} }

func UpdateTeachersinDB(ctx context.Context, PbTeachers []*pb.Teacher) ([]*pb.Teacher, error) {
	client, err := CreateMongoClient()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal error")
	}
	defer client.Disconnect(ctx)
	var updatedTeachers []*pb.Teacher
	for _, teacher := range PbTeachers {
		if teacher.Id == "" {
			return nil, utils.ErrorHandler(errors.New("id Cannot be blank"), "id cannot be blank")
		}
		modelTeacher := MapPbTeacherToModelTeacher(teacher)
		objId, err := primitive.ObjectIDFromHex(teacher.Id)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Internal error")
		}
		// Convert ModelTeacher to BSON document:
		modelDoc, err := bson.Marshal(modelTeacher)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Internal error")
		}
		var updateDoc bson.M
		err = bson.Unmarshal(modelDoc, &updateDoc)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Internal error")
		}
		// Removing the Id field from the update document
		delete(updateDoc, "_id")

		_, err = client.Database("School").Collection("teachers").UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": updateDoc})
		if err != nil {
			return nil, utils.ErrorHandler(err, fmt.Sprintf("Error updating teacher id : %v", teacher.Id))
		}
		updatedTeacher := MapModelTeacherToPb(*modelTeacher)
		updatedTeachers = append(updatedTeachers, updatedTeacher)
	}
	return updatedTeachers, nil
}
