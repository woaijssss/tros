package trerror

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	SUCCESS       = 200 //成功
	ERROR         = 500 //失败
	InvalidParams = 400 //参数错误
)

var (
	TR_SUCCESS               = &TrError{200, "成功"}
	TR_BAD_REQUEST           = &TrError{400, "无效的请求"}
	TR_ERROR                 = &TrError{500, "服务器错误"}
	TR_INVALID_ERROR         = &TrError{4000, "未知的错误类型"}
	TR_SYSTEM_ERROR          = &TrError{5001, "系统错误"}
	TR_SYSTEM_BUSY           = &TrError{5002, "系统繁忙"}
	TR_TIMEOUT               = &TrError{5003, "请求超时"}
	TR_INVALID_TOKEN         = &TrError{4001, "无效token"}
	TR_FILE_TOO_LARGE        = &TrError{4002, "文件过大"}
	TR_DUPLICATE_PRIMARY_KEY = &TrError{4003, "重复主键"}
	TR_LOGIN_ERROR           = &TrError{4004, "登录错误"}
	TR_NOT_LOGIN             = &TrError{4005, "用户未登录"}
	TR_USER_NOT_EXIST        = &TrError{4006, "用户不存在"}
	TR_DISABLED_USER         = &TrError{4007, "用户已被禁用"}
	TR_WRONG_PASSWORD        = &TrError{4008, "密码错误"}
	TR_NO_PERMISSION         = &TrError{4019, "无权限"}
	TR_ILLEGAL_OPERATION     = &TrError{4010, "非法操作"}
	TR_RECORD_NOT_FOUND      = &TrError{4011, "记录不存在"}
	TR_EMAIL_REGISTERED      = &TrError{4012, "邮箱已被注册"}
	TR_LOGIN_UNSUPPORT       = &TrError{4013, "暂不支持此方式登录"}

	// ugc校验错误
	ContentIllegal = &TrError{10201, "文字包含违规信息"}
	ImageIllegal   = &TrError{10202, "图片包含违规信息"}
	VideoIllegal   = &TrError{10203, "视频包含违规信息"}
	// 数据库错误
	DBNotFoundError = &TrError{Code: 80001, Message: "data not found"}

	TR_ACCESS_TOO_FREQUENTLY = &TrError{99999, "访问太频繁"}
)

type TrError struct {
	Code    int32
	Message string
}

func (te *TrError) Error() string {
	return te.Message
}

// GetMessage 建议使用，获取错误消息
func (te *TrError) GetMessage() string {
	return te.Message
}

// GetCodeInt32 建议使用，获取int32类型的错误码
func (te *TrError) GetCodeInt32() int32 {
	return te.Code
}

// GRPCStatus 需要继承该方法，才能正常返回应用定义的错误码
func (te *TrError) GRPCStatus() *status.Status {
	return status.New(codes.Code(te.Code), te.Message)
}

// GetCodeInt 建议使用，获取int类型的错误码
func (te *TrError) GetCodeInt() int {
	return int(te.Code)
}

func DefaultTrError(msg string) *TrError {
	return &TrError{
		Code:    -1,
		Message: msg,
	}
}

func NewTrError(code int32, msg string) *TrError {
	return &TrError{
		Code:    code,
		Message: msg,
	}
}

// NewF Error with reason fmt
func NewF(code uint32, message, reason string, args ...interface{}) *TrError {
	if len(args) > 0 {
		reason = fmt.Sprintf(reason, args...)
	}
	//return New(code, message, reason, code)
	return New(code, message)
}

// NewErrorWithF Error with reason fmt
func NewErrorWithF(code uint32, internalCode uint32, message, reason string, args ...interface{}) *TrError {
	if len(args) > 0 {
		reason = fmt.Sprintf(reason, args...)
	}
	//return New(code, message, reason, internalCode)
	return New(code, message)
}

// New Error with message and reason
// func New(code uint32, message, reason string, internalCode uint32) *TrError {
func New(code uint32, message string) *TrError {
	return &TrError{
		Code:    int32(code),
		Message: message,
		//Reason:       reason,
		//Metadata:     metadata,
	}
}
