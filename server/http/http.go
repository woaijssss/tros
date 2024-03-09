package http

import (
	"context"
	"errors"
	"fmt"
	"gitee.com/idigpower/tros/conf"
	trlogger "gitee.com/idigpower/tros/logx"
	http4 "gitee.com/idigpower/tros/server/middleware/http"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

// DefaultServer 默认的http server入口
func DefaultServer(opts ...ServerOption) *Server {
	s := NewServer(opts...)

	s.Use(
		http4.Recovery(),
		http4.Cors(),
		http4.Monitor(),
		http4.HeartCheck(),
	)

	return s
}

// NewServer new server with options
func NewServer(opts ...ServerOption) *Server {
	e := NewEngine(
		http4.Recovery(),
		http4.Cors(),
		http4.Monitor(),
		http4.HeartCheck(),
	)

	AddMiddleWares(e,
		//gin.Logger(),
		sentrygin.New(sentrygin.Options{Repanic: true}),
		http4.AddTraceID,
		//http4.AddUserID,
	)
	setRuntimeMode(e)
	server := &Server{
		Engine:            gin.New(),
		address:           fmt.Sprintf(":%d", conf.GetHttpPort()),
		gRPCServerAddress: ":" + strconv.Itoa(conf.GetGrpcPort()),
		//log:               logx.WithoutContext(),
	}
	server.Engine = e

	for _, o := range opts {
		o(server)
	}

	return server
}

var trustProxies = []string{
	"0.0.0.0/0",
}

type (
	// ServerOption options for http server
	ServerOption func(*Server)
	// ServiceHandler gRpc-Gateway handler
	ServiceHandler func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error
)

// Router http router
type Router interface {
	gin.IRouter
	RegisterServiceHandler(handler ServiceHandler)
}

// Listener listener
func Listener(lis net.Listener) ServerOption {
	return func(s *Server) {
		s.lis = lis
	}
}

// Address listen address
func Address(address string) ServerOption {
	return func(s *Server) {
		s.address = address
	}
}

// GRpcServerAddress gRpc server address for gRpc-Gateway
func GRpcServerAddress(address string) ServerOption {
	return func(s *Server) {
		s.gRPCServerAddress = address
	}
}

// Server http server
type Server struct {
	*gin.Engine

	handlers          []ServiceHandler
	address           string
	gRPCServerAddress string

	hs  *http.Server
	mu  sync.Mutex
	lis net.Listener

	//log logx.Logger
}

func setRuntimeMode(e *gin.Engine) {
	// gin mode
	if trlogger.IsDebugLevel() {
		gin.SetMode(gin.DebugMode)
		AddMiddleWares(e, gin.Logger())
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

// Start start http server
func (s *Server) Start(ctx context.Context) error {
	trlogger.Infof(ctx, "[HTTP] Start")

	err := s.Engine.SetTrustedProxies(trustProxies)
	if err != nil {
		return fmt.Errorf("set gin trusted proxies failed: %v", err)
	}

	if len(s.handlers) > 0 {
		err = attachGRpcGateway(ctx, s)
		if err != nil {
			return err
		}
	}

	// use http2 to adapt agent request
	s.mu.Lock()
	s.hs = &http.Server{
		Handler: h2c.NewHandler(s.Engine, &http2.Server{}),
	}
	s.mu.Unlock()

	if s.lis == nil {
		if s.lis, err = net.Listen("tcp", s.address); err != nil {
			return err
		}
	}
	trlogger.Infof(ctx, "[HTTP] server listening address %s", s.lis.Addr().String())
	if err := s.hs.Serve(s.lis); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Stop stop http server
func (s *Server) Stop() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	trlogger.Infof(context.Background(), "[HTTP] server stopping")
	if s.hs != nil {
		return s.hs.Shutdown(context.Background())
	}

	return nil
}

// RegisterServiceHandler register gRpc-gateway handler
func (s *Server) RegisterServiceHandler(handler ServiceHandler) {
	s.handlers = append(s.handlers, handler)
}

// GetListener get servers listener
func (s *Server) GetListener() net.Listener {
	return s.lis
}
