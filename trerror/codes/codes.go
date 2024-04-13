// Package codes provide codes list
package codes

import (
	"gitee.com/idigpower/tros/lang"
	"net/http"

	"google.golang.org/grpc/codes"
)

// Code define
const (
	OK uint32 = 0

	Canceled          uint32 = 10001
	Unknown           uint32 = 10002
	InvalidArgument   uint32 = 10003
	DeadlineExceeded  uint32 = 10004
	NotFound          uint32 = 10005
	AlreadyExists     uint32 = 10006
	PermissionDenied  uint32 = 10007
	ResourceExhausted uint32 = 10008
	Aborted           uint32 = 10010
	Unimplemented     uint32 = 10012
	Internal          uint32 = 10013
	Unavailable       uint32 = 10014
	Unauthenticated   uint32 = 10016

	ConfigurationError uint32 = 10101
	DBError            uint32 = 10102

	ContentIllegal uint32 = 10201
	ImageIllegal   uint32 = 10202
	VideoIllegal   uint32 = 10203
)

// CodeMap is a mapping for codes and error info
var codeAbstracts = map[uint32]string{
	OK:                 "成功",
	Canceled:           "操作已取消",
	Unknown:            "未知错误",
	InvalidArgument:    "请求参数错误",
	DeadlineExceeded:   "操作已过期",
	NotFound:           "未找到需要的资源",
	AlreadyExists:      "资源已存在",
	PermissionDenied:   "您没有权限进行该操作",
	ResourceExhausted:  "资源已耗尽",
	Aborted:            "操作已中止",
	Unimplemented:      "操作未实现",
	Internal:           "服务器内部错误",
	Unavailable:        "服务暂时不可用",
	Unauthenticated:    "认证失败",
	ConfigurationError: "配置错误",
	DBError:            "数据库错误",
	ContentIllegal:     "文字包含违规信息",
	ImageIllegal:       "图片包含违规信息",
	VideoIllegal:       "视频包含违规信息",
}

// grpcToCode is a mapping between gRPC codes and QT codes
var grpcToCode = map[codes.Code]uint32{
	codes.OK:                 OK,
	codes.Canceled:           Canceled,
	codes.Unknown:            Unknown,
	codes.InvalidArgument:    InvalidArgument,
	codes.DeadlineExceeded:   DeadlineExceeded,
	codes.NotFound:           NotFound,
	codes.AlreadyExists:      AlreadyExists,
	codes.PermissionDenied:   PermissionDenied,
	codes.ResourceExhausted:  ResourceExhausted,
	codes.FailedPrecondition: InvalidArgument,
	codes.Aborted:            Aborted,
	codes.OutOfRange:         InvalidArgument,
	codes.Unimplemented:      Unimplemented,
	codes.Internal:           Internal,
	codes.Unavailable:        Unavailable,
	codes.DataLoss:           Internal,
	codes.Unauthenticated:    Unauthenticated,
}

// codeToGRPC is a mapping between QT codes and gRPC codes
var codeToGRPC = map[uint32]codes.Code{
	OK:                codes.OK,
	Canceled:          codes.Canceled,
	Unknown:           codes.Unknown,
	InvalidArgument:   codes.InvalidArgument,
	DeadlineExceeded:  codes.DeadlineExceeded,
	NotFound:          codes.NotFound,
	AlreadyExists:     codes.AlreadyExists,
	PermissionDenied:  codes.PermissionDenied,
	ResourceExhausted: codes.ResourceExhausted,
	Aborted:           codes.Aborted,
	Unimplemented:     codes.Unimplemented,
	Internal:          codes.Internal,
	Unavailable:       codes.Unavailable,
	Unauthenticated:   codes.Unauthenticated,
}

var codeToStatus = map[uint32]int{
	OK:                http.StatusOK,
	Canceled:          http.StatusRequestTimeout,
	Unknown:           http.StatusInternalServerError,
	InvalidArgument:   http.StatusBadRequest,
	DeadlineExceeded:  http.StatusGatewayTimeout,
	NotFound:          http.StatusNotFound,
	AlreadyExists:     http.StatusConflict,
	PermissionDenied:  http.StatusForbidden,
	ResourceExhausted: http.StatusTooManyRequests,
	Aborted:           http.StatusConflict,
	Unimplemented:     http.StatusNotImplemented,
	Internal:          http.StatusInternalServerError,
	Unavailable:       http.StatusServiceUnavailable,
	Unauthenticated:   http.StatusUnauthorized,
}

// CodeToGRPC convert int codes to codes.Code
func CodeToGRPC(code uint32) codes.Code {
	c, ok := codeToGRPC[code]
	if ok {
		return c
	}
	return codes.Unknown
}

// GRPCToCode convert grpc codes to int codes
func GRPCToCode(code codes.Code) uint32 {
	c, ok := grpcToCode[code]
	if ok {
		return c
	}
	return Unknown
}

// CodeToStatus codes to http status
func CodeToStatus(code uint32) int {
	c, ok := codeToStatus[code]
	if ok {
		return c
	}
	return http.StatusInternalServerError
}

// CodeAbstract codes memo
func CodeAbstract(code uint32) string {
	r, ok := codeAbstracts[code]
	if !ok {
		r = codeAbstracts[Unknown]
	}
	return r
}

// Default Lang, config by cfgkey.GlobalLang
func Default() lang.Lang {
	//str := conf.GetString(cfgkey.GlobalLang)
	//switch str {
	//case "cn":
	//	return CN
	//case "en":
	//	return EN
	//}
	return lang.CN
}
