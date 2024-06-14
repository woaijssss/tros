package trerror

import (
	"gitee.com/idigpower/tros/trerror/codes"
)

// Canceled returns Error for canceled
func Canceled(reason string, args ...interface{}) *TrError {
	return newByStdCode(codes.Canceled, reason, args...)
}

// Unknown returns unknown Error
func Unknown(reason string, args ...interface{}) *TrError {
	return newByStdCode(codes.Unknown, reason, args...)
}

// InvalidArgument returns invalid argument Error
func InvalidArgument(reason string, args ...interface{}) *TrError {
	return newByStdCode(codes.InvalidArgument, reason, args...)
}

// DeadlineExceeded returns deadline exceeded Error
func DeadlineExceeded(reason string, args ...interface{}) *TrError {
	return newByStdCode(codes.DeadlineExceeded, reason, args...)
}

// NotFound returns resource not found Error
func NotFound(reason string, args ...interface{}) *TrError {
	return newByStdCode(codes.NotFound, reason, args...)
}

// AlreadyExists returns resource already exists Error
func AlreadyExists(reason string, args ...interface{}) *TrError {
	return newByStdCode(codes.AlreadyExists, reason, args...)
}

// PermissionDenied returns permission denied Error
func PermissionDenied(reason string, args ...interface{}) *TrError {
	return newByStdCode(codes.PermissionDenied, reason, args...)
}

// ResourceExhausted returns resource exhaust Error
func ResourceExhausted(reason string, args ...interface{}) *TrError {
	return newByStdCode(codes.ResourceExhausted, reason, args...)
}

// Aborted return abort Error
func Aborted(reason string, args ...interface{}) *TrError {
	return newByStdCode(codes.Aborted, reason, args...)
}

// Unimplemented return unimplemented Error
func Unimplemented(reason string, args ...interface{}) *TrError {
	return newByStdCode(codes.Unimplemented, reason, args...)
}

// Internal return inner server Error
func Internal(reason string, args ...interface{}) *TrError {
	return newByStdCode(codes.Internal, reason, args...)
}

// Unavailable return service unavailable Error
func Unavailable(reason string, args ...interface{}) *TrError {
	return newByStdCode(codes.Unavailable, reason, args...)
}

// Unauthenticated returns unauthenticated Error
func Unauthenticated(reason string, args ...interface{}) *TrError {
	return newByStdCode(codes.Unauthenticated, reason, args...)
}

// ConfigurationError returns configuration Error
func ConfigurationError(reason string, args ...interface{}) *TrError {
	return newByStdCode(codes.ConfigurationError, reason, args...)
}

// DBError returns database Error
func DBError(reason string, args ...interface{}) *TrError {
	return newByStdCode(codes.DBError, reason, args...)
}

// CustomError returns custom Error
func CustomError(code uint32, internalCode uint32, message, reason string, args ...interface{}) *TrError {
	return NewErrorWithF(code, internalCode, message, reason, args...)
}

// CustomCanceledError returns Error for canceled
func CustomCanceledError(code uint32, msg, reason string, args ...interface{}) *TrError {
	return NewErrorWithF(code, codes.Canceled, msg, reason, args...)
}

// CustomUnknownError returns unknown Error
func CustomUnknownError(code uint32, msg, reason string, args ...interface{}) *TrError {
	return NewErrorWithF(code, codes.Unknown, msg, reason, args...)
}

// CustomInvalidArgumentError returns invalid argument Error
func CustomInvalidArgumentError(code uint32, msg, reason string, args ...interface{}) *TrError {
	return NewErrorWithF(code, codes.InvalidArgument, msg, reason, args...)
}

// CustomDeadlineExceededError returns deadline exceeded Error
func CustomDeadlineExceededError(code uint32, msg, reason string, args ...interface{}) *TrError {
	return NewErrorWithF(code, codes.DeadlineExceeded, msg, reason, args...)
}

// CustomNotFoundError returns resource not found Error
func CustomNotFoundError(code uint32, msg, reason string, args ...interface{}) *TrError {
	return NewErrorWithF(code, codes.NotFound, msg, reason, args...)
}

// CustomAlreadyExistsError returns resource already exists Error
func CustomAlreadyExistsError(code uint32, msg, reason string, args ...interface{}) *TrError {
	return NewErrorWithF(code, codes.AlreadyExists, msg, reason, args...)
}

// CustomPermissionDeniedError returns permission denied Error
func CustomPermissionDeniedError(code uint32, msg, reason string, args ...interface{}) *TrError {
	return NewErrorWithF(code, codes.PermissionDenied, msg, reason, args...)
}

// CustomResourceExhaustedError returns resource exhaust Error
func CustomResourceExhaustedError(code uint32, msg, reason string, args ...interface{}) *TrError {
	return NewErrorWithF(code, codes.ResourceExhausted, msg, reason, args...)
}

// CustomAbortedError return abort Error
func CustomAbortedError(code uint32, msg, reason string, args ...interface{}) *TrError {
	return NewErrorWithF(code, codes.Aborted, msg, reason, args...)
}

// CustomUnimplementedError return unimplemented Error
func CustomUnimplementedError(code uint32, msg, reason string, args ...interface{}) *TrError {
	return NewErrorWithF(code, codes.Unimplemented, msg, reason, args...)
}

// CustomInternalError return inner server Error
func CustomInternalError(code uint32, msg, reason string, args ...interface{}) *TrError {
	return NewErrorWithF(code, codes.Internal, msg, reason, args...)
}

// CustomUnavailableError return service unavailable Error
func CustomUnavailableError(code uint32, msg, reason string, args ...interface{}) *TrError {
	return NewErrorWithF(code, codes.Unavailable, msg, reason, args...)
}

// CustomUnauthenticatedError returns unauthenticated Error
func CustomUnauthenticatedError(code uint32, msg, reason string, args ...interface{}) *TrError {
	return NewErrorWithF(code, codes.Unauthenticated, msg, reason, args...)
}

// CustomConfigurationError returns configuration Error
func CustomConfigurationError(code uint32, msg, reason string, args ...interface{}) *TrError {
	return NewErrorWithF(code, codes.ConfigurationError, msg, reason, args...)
}

// CustomDBError returns database Error
func CustomDBError(code uint32, msg, reason string, args ...interface{}) *TrError {
	return NewErrorWithF(code, codes.DBError, msg, reason, args...)
}

func newByStdCode(code uint32, reason string, args ...interface{}) *TrError {
	msg := codes.CodeAbstract(code)
	return NewF(code, msg, reason, args...)
}

func newByCode(code uint32, internalCode uint32, reason string, args ...interface{}) *TrError {
	msg := codes.CodeAbstract(internalCode)
	return NewErrorWithF(code, internalCode, msg, reason, args...)
}
