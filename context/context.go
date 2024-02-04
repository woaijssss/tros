package context

import (
	"context"
	"gitee.com/idigpower/tros/constants"
	"gitee.com/idigpower/tros/pkg/utils/encrypt"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
)

type TrContext struct {
	context.Context
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

const (
	headerRealIP    = "X-Real-Ip"
	headerUserAgent = "User-Agent"
	headerRefer     = "Referer"
	metadataTrace   = "Tr-Trace"
)

func (ctx *TrContext) SetExtraKeyValue(key string, val any) {
	if ctx.Extra == nil {
		ctx.Extra = map[string]any{key: val}
	} else {
		ctx.Extra[key] = val
	}
}

func (ctx *TrContext) GetExtraValue(key string) any {
	if ctx.Extra == nil {
		return nil
	} else {
		return ctx.Extra[key]
	}
}

func GenTraceID(ctx context.Context) string {
	newCtx, ok := ctx.(*gin.Context)
	if ok && newCtx.Request != nil {
		ctx = newCtx.Request.Context()
	}
	return encrypt.EncodeMD5AsEmpty()
}

func GetContextWithTraceId() context.Context {
	ctx := context.Background()
	return injectToContext(ctx, GenTraceID(ctx))
}

func GetTraceIdFromContext(ctx context.Context) string {
	traceId, ok := ctx.Value(constants.TraceId).(string)
	if !ok {
		return ""
	}

	return traceId
}

// FromContext get trace entry from context.
func FromContext(ctx context.Context) context.Context {
	traceId := GetTraceIdFromContext(ctx)
	if traceId == "" {
		return GetContextWithTraceId()
	}
	return injectToContext(context.Background(), traceId)
}

func InsertTraceID(ctx context.Context) context.Context {
	md := metautils.ExtractIncoming(ctx)
	s := md.Get(metadataTrace)
	if s == "" {
		s = GenTraceID(ctx)
	}

	return injectToContext(ctx, s)
}

func injectToContext(ctx context.Context, traceId string) context.Context {
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
