// Package codes provide codes list
package codes

import (
	"github.com/woaijssss/tros/lang"
	"net/http"
)

// Code define
const (
	OK = 0

	TrSuccess             = 200
	TrBadRequest          = 400
	TrError               = 500
	TrInvalidError        = 4000
	TrInvalidToken        = 4001
	TrFileTooLarge        = 4002
	TrDuplicatePrimaryKey = 4003
	TrLoginError          = 4004
	TrNotLogin            = 4005
	TrUserNotExist        = 4006
	TrDisableUser         = 4007
	TrWrongPassword       = 4008
	TrNoPermission        = 4009
	TrIllegalOperation    = 4010
	TrRecordNotFound      = 4011
	TrLoginUnSupport      = 4012

	TrSystemError = 5001
	TrSystemBusy  = 5002
	TrTimeout     = 5003

	TrContentIllegal = 10201
	TrImageIllegal   = 10202
	TrVideoIllegal   = 10203

	TrDbError      = 80001
	TrDataNotFound = 80002

	TrAccessTooFrequently = 99999
)

var appCodeToHttpStatus = map[int32]int32{
	TrSuccess:             http.StatusOK,
	TrBadRequest:          http.StatusBadRequest,
	TrError:               http.StatusInternalServerError,
	TrInvalidError:        http.StatusInternalServerError,
	TrInvalidToken:        http.StatusForbidden,
	TrFileTooLarge:        http.StatusBadRequest,
	TrDuplicatePrimaryKey: http.StatusBadRequest,
	TrLoginError:          http.StatusForbidden,
	TrNotLogin:            http.StatusForbidden,
	TrUserNotExist:        http.StatusForbidden,
	TrDisableUser:         http.StatusLocked,
	TrWrongPassword:       http.StatusBadRequest,
	TrNoPermission:        http.StatusForbidden,
	TrIllegalOperation:    http.StatusBadRequest,
	TrRecordNotFound:      http.StatusNotFound,
	TrLoginUnSupport:      http.StatusMethodNotAllowed,

	TrSystemError: http.StatusInternalServerError,
	TrSystemBusy:  http.StatusInternalServerError,
	TrTimeout:     http.StatusGatewayTimeout,

	TrContentIllegal: http.StatusBadRequest,
	TrImageIllegal:   http.StatusBadRequest,
	TrVideoIllegal:   http.StatusBadRequest,

	TrDbError:      http.StatusInternalServerError,
	TrDataNotFound: http.StatusNotFound,

	TrAccessTooFrequently: http.StatusTooManyRequests,
}

// AppCodeToHttpStatus app codes to http status
func AppCodeToHttpStatus(code int32) int32 {
	c, ok := appCodeToHttpStatus[code]
	if ok {
		return c
	}
	return http.StatusInternalServerError
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
