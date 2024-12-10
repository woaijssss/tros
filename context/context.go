package context

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/woaijssss/tros/constants"
	"github.com/woaijssss/tros/pkg/utils/encrypt"
	"google.golang.org/grpc/metadata"
	"strings"
	"time"
)

type TrContext struct {
	//context.Context
	TraceId    string         `json:"traceId,omitempty"`
	UserId     int64          `json:"userId,omitempty"`
	OpId       int64          `json:"opId,omitempty"`
	RunAs      int64          `json:"runAs,omitempty"`
	Roles      string         `json:"roles,omitempty"`
	BizTypes   int            `json:"bizTypes,omitempty"`
	GroupId    int64          `json:"groupId,omitempty"`
	Platform   string         `json:"platform,omitempty"`
	UserAgent  string         `json:"userAgent,omitempty"`
	Lang       string         `json:"lang,omitempty"`
	GoId       uint64         `json:"goId,omitempty"`
	Token      string         `json:"token,omitempty"`
	ShareToken string         `json:"shareToken,omitempty"`
	RemoteIp   string         `json:"remoteIp,omitempty"`
	CompanyId  int64          `json:"companyId,omitempty"`
	Product    int            `json:"product,omitempty"`
	Extra      map[string]any `json:"extra,omitempty"`
}

func (tc *TrContext) Deadline() (time.Time, bool)       { return time.Time{}, false }
func (tc *TrContext) Done() <-chan struct{}             { return nil }
func (tc *TrContext) Err() error                        { return nil }
func (tc *TrContext) Value(key interface{}) interface{} { return tc.Value(key) }

const (
	headerRealIP    = "X-Real-Ip"
	headerUserAgent = "User-Agent"
	headerRefer     = "Referer"
	metadataTrace   = "Tr-Trace"
)

func (tc *TrContext) SetExtraKeyValue(key string, val any) {
	if tc.Extra == nil {
		tc.Extra = map[string]any{key: val}
	} else {
		tc.Extra[key] = val
	}
}

func (tc *TrContext) GetExtraValue(key string) any {
	if tc.Extra == nil {
		return nil
	} else {
		return tc.Extra[key]
	}
}

func GenTraceID() string {
	return encrypt.EncodeMD5AsEmpty()
}

func GetIpFromCtx(ctx context.Context) string {
	ip, ok := ctx.Value(constants.RemoteIp).(string)
	if !ok {
		return ""
	}

	return ip
}

func GetContextWithTraceId() context.Context {
	ctx := context.Background()
	return AddTraceID(ctx, GenTraceID())
}

func GetTraceIdFromContext(ctx context.Context) string {
	traceId, ok := ctx.Value(constants.TraceId).(string)
	if !ok {
		return ""
	}

	return traceId
}

func GetUserIdFromContext(ctx context.Context) string {
	if c, ok := ctx.(*gin.Context); ok {
		return c.GetString(constants.UserId)
	}

	return ctx.Value(constants.UserId).(string)
}

func InsertTraceID(ctx context.Context) context.Context {
	md := metautils.ExtractIncoming(ctx)
	s := md.Get(metadataTrace)
	if s == "" {
		s = GenTraceID()
	}

	return AddTraceID(ctx, s)
}

func InsertRemoteIp(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if ipHeaders, ok := md["x-forwarded-for"]; ok && len(ipHeaders) > 0 {
			/*
				The IP passed in by grpc gateway may still have "127.0.0.1" or other IPs after the real IP, so a replacement is needed
			*/
			ips := strings.Split(ipHeaders[0], ",")
			ip := ""
			if len(ips) >= 1 {
				ip = ips[0]
			}
			fmt.Println("Remote IP:", ip)
			return context.WithValue(ctx, constants.RemoteIp, ip)
		}
	}
	if c, ok := ctx.(*gin.Context); ok {
		c.Set(constants.RemoteIp, c.RemoteIP())
	}
	return ctx
}

func AddTraceID(ctx context.Context, traceId string) context.Context {
	if traceId == "" {
		return ctx
	}
	if c, ok := ctx.(*gin.Context); ok {
		c.Set(constants.TraceId, traceId)
		if c.Request != nil {
			newCtx := context.WithValue(c.Request.Context(), constants.TraceId, traceId)
			c.Request = c.Request.WithContext(newCtx)
			c.Request.Header.Set(metadataTrace, traceId)
		}
		return c
	}
	return context.WithValue(ctx, constants.TraceId, traceId)
}

func AddUserID(ctx context.Context, userId string) context.Context {
	if c, ok := ctx.(*gin.Context); ok { // 兼容纯gin模式的请求
		c.Set(constants.UserId, userId)
		if c.Request != nil {
			newCtx := context.WithValue(c.Request.Context(), constants.UserId, userId)
			c.Request = c.Request.WithContext(newCtx)
		}
		return c
	}
	return context.WithValue(ctx, constants.UserId, userId) // 兼容grpc-gateway模式的请求
}
