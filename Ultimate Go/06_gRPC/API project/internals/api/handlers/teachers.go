package handlers

import (
	"context"
	"errors"
	"gRPC_school_api/internals/models"
	"gRPC_school_api/internals/repositories/mongodb"
	"gRPC_school_api/pkg/utils"
	pb "gRPC_school_api/proto/gen"

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
	filter, err := BuildFilter(req.Teacher, &models.Teacher{})
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

func (s *Server) UpdateTeachers(ctx context.Context, req *pb.Teachers) (*pb.Teachers, error) {
	updatedTeachers, err := mongodb.UpdateTeachersinDB(ctx, req.Teachers)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.Teachers{Teachers: updatedTeachers}, nil
}

func (s *Server) DeleteTeachers(ctx context.Context, req *pb.TeacherIDs) (*pb.DeleteTeachersConfirmation, error) {
	ids := req.GetIds()
	var teacherIdsToDelete []string

	for _, v := range ids {
		if v.Id == "" {
			return nil, errors.New("id Field cannot be blank")
		}
		teacherIdsToDelete = append(teacherIdsToDelete, v.Id)
	}

	client, err := mongodb.CreateMongoClient()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal error")
	}
	defer client.Disconnect(ctx)

	objectIds := make([]primitive.ObjectID, len(teacherIdsToDelete))
	for i, id := range teacherIdsToDelete {
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Invalid ID")
		}
		objectIds[i] = objectId
	}

	filter := bson.M{"_id": bson.M{"$in": objectIds}}

	res, err := client.Database("School").Collection("teachers").DeleteMany(ctx, filter)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error in deleting the _Id")
	}

	if res.DeletedCount == 0 {
		return nil, utils.ErrorHandler(err, "No teachers were deleted")
	}

	deletedIds := make([]string, res.DeletedCount)
	for i, id := range objectIds {
		deletedIds[i] = id.Hex()
	}

	return &pb.DeleteTeachersConfirmation{
		Status:     "Teachers Successfully Deleted",
		DeletedIds: deletedIds,
	}, nil
}
