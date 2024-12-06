package user

import (
	"context"
	"fmt"
	"github.com/woaijssss/tros/constants"
	context2 "github.com/woaijssss/tros/context"
	"github.com/woaijssss/tros/pkg/utils"
	"github.com/woaijssss/tros/trerror"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	userNoChar              = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	userNoPureLowercaseChar = "abcdefghijklmnopqrstuvwxyz0123456789"
	phoneLengthChn          = 11 // Chinese phone number length
)

// CheckPermission 需要校验token并从中提取user_id的接口，都需要调用该函数
func CheckPermission(ctx context.Context) (string, error) {
	tokenInfo, err := getUserInfoFromToken(ctx)
	if err != nil {
		return "", err
	}
	return tokenInfo.UserId, err
}

func GetTokenFromContext(ctx context.Context) (*utils.TokenInfo, error) {
	return getUserInfoFromToken(ctx)
}

// GenUserNoPrefix 生成用户编号的前5位
func GenUserNoPrefix() string {
	length := 5
	//rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = userNoChar[rand.Intn(len(userNoChar))]
	}
	return string(result)
}

// GenerateUniqueIdPureLowercase Generate custom length pure lowercase user No.
func GenerateUniqueIdPureLowercase(length int) string {
	rand.Seed(time.Now().UnixNano())
	// 用于存储生成的用户编号
	var uniqueID strings.Builder
	// 循环生成12位字符
	for i := 0; i < length; i++ {
		// 随机选择字符集合中的一个字符
		randomIndex := rand.Intn(len(userNoPureLowercaseChar))
		uniqueID.WriteRune(rune(userNoPureLowercaseChar[randomIndex]))
	}
	return uniqueID.String()
}

// GenerateUniqueId Generate custom length user No.
func GenerateUniqueId(length int) string {
	rand.Seed(time.Now().UnixNano())
	// 用于存储生成的用户编号
	var uniqueID strings.Builder
	// 循环生成12位字符
	for i := 0; i < length; i++ {
		// 随机选择字符集合中的一个字符
		randomIndex := rand.Intn(len(userNoChar))
		uniqueID.WriteRune(rune(userNoChar[randomIndex]))
	}
	return uniqueID.String()
}

// HidePhoneNumber Keep the last 4 digits of phone number
func HidePhoneNumber(phone string) string {
	if len(phone) == phoneLengthChn {
		return phone[:3] + "****" + phone[7:]
	}
	return phone
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
		//return nil, status.Errorf(codes.PermissionDenied, "missing authorization token")
		return nil, trerror.TR_NOT_LOGIN
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

// GetUserGender todo 临时的兼容，后面db里gender字段修改为整数类型后，此方法作废
func GetUserGender(gender string) int32 {
	gdi, err := strconv.Atoi(gender)
	if err != nil {
		return 0
	}
	return int32(gdi)
}
