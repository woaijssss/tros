package grpc

import trlogger "github.com/woaijssss/tros/logx"

// GRpcConfig gRpc日志打印配置
type GRpcConfig struct {
	// Logger logx.Logger instance
	Logger *trlogger.TrLogger
	// Excludes ignore paths
	Excludes []string
	// ExcludeGRpcGatewayRequest ignore gRpc-Gateway request
	ExcludeGRpcGatewayRequest bool
}
