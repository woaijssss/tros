package grpc

import (
	"context"
	go_grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/woaijssss/tros/conf"
	trlogger "github.com/woaijssss/tros/logx"
	grpc3 "github.com/woaijssss/tros/server/middleware/grpc"
	"github.com/woaijssss/tros/server/middleware/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net"
	"strconv"
)

// DefaultServer 默认的grpc server入口
func DefaultServer(opts ...ServerOption) *Server {
	var options []ServerOption

	lc := grpc3.GRpcConfig{
		ExcludeGRpcGatewayRequest: true,
	}
	ro := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandlerContext(recoveryHandler),
	}

	options = append(
		options,
		UnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(ro...),
			//openTelemeTryMiddleware.UnaryServerInterceptorrceptor(),
			grpc3.UnaryServerInterceptor(lc),
			//grpc_recovery.UnaryServerInterceptor(ro...),
			//identity.UnaryServerInterceptor(),
			http.UnaryServerInterceptor(),
		),
		StreamInterceptor(
			grpc_recovery.StreamServerInterceptor(ro...),
			//openTelemeTryMiddleware.StreamServerInterceptor(),
			//logging.StreamServerInterceptor(lc),
			grpc_recovery.StreamServerInterceptor(ro...),
			//identity.StreamServerInterceptor(),
			//tracing.StreamServerInterceptor(),
		),
	)
	options = append(options, opts...)

	return NewServer(options...)
}

// NewServer new server with options
func NewServer(opts ...ServerOption) *Server {
	addr := ":" + strconv.Itoa(conf.GetGrpcPort())
	server := &Server{
		address: addr,
		//log:     logx.WithoutContext(),
	}

	for _, o := range opts {
		o(server)
	}

	var list []grpc.ServerOption
	if len(server.unaryInterceptors) > 0 {
		list = append(list, grpc.UnaryInterceptor(go_grpc_middleware.ChainUnaryServer(server.unaryInterceptors...)))
	}

	if len(server.streamInterceptors) > 0 {
		list = append(list, grpc.StreamInterceptor(go_grpc_middleware.ChainStreamServer(server.streamInterceptors...)))
	}

	list = append(list, server.grpcOpts...)
	server.gs = grpc.NewServer(list...)

	// register builtin service
	reflection.Register(server.gs)

	return server
}

// ServerOption gRpc server option
type ServerOption func(server *Server)

// Address gRpc server listen address
func GrpcAddress(address string) ServerOption {
	return func(s *Server) {
		s.address = address
	}
}

// Listener listener
func GrpcListener(lis net.Listener) ServerOption {
	return func(s *Server) {
		s.lis = lis
	}
}

// Options set gRpc server options
func GrpcOptions(opts ...grpc.ServerOption) ServerOption {
	return func(s *Server) {
		s.grpcOpts = opts
	}
}

// StreamInterceptor register stream interceptors
func StreamInterceptor(in ...grpc.StreamServerInterceptor) ServerOption {
	return func(s *Server) {
		s.streamInterceptors = append(s.streamInterceptors, in...)
	}
}

// Server of gRpc
type Server struct {
	gs                 *grpc.Server
	address            string
	grpcOpts           []grpc.ServerOption
	unaryInterceptors  []grpc.UnaryServerInterceptor
	streamInterceptors []grpc.StreamServerInterceptor
	//log                logx.Logger
	lis net.Listener
}

// Start gRpc server
func (s *Server) Start(ctx context.Context) error {
	trlogger.Infof(ctx, "[GRPC] Start")
	if s.lis == nil {
		lis, err := net.Listen("tcp", s.address)
		if err != nil {
			return err
		}
		s.lis = lis
	}
	trlogger.Infof(ctx, "[GRPC] server listening address %s", s.lis.Addr().String())
	return s.gs.Serve(s.lis)
}

// Stop gRpc server
func (s *Server) Stop() error {
	s.gs.GracefulStop()
	trlogger.Infof(context.Background(), "[GRPC] server stopping")
	return nil
}

func recoveryHandler(ctx context.Context, p interface{}) (err error) {
	trlogger.Errorf(ctx, "grpc server panic recovery panic: [%+v]", p)
	return status.Errorf(codes.Unknown, "panic triggered: %v", p)
}

// UnaryInterceptor register unary interceptors
func UnaryInterceptor(in ...grpc.UnaryServerInterceptor) ServerOption {
	return func(s *Server) {
		s.unaryInterceptors = append(s.unaryInterceptors, in...)
	}
}

// RegisterService register gRpc service
func (s *Server) RegisterService(desc *grpc.ServiceDesc, impl interface{}) {
	s.gs.RegisterService(desc, impl)
}

// GetListener get servers listener
func (s *Server) GetListener() net.Listener {
	return s.lis
}
