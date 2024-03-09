package grpc

import (
	"context"
	"gitee.com/idigpower/tros/constants"
	context2 "gitee.com/idigpower/tros/context"
	trlogger "gitee.com/idigpower/tros/logx"
	"gitee.com/idigpower/tros/server/middleware"
	"time"

	grpclogging "github.com/grpc-ecosystem/go-grpc-middleware/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc"
)

const (
	// HeaderRequestID header name for request id
	HeaderRequestID     = "X-Request-Id"
	defaultKeyValuesCap = 8
)

// UnaryServerInterceptor returns a new unary server interceptors that performs logging
func UnaryServerInterceptor(config GRpcConfig) grpc.UnaryServerInterceptor {
	gl := newGpcLogger(config)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		method := info.FullMethod
		md := metautils.ExtractIncoming(ctx)
		if !gl.shouldLog(md, method) {
			return handler(ctx, req)
		}

		kv := make([]interface{}, 0, defaultKeyValuesCap)
		st := time.Now()

		kv = gl.beforePopulate(ctx, kv, md, method, false)
		resp, err := handler(ctx, req)
		defer func() {
			kv = gl.afterPopulate(kv, st, err)
			gl.logger.Infof(ctx, "grpc request", kv...)
		}()

		return resp, err
	}
}

// StreamServerInterceptor returns a new stream server interceptors that performs logging
func StreamServerInterceptor(config GRpcConfig) grpc.StreamServerInterceptor {
	gl := newGpcLogger(config)
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := ss.Context()
		method := info.FullMethod
		md := metautils.ExtractIncoming(ctx)

		if !gl.shouldLog(md, method) {
			return handler(srv, ss)
		}

		kv := make([]interface{}, 0, defaultKeyValuesCap)
		st := time.Now()

		kv = gl.beforePopulate(ctx, kv, md, method, true)
		err := handler(srv, ss)
		defer func() {
			kv = gl.afterPopulate(kv, st, err)
			gl.logger.Infof(ctx, "grpc request", kv...)
		}()

		return err
	}
}

type grpcLogger struct {
	logger                    *trlogger.TrLogger
	excludes                  map[string]struct{}
	excludeGRpcGatewayRequest bool
}

func newGpcLogger(config GRpcConfig) *grpcLogger {
	logger := config.Logger
	if logger == nil {
		logger = trlogger.DefaultTrLogger()
	}

	return &grpcLogger{
		logger:                    logger,
		excludes:                  middleware.ExcludePaths(config.Excludes),
		excludeGRpcGatewayRequest: config.ExcludeGRpcGatewayRequest,
	}
}

func (gl *grpcLogger) shouldLog(md metautils.NiceMD, method string) bool {
	if _, ok := gl.excludes[method]; ok {
		return false
	}

	if gl.excludeGRpcGatewayRequest {
		if middleware.IsRequestFromGRpcGateway(md) {
			return false
		}
	}

	return true
}

func (gl *grpcLogger) beforePopulate(ctx context.Context, kv []interface{}, md metautils.NiceMD, method string, stream bool) []interface{} {
	kv = append(kv,
		"stream", stream,
		"method", method,
		"req_id", md.Get(HeaderRequestID),
		"gw", middleware.IsRequestFromGRpcGateway(md),
		constants.TraceId, context2.GenTraceID(),
	)
	return kv
}

func (gl *grpcLogger) afterPopulate(kv []interface{}, st time.Time, err error) []interface{} {
	kv = append(kv,
		"code", grpclogging.DefaultErrorToCode(err),
		"cost", time.Since(st),
	)

	return kv
}
