package http

import (
	"context"
	context2 "github.com/woaijssss/tros/context"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor returns unary gRpc tracing middleware
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx = context2.InsertTraceID(ctx) // set log id
		//ctx = context2.InsertRemoteIp(ctx) // set remote ip
		ctx = context2.InsertAllInfo(ctx)

		return handler(ctx, req)
	}
}
