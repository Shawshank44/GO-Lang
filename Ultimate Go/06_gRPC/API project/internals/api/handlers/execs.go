package handlers

import (
	"context"
	"gRPC_school_api/internals/models"
	"gRPC_school_api/internals/repositories/mongodb"
	"gRPC_school_api/pkg/utils"
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

func (s *Server) UpdateExecs(ctx context.Context, req *pb.Execs) (*pb.Execs, error) {
	updatedExecs, err := mongodb.UpdateExecsinDB(ctx, req.Execs)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.Execs{Execs: updatedExecs}, nil
}

func (s *Server) DeleteExecs(ctx context.Context, req *pb.ExecIDs) (*pb.DeleteExecsConfirmation, error) {
	// ids := req.GetIds()
	// var execIdsToDelete []string

	// for _, v := range ids {
	// 	if v.Id == "" {
	// 		return nil, errors.New("id Field cannot be blank")
	// 	}
	// 	execIdsToDelete = append(execIdsToDelete, v.Id)
	// }

	deletedIds, err := mongodb.DeleteExecsFromDB(ctx, req.GetIds())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DeleteExecsConfirmation{
		Status:     "Execs Successfully Deleted",
		DeletedIds: deletedIds,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.ExecLoginRequest) (*pb.ExecLoginResponse, error) {
	if req.Username == "" || req.Password == "" {
		return nil, status.Error(codes.Unauthenticated, "Fields of username and password cannot be empty")
	}

	exec, err := mongodb.GetUserByUserName(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if exec.InactiveStatus {
		return nil, utils.ErrorHandler(err, "Internal error")
	}

	err = utils.VerifyPassword(req.GetPassword(), exec.Password)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Incorrect username/password")
	}

	token, err := utils.SignToken(exec.Id, exec.Username, exec.Role)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Could not create the token")
	}

	return &pb.ExecLoginResponse{
		Status: true,
		Token:  token,
	}, nil
}

func (s *Server) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordResponse, error) {
	if req.Id == "" || req.NewPassword == "" || req.CurrentPassword == "" {
		return nil, status.Error(codes.Unauthenticated, "Fields cannot be blank")
	}
	username, role, err := mongodb.UpdatePasswordInDB(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	token, err := utils.SignToken(req.Id, username, role)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal error")
	}

	return &pb.UpdatePasswordResponse{
		PasswordUpdated: true,
		Token:           token,
	}, nil
}

func (s *Server) DeactivateUser(ctx context.Context, req *pb.ExecIDs) (*pb.Confirmation, error) {
	result, err := mongodb.DeactivateUserInDB(ctx, req.GetIds())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Unable to deactivate user")
	}

	return &pb.Confirmation{
		Confirmation: result.ModifiedCount > 0,
	}, nil
}

func (s *Server) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error) {
	email := req.GetEmail()

	message, err := mongodb.ForgotPasswordDB(ctx, email)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Unable to reset the password")
	}
	return &pb.ForgotPasswordResponse{
		Confirmation: true,
		Message:      message,
	}, nil
}
