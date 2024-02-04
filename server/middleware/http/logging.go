package http

import (
	trlogger "gitee.com/idigpower/tros/logx"
	"gitee.com/idigpower/tros/server/middleware"
	"github.com/gin-gonic/gin"
)

// HTTPLogger new http logger middleware
//func HTTPLogger() gin.HandlerFunc {
//	return HTTPLoggerWithConfig(HTTPConfig{
//		Excludes: []string{
//			healthz.PathHealthz,
//			healthz.PathReady,
//		},
//	})
//}

// HTTPLoggerWithConfig new logger middleware with LoggerConfig
func HTTPLoggerWithConfig(config HTTPConfig) gin.HandlerFunc {
	log := config.Logger
	if log == nil {
		log = trlogger.DefaultTrLogger()
	}

	hl := &httpLogger{
		logger:   log,
		excludes: middleware.ExcludePaths(config.Excludes),
	}

	return hl.handle
}

type httpLogger struct {
	logger   *trlogger.TrLogger
	excludes map[string]struct{}
}

func (hl *httpLogger) handle(c *gin.Context) {
	// 忽略打印
	if _, ok := hl.excludes[c.Request.URL.Path]; ok {
		c.Next()
		return
	}

	//start := time.Now()
	c.Next()
	//hl.logger.Infof(c, "http request", hl.kv(c, start)...)
}

//func (hl *httpLogger) kv(c *gin.Context, start time.Time) []interface{} {
//	kv := utils.GetLogKv(c)
//	traceID := context.GenTraceID(c)
//	kv = append(kv,
//		"cost", time.Since(start),
//		constants.TraceId, traceID,
//	)
//
//	return kv
//}
