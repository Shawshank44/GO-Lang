package interceptors

import (
	"context"
	"fmt"
	"gRPC_school_api/pkg/utils"
	"log"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthenticationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("Authenticator begins.")

	skipMethods := map[string]bool{
		"/main.ExecsService/Login":          true,
		"/main.ExecsService/ForgotPassword": true,
		"/main.ExecsService/ResetPassword":  true,
	}

	if skipMethods[info.FullMethod] {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "metadata unavailable")
	}
	authHeader, ok := md["authorization"]
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "authorization unavailable")
	}

	tokenstr := strings.TrimPrefix(authHeader[0], "Bearer ")
	tokenstr = strings.TrimSpace(tokenstr)

	ok = utils.JwtStore.IsLoggedOut(tokenstr)
	if ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized access")
	}
	jwtSecret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "unauthorized access")
	}

	if !token.Valid {
		return nil, status.Error(codes.Unauthenticated, "unauthorized access")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "unauthorized access")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}

	userId := claims["uid"].(string)
	user := claims["user"].(string)
	expf := claims["exp"].(float64)
	expI := int64(expf)
	exp := fmt.Sprintf("%v", expI)

	newCtx := context.WithValue(ctx, utils.ContextKey("role"), role)
	newCtx = context.WithValue(newCtx, utils.ContextKey("userId"), userId)
	newCtx = context.WithValue(newCtx, utils.ContextKey("username"), user)
	newCtx = context.WithValue(newCtx, utils.ContextKey("expiresAt"), exp)

	log.Println("Authenticator Ends.")

	return handler(newCtx, req)
}
