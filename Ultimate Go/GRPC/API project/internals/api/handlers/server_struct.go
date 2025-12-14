package handlers

import pb "gRPC_school_api/proto/gen"

type Server struct {
	pb.UnimplementedTeachersServiceServer
	pb.UnimplementedStudentsServiceServer
	pb.UnimplementedExecsServiceServer
}
