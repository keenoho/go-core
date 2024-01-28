package grpc_engine

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func LoggerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		log.Println(info)
		return handler(ctx, req)
	}
}
