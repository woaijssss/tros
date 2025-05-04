package http

import (
	"context"
	context2 "github.com/woaijssss/tros/context"
	trlogger "github.com/woaijssss/tros/logx"
	"github.com/woaijssss/tros/pkg/utils"
	"google.golang.org/grpc"
	"time"
)

// UnaryServerInterceptor returns unary gRpc tracing middleware
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx = context2.InsertTraceID(ctx) // set log id
		//ctx = context2.InsertRemoteIp(ctx) // set remote ip
		ctx = context2.InsertAllInfo(ctx)

		t1 := time.Now()
		resp, err = handler(ctx, req)
		t2 := time.Now()
		elapsed := utils.CalcMillisecondsBetween(t1, t2)
		trlogger.Infof(ctx, "[%s] api execute time: [%+v] ms\n", context2.GetRequestUrlFromCtx(ctx), elapsed)

		return
	}
}
