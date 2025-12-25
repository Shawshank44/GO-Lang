package mongodb

import (
	"context"
	"gRPC_school_api/internals/models"
	"gRPC_school_api/pkg/utils"
	pb "gRPC_school_api/proto/gen"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
		newTeachers[i] = MapToModelTeacher(pbTeacher)
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

		pbTeacher := MaptoModelTeacherDB(teacher)
		addedTeachers = append(addedTeachers, pbTeacher)
	}
	return addedTeachers, nil
}

func MaptoModelTeacherDB(teacher *models.Teacher) *pb.Teacher {
	pbTeacher := &pb.Teacher{}
	modelVal := reflect.ValueOf(*teacher)
	pbVal := reflect.ValueOf(pbTeacher).Elem()
	for i := 0; i < modelVal.NumField(); i++ {
		modelField := modelVal.Field(i)
		modelFieldtype := modelVal.Type().Field(i)
		pbField := pbVal.FieldByName(modelFieldtype.Name)
		if pbField.IsValid() && pbField.CanSet() {
			pbField.Set(modelField)
		}
	}
	return pbTeacher
}

func MapToModelTeacher(pbTeacher *pb.Teacher) *models.Teacher {
	modelTeacher := models.Teacher{}
	pbval := reflect.ValueOf(pbTeacher).Elem()
	modelVal := reflect.ValueOf(&modelTeacher).Elem()
	for i := 0; i < pbval.NumField(); i++ {
		pbField := pbval.Field(i)
		fieldName := pbval.Type().Field(i).Name
		modelfield := modelVal.FieldByName(fieldName)
		if modelfield.IsValid() && modelfield.CanSet() {
			modelfield.Set(pbField)
		}
	}
	return &modelTeacher
}

func GetTeachersfromDB(ctx context.Context, sortOptions bson.D, filter bson.M) ([]*pb.Teacher, error) {
	client, err := CreateMongoClient()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal Error")
	}
	defer client.Disconnect(ctx)
	collection := client.Database("School").Collection("teachers")
	var cursor *mongo.Cursor
	if len(sortOptions) < 1 {
		cursor, err = collection.Find(ctx, filter)
	} else {
		cursor, err = collection.Find(ctx, filter, options.Find().SetSort(sortOptions))
	}
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal Error")
	}
	defer cursor.Close(ctx)

	var teachers []*pb.Teacher

	for cursor.Next(ctx) {
		var teacher models.Teacher
		err := cursor.Decode(&teacher)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Internal error")
		}
		teachers = append(teachers, &pb.Teacher{
			Id:        teacher.Id,
			FirstName: teacher.FirstName,
			LastName:  teacher.LastName,
			Email:     teacher.Email,
			Class:     teacher.Class,
			Subject:   teacher.Subject,
		})
	}
	return teachers, nil
}
