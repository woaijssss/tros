package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"github.com/woaijssss/tros/server/middleware"
	codes2 "github.com/woaijssss/tros/trerror/codes"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"io"
	"net/http"
	"strings"
)

const (
	headerTrailer          = "Trailer"
	headerTransferEncoding = "Transfer-Encoding"
	headerContentType      = "Content-Type"
	defaultErrCode         = "0x000001F4"
	// PathPatternKey HTTP path pattern key
	PathPatternKey = "pattern"
	maxRecvMsgSize = 20 * 1024 * 1024
)

const (
	marshalError = "{\"code\":\"0x000001F4\",\"message\":\"failed to marshal error message\"}"
)

// malformedHTTPHeaders lists the headers that the gRPC server may reject outright as malformed.
// See https://github.com/grpc/grpc-go/pull/4803#issuecomment-986093310 for more context.
// keep the same to https://github.com/grpc-ecosystem/grpc-gateway/blob/master/runtime/context.go#L46 malformedHTTPHeaders
var malformedHTTPHeaders = map[string]struct{}{
	"connection": {},
}

type wrapError struct {
	Code    int32         `json:"code"`
	Message string        `json:"message"`
	Details []interface{} `json:"details"`
}

func attachGRpcGateway(ctx context.Context, s *Server) error {
	conn, err := grpc.DialContext(
		ctx,
		s.gRPCServerAddress,
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxRecvMsgSize)),
	)
	if err != nil {
		logrus.Error("failed to dial server", "error", err)
		return err
	}

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	matcher := httpHeaderMatcher

	// 使用下划线方式接收请求参数
	marshaler := &runtime.HTTPBodyMarshaler{
		Marshaler: &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		},
	}

	gux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(matcher),
		runtime.WithErrorHandler(gwErrorHandler),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, marshaler),
		runtime.WithMetadata(func(ctx context.Context, _ *http.Request) metadata.MD {
			md := make(map[string]string)
			if pattern, ok := runtime.HTTPPathPattern(ctx); ok {
				md[PathPatternKey] = pattern
			}
			return metadata.New(md)
		}),
	)

	for _, r := range s.handlers {
		if err := r(ctx, gux, conn); err != nil {
			return err
		}
	}

	s.NoRoute(forwardHandler(gux))
	return nil
}

// 自定义错误匹配器，可以根据需要调整
func customHeaderMatcher(key string) (string, bool) {
	// 根据需要添加自定义的gRPC头到HTTP响应中
	switch key {
	case "header-key1":
		return "grpc-header-key1", true
	case "header-key2":
		return "grpc-header-key2", true
	default:
		return "", false
	}
}

// 自定义HTTP错误处理函数
func customErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	// 检查错误是否是gRPC错误
	if s, ok := status.FromError(err); ok {
		fmt.Println(s)
		// 设置自定义HTTP状态码
		w.WriteHeader(http.StatusBadRequest)
		// 返回自定义错误信息
		w.Write([]byte(`{"error": "custom error message"}`))
		return
	}
	// 默认处理错误...
}

func forwardHandler(gux *runtime.ServeMux) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 预设成功状态, 覆盖默认的404状态，如果异常在gux内部会再次修改
		c.Writer.WriteHeader(http.StatusOK)
		middleware.MarkRequestFromGRpcGateway(c.Request)
		gux.ServeHTTP(c.Writer, c.Request)
	}
}

func httpHeaderMatcher(key string) (string, bool) {
	// Allow all header key, except malformedHTTPHeaders
	for malformedHTTPHeader := range malformedHTTPHeaders {
		if malformedHTTPHeader == strings.ToLower(key) {
			return key, false
		}
	}
	return key, true
}

//func gwErrorHandler(_ context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
//	var (
//		custom     *runtime.HTTPStatusError
//		ns         *status.Status
//		code       uint32
//		customCode string
//		buf        []byte
//		marshalErr error
//	)
//	if errors.As(err, &custom) {
//		err = custom.Err
//	}
//
//	s := status.Convert(err) // s为应用返回的 TrError
//	message := s.Message()
//	originCode := s.ErrCode()
//	code = codes2.GRPCToCode(s.ErrCode())
//
//	if len(s.Details()) == 0 {
//		// 全局异常details为空
//		ns = status.New(codes.ErrCode(code), codes2.CodeAbstract(code))
//		ns, err = ns.WithDetails(&epb.ErrorInfo{
//			Reason: message,
//		})
//		if err != nil {
//			grpclog.Errorf("Failed to ns.WithDetails %q: %v", s, err)
//		}
//	} else {
//		// 通用异常details不为空
//		ns = status.New(codes.ErrCode(code), message)
//		for _, detail := range s.Details() {
//			if vv, ok := detail.(proto.Message); ok {
//				ns, err = ns.WithDetails(vv)
//				if err != nil {
//					grpclog.Errorf("Failed to ns.WithDetails %q: %v", s, err)
//				}
//			}
//		}
//	}
//
//	pb := ns.Proto()
//
//	w.Header().Del(headerTrailer)
//	w.Header().Del(headerTransferEncoding)
//	w.Header().Set(headerContentType, marshaler.ContentType(pb))
//	wrapDetails, customCode := buildWrapDetail(s.Details())
//	if len(customCode) == 0 {
//		customCode = defaultErrCode
//	}
//	resError := &wrapError{
//		ErrCode: customCode,
//		Message: message,
//		Details: wrapDetails,
//	}
//	buf, marshalErr = marshaler.Marshal(resError)
//	if marshalErr != nil {
//		grpclog.Infof("Failed to marshal error message %q: %v", s, marshalErr)
//		w.WriteHeader(http.StatusInternalServerError)
//		if _, err := io.WriteString(w, marshalError); err != nil {
//			grpclog.Infof("Failed to write response: %v", err)
//		}
//		return
//	}
//
//	st := runtime.HTTPStatusFromCode(originCode)
//	if custom != nil {
//		st = custom.HTTPStatus
//	}
//
//	w.WriteHeader(st)
//	if _, err := w.Write(buf); err != nil {
//		grpclog.Infof("Failed to write response: %v", err)
//	}
//}

func gwErrorHandler(ctx context.Context, rs *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	var (
		custom *runtime.HTTPStatusError
		//ns         *status.Status
		//code       uint32
		buf        []byte
		marshalErr error
	)
	if errors.As(err, &custom) {
		err = custom.Err
	}

	s := status.Convert(err)   // s为应用返回的 TrError
	appMessage := s.Message()  // s为应用返回的 TrError 的 Message
	appCode := int32(s.Code()) // s为应用返回的 TrError 的 Code

	// 转换为对应的http错误码
	//code := codes2.GRPCToCode(appCode)

	resError := &wrapError{
		Code:    appCode,
		Message: appMessage,
		//Details: wrapDetails,
	}
	buf, marshalErr = marshaler.Marshal(resError)
	if marshalErr != nil {
		grpclog.Infof("Failed to marshal error message %q: %v", s, marshalErr)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := io.WriteString(w, marshalError); err != nil {
			grpclog.Infof("Failed to write response: %v", err)
		}
		return
	}

	statusCode := codes2.AppCodeToHttpStatus(appCode)
	//statusCode := runtime.HTTPStatusFromCode(codes.Code(st))
	w.WriteHeader(int(statusCode))
	if _, err := w.Write(buf); err != nil {
		grpclog.Infof("Failed to write response: %v", err)
	}
}

// build http error detail
func buildWrapDetail(details []interface{}) ([]interface{}, string) {
	wrapDetails := make([]interface{}, len(details))
	customCode := ""
	copy(wrapDetails, details)
	if len(wrapDetails) > 0 {
		detail := wrapDetails[0]
		if de, ok := detail.(*epb.ErrorInfo); ok {
			if de.Metadata != nil {
				customCode = de.Metadata["code"]
				delete(de.Metadata, "code")
			}
		}
	}
	return wrapDetails, customCode
}
