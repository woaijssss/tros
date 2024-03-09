package http

import (
	"context"
	"fmt"
	"gitee.com/idigpower/tros/constants"
	context2 "gitee.com/idigpower/tros/context"
	"gitee.com/idigpower/tros/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// UnaryServerInterceptor returns unary gRpc tracing middleware
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx = context2.InsertTraceID(ctx)

		// 从context中获取metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, fmt.Errorf("no metadata in context")
		}

		// 获取metadata中的"authorization"键对应的值，假设token在这个键下
		tokens := md[constants.Token]
		if len(tokens) == 0 {
			return nil, fmt.Errorf("missing authorization token")
		}

		// 使用token进行后续处理，例如验证等
		// ...
		token := tokens[0]
		if token != "" {

			tokenInfo, err := utils.ParseTokenWithoutVerify(token)
			var userId string
			if err == nil {
				userId = tokenInfo.UserId
			}

			ctx = context2.AddUserID(ctx, userId)
		}

		return handler(ctx, req)
	}
}
