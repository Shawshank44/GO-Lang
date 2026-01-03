package handlers

import (
	"context"
	"gRPC_school_api/internals/models"
	"gRPC_school_api/internals/repositories/mongodb"
	pb "gRPC_school_api/proto/gen"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AddExecs(ctx context.Context, req *pb.Execs) (*pb.Execs, error) {
	for _, exec := range req.GetExecs() {
		if exec.Id != "" {
			return nil, status.Error(codes.InvalidArgument, "Request is in incorrect format")
		}
	}
	addedExecs, err := mongodb.AddExecsToDB(ctx, req.Execs)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Execs{Execs: addedExecs}, nil
}

func (s *Server) GetExecs(ctx context.Context, req *pb.GetExecsRequest) (*pb.Execs, error) {
	// Getting the filters from the request
	filter, err := BuildFilter(req.Exec, &models.Exec{})
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
	execs, err := mongodb.GetExecsfromDB(ctx, sortOptions, filter, pageNumber, pageSize)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Execs{Execs: execs}, nil
}
