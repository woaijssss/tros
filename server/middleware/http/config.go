package http

import trlogger "gitee.com/idigpower/tros/logx"

// HTTPConfig http日志打印配置
type HTTPConfig struct {
	// Logger logx.Logger instance
	Logger *trlogger.TrLogger
	// Excludes ignore paths
	Excludes []string
}
