package handlers

import (
	"context"
	"gRPC_school_api/internals/models"
	"gRPC_school_api/internals/repositories/mongodb"
	"gRPC_school_api/pkg/utils"
	pb "gRPC_school_api/proto/gen"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AddTeachers(ctx context.Context, req *pb.Teachers) (*pb.Teachers, error) {
	for _, teacher := range req.GetTeachers() {
		if teacher.Id != "" {
			return nil, status.Error(codes.InvalidArgument, "Request is in incorrect format")
		}
	}
	addedTeachers, err := mongodb.AddTeachersToDB(ctx, req.Teachers)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Teachers{Teachers: addedTeachers}, nil
}

func (s *Server) GetTeachers(ctx context.Context, req *pb.GetTeachersRequest) (*pb.Teachers, error) {
	// Getting the filters from the request
	filter, err := BuildFilterForTeacher(req.Teacher)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// Sorting, getting the sort options from the request
	sortOptions := BuildSortOptions(req.GetSortBy())
	// Access the database to fetch data
	teachers, err := mongodb.GetTeachersfromDB(ctx, sortOptions, filter)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Teachers{Teachers: teachers}, nil
}

func BuildFilterForTeacher(teacherObj *pb.Teacher) (bson.M, error) {
	filter := bson.M{}

	if teacherObj == nil {
		return filter, nil
	}

	var ModelTeacher models.Teacher
	modelVal := reflect.ValueOf(&ModelTeacher).Elem()
	modelType := modelVal.Type()

	reqVal := reflect.ValueOf(teacherObj).Elem()
	reqType := reqVal.Type()

	for i := 0; i < reqVal.NumField(); i++ {
		fieldVal := reqVal.Field(i)
		fieldName := reqType.Field(i).Name

		if fieldVal.IsValid() && !fieldVal.IsZero() {
			modelfield := modelVal.FieldByName(fieldName)
			if modelfield.IsValid() && modelfield.CanSet() {
				modelfield.Set(fieldVal)
			}
		}
	}

	// Now we iterate over the modelTeacher to build filter using bson.M
	for i := 0; i < modelVal.NumField(); i++ {
		fieldVal := modelVal.Field(i)
		// fieldName := modelType.Field(i).Name
		if fieldVal.IsValid() && !fieldVal.IsZero() {
			bsonTag := modelType.Field(i).Tag.Get("bson")
			bsonTag = strings.TrimSuffix(bsonTag, ",omitempty")
			if bsonTag == "_id" {
				objId, err := primitive.ObjectIDFromHex(teacherObj.Id)
				if err != nil {
					return nil, utils.ErrorHandler(err, "Internal error")
				}
				filter[bsonTag] = objId
			} else {
				filter[bsonTag] = fieldVal.Interface().(string)
			}

		}
	}
	return filter, nil
}

func BuildSortOptions(sortfields []*pb.SortField) bson.D {
	var sortOptions bson.D

	for _, sortfield := range sortfields {
		order := 1
		if sortfield.GetOrder() == pb.Order_DESC {
			order = -1
		}
		sortOptions = append(sortOptions, bson.E{Key: sortfield.Field, Value: order})
	}
	return sortOptions
}
