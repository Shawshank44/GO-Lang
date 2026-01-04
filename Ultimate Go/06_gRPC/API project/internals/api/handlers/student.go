package handlers

import (
	"context"
	"errors"
	"gRPC_school_api/internals/models"
	"gRPC_school_api/internals/repositories/mongodb"
	pb "gRPC_school_api/proto/gen"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AddStudents(ctx context.Context, req *pb.Students) (*pb.Students, error) {
	for _, student := range req.GetStudents() {
		if student.Id != "" {
			return nil, status.Error(codes.InvalidArgument, "Request is in incorrect format")
		}
	}
	addedStudents, err := mongodb.AddStudentsToDB(ctx, req.Students)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Students{Students: addedStudents}, nil
}

func (s *Server) GetStudents(ctx context.Context, req *pb.GetStudentsRequest) (*pb.Students, error) {
	// Getting the filters from the request
	filter, err := BuildFilter(req.Student, &models.Student{})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// Sorting, getting the sort options from the request
	sortOptions := BuildSortOptions(req.GetSortBy())
	// Access the database to fetch data
	pageNumber := req.GetPageNumber()
	pageSize := req.GetPageSize()

	if pageNumber < 1 {
		pageNumber = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	students, err := mongodb.GetStudentsfromDB(ctx, sortOptions, filter, pageNumber, pageSize)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Students{Students: students}, nil
}

func (s *Server) UpdateStudents(ctx context.Context, req *pb.Students) (*pb.Students, error) {
	updatedStudents, err := mongodb.UpdateStudentsinDB(ctx, req.Students)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.Students{Students: updatedStudents}, nil
}

func (s *Server) DeleteStudents(ctx context.Context, req *pb.StudentIDs) (*pb.DeleteStudentsConfirmation, error) {
	ids := req.GetIds()
	var studentIdsToDelete []string

	for _, v := range ids {
		if v.Id == "" {
			return nil, errors.New("id Field cannot be blank")
		}
		studentIdsToDelete = append(studentIdsToDelete, v.Id)
	}

	deletedIds, err := mongodb.DeleteStudentsFromDB(ctx, studentIdsToDelete)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DeleteStudentsConfirmation{
		Status:     "Students Successfully Deleted",
		DeletedIds: deletedIds,
	}, nil
}
