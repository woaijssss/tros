package trerror

import (
	"errors"
	"fmt"
	"github.com/woaijssss/tros/trerror/codes"
	grpcCodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	SUCCESS       = 200 //成功
	ERROR         = 500 //失败
	InvalidParams = 400 //参数错误
)

var (
	TR_SUCCESS               = &TrError{codes.TrSuccess, "成功"}
	TR_BAD_REQUEST           = &TrError{codes.TrBadRequest, "无效的请求"}
	TR_ERROR                 = &TrError{codes.TrError, "服务器错误"}
	TR_INVALID_ERROR         = &TrError{codes.TrInvalidError, "未知的错误类型"}
	TR_INVALID_TOKEN         = &TrError{codes.TrInvalidToken, "无效token"}
	TR_FILE_TOO_LARGE        = &TrError{codes.TrFileTooLarge, "文件过大"}
	TR_DUPLICATE_PRIMARY_KEY = &TrError{codes.TrDuplicatePrimaryKey, "重复主键"}
	TR_LOGIN_ERROR           = &TrError{codes.TrLoginError, "登录错误"}
	TR_NOT_LOGIN             = &TrError{codes.TrNotLogin, "请登录"}
	TR_USER_NOT_EXIST        = &TrError{codes.TrUserNotExist, "用户不存在"}
	TR_DISABLED_USER         = &TrError{codes.TrDisableUser, "用户已被禁用"}
	TR_WRONG_PASSWORD        = &TrError{codes.TrWrongPassword, "密码错误"}
	TR_NO_PERMISSION         = &TrError{codes.TrNoPermission, "无权限"}
	TR_ILLEGAL_OPERATION     = &TrError{codes.TrIllegalOperation, "非法操作"}
	TR_RECORD_NOT_FOUND      = &TrError{codes.TrRecordNotFound, "记录不存在"}
	TR_LOGIN_UNSUPPORT       = &TrError{codes.TrLoginUnSupport, "暂不支持此方式登录"}

	TR_SYSTEM_ERROR          = &TrError{codes.TrSystemError, "系统错误"}
	TR_SYSTEM_BUSY           = &TrError{codes.TrSystemBusy, "系统繁忙"}
	TR_TIMEOUT               = &TrError{codes.TrTimeout, "请求超时"}
	TR_URL_CANNOT_ACCESS_ERR = &TrError{codes.TrUrlCannotAccess, "链接无法访问"}

	// ugc校验错误
	ContentIllegal = &TrError{codes.TrContentIllegal, "文字包含违规信息"}
	ImageIllegal   = &TrError{codes.TrImageIllegal, "图片包含违规信息"}
	VideoIllegal   = &TrError{codes.TrVideoIllegal, "视频包含违规信息"}

	// 数据库错误
	TR_DB_ERROR     = &TrError{codes.TrDbError, "数据库错误"}
	DBNotFoundError = &TrError{codes.TrDataNotFound, "未找到数据"}

	TR_ACCESS_TOO_FREQUENTLY = &TrError{codes.TrAccessTooFrequently, "访问太频繁"}
)

type TrError struct {
	Code    int32
	Message string
}

// WithOutNotFound Support skipping 'no record found' for list classes and setting the judgment operation for empty lists
func WithOutNotFound(err error) error {
	if errors.Is(err, DBNotFoundError) {
		return nil
	}

	return err
}

/*
	GRPCStatus

No need to call, only inherit from 'type grpcstatus interface {GRPCStatus () * Status}',
ensure that the returned application error code is correct
*/
func (te *TrError) GRPCStatus() *status.Status {
	return status.New(grpcCodes.Code(te.Code), te.Message)
}

func (te *TrError) Error() string {
	return te.Message
}

// GetMessage Suggested use to obtain error messages
func (te *TrError) GetMessage() string {
	return te.Message
}

// GetCodeInt32 Suggest using to obtain error codes of type int32
func (te *TrError) GetCodeInt32() int32 {
	return te.Code
}

// GetCodeInt Suggest using to obtain error codes of type int
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
func NewF(code int32, message, reason string, args ...interface{}) *TrError {
	if len(args) > 0 {
		reason = fmt.Sprintf(reason, args...)
	}
	return New(code, message)
}

// NewErrorWithF Error with reason fmt
func NewErrorWithF(code int32, internalCode int32, message, reason string, args ...interface{}) *TrError {
	if len(args) > 0 {
		reason = fmt.Sprintf(reason, args...)
	}
	return New(code, message)
}

// New Error with message and reason
// func New(code uint32, message, reason string, internalCode uint32) *TrError {
func New(code int32, message string) *TrError {
	return &TrError{
		Code:    code,
		Message: message,
		//Reason:       reason,
		//Metadata:     metadata,
	}
}
