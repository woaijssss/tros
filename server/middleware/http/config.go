package http

import trlogger "github.com/woaijssss/tros/logx"

// HTTPConfig http日志打印配置
type HTTPConfig struct {
	// Logger logx.Logger instance
	Logger *trlogger.TrLogger
	// Excludes ignore paths
	Excludes []string
}
