package user

import (
	"context"
	"fmt"
	"gitee.com/idigpower/tros/constants"
	context2 "gitee.com/idigpower/tros/context"
	"gitee.com/idigpower/tros/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// CheckPermission 需要校验token并从中提取user_id的接口，都需要调用该函数
func CheckPermission(ctx context.Context) (string, error) {
	tokenInfo, err := getUserInfoFromToken(ctx)
	if err != nil {
		return "", err
	}
	return tokenInfo.UserId, err
}

func getUserInfoFromToken(ctx context.Context) (*utils.TokenInfo, error) {
	// 从context中获取metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("no metadata in context")
	}

	// 获取metadata中的"authorization"键对应的值，假设token在这个键下
	tokens := md[constants.Token]
	if len(tokens) == 0 || tokens[0] == "" {
		return nil, status.Errorf(codes.PermissionDenied, "missing authorization token")
	}

	// 使用token进行后续处理，例如验证等
	// ...
	token := tokens[0]
	// todo 增加检验token是否过期，过期的token依然要返回403
	tokenInfo, err := utils.ParseTokenWithoutVerify(token)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "parse token is invalid")
	}
	ctx = context2.AddUserID(ctx, tokenInfo.UserId)
	return tokenInfo, nil // 解析成功
}
