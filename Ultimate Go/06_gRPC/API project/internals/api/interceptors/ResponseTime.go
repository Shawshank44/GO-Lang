package interceptors

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func ResponseTimeInterceptors(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("ResponseTimeInterceptor Begins")
	// Record the Start Time
	start := time.Now()

	// Call the handler to proceed with the request.
	res, err := handler(ctx, req)

	// Calculate the duration
	duration := time.Since(start)

	// Log the request deatils with duration.
	st, _ := status.FromError(err)
	fmt.Printf("Method : %s, Status : %d, Duration : %v \n", info.FullMethod, st.Code(), duration)

	md := metadata.Pairs("X-Response-Time", duration.String())
	grpc.SetHeader(ctx, md)

	log.Println("ResponseTimeInterceptor Ends")
	return res, err
}
