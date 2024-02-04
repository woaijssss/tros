package grpc

import (
	trlogger "gitee.com/idigpower/tros/logx"
	"sync"
	"testing"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_testing "github.com/grpc-ecosystem/go-grpc-middleware/testing"
	mwitkow_testproto "github.com/grpc-ecosystem/go-grpc-middleware/testing/testproto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

type traceLogAdapter struct {
	records []logRecord
	mu      sync.RWMutex
}

type logRecord struct {
	level trlogger.Level
	msg   string
	kv    []interface{}
}

type gRpcTestSuite struct {
	adapter *traceLogAdapter
	*grpc_testing.InterceptorTestSuite
}

func TestGRpcSuite(t *testing.T) {
	adapter := &traceLogAdapter{}
	config := GRpcConfig{
		Logger:                    trlogger.DefaultTrLogger(),
		Excludes:                  []string{"/mwitkow.testproto.TestService/PingEmpty"},
		ExcludeGRpcGatewayRequest: true,
	}

	s := &gRpcTestSuite{
		adapter: adapter,
		InterceptorTestSuite: &grpc_testing.InterceptorTestSuite{
			TestService: &grpc_testing.TestPingService{T: t},
			ServerOpts: []grpc.ServerOption{
				grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
					UnaryServerInterceptor(config),
				)),
				grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
					StreamServerInterceptor(config),
				)),
			},
		},
	}

	suite.Run(t, s)
}

func (s *gRpcTestSuite) SetupTest() {
	s.adapter.mu.Lock()
	defer s.adapter.mu.Unlock()
	s.adapter.records = nil
}

func (s *gRpcTestSuite) TestUnary() {
	_, err := s.Client.Ping(s.SimpleCtx(), &mwitkow_testproto.PingRequest{})
	assert.Nil(s.T(), err)
	// assert.Equal(s.T(), 1, len(s.adapter.records))
}

func (s *gRpcTestSuite) TestExclude() {
	_, err := s.Client.PingEmpty(s.SimpleCtx(), &mwitkow_testproto.Empty{})
	assert.Nil(s.T(), err)
	// assert.Equal(s.T(), 0, len(s.adapter.records))
}

func (s *gRpcTestSuite) TestStream() {
	cli, err := s.Client.PingList(s.SimpleCtx(), &mwitkow_testproto.PingRequest{})
	assert.NotNil(s.T(), cli)

	_, err = cli.Recv()
	assert.Nil(s.T(), err)
	// assert.Equal(s.T(), 1, len(s.adapter.records))
}
