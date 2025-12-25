package handlers

import (
	"context"
	"fmt"
	"gRPC_school_api/internals/models"
	"gRPC_school_api/internals/repositories/mongodb"
	pb "gRPC_school_api/proto/gen"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
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
	BuildFilterForTeacher(req)
	// Sorting, getting the sort options from the request
	BuildSortOptions(req.GetSortBy())
	// Access the database to fetch data

	return nil, nil
}

func BuildFilterForTeacher(req *pb.GetTeachersRequest) {
	filter := bson.M{}

	var ModelTeacher models.Teacher
	modelVal := reflect.ValueOf(&ModelTeacher).Elem()
	modelType := modelVal.Type()

	reqVal := reflect.ValueOf(req.Teacher).Elem()
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
			filter[bsonTag] = fieldVal.Interface().(string)
		}
	}
	fmt.Println(filter)
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
	fmt.Println(sortOptions)
	return sortOptions
}
