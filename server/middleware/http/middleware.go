package http

import (
	"context"
	"errors"
	"github.com/gin-contrib/cors"
	trlogger "github.com/woaijssss/tros/logx"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var AllowOrigins = []string{
	"*",
}

var AllowHeaders = []string{
	"device-id",
	"hardware",
	"tros",
	"os",
	"os_version",
	"location",
	"ip",
	"network_type",
	"timestamp",
	"user_agent",
	"resolution",
	"platform",
	"app_key",
	"app_version",
	"app_vsn",
	"trace_id",
	"token",
	"s-token",
	"run-as",
	"company-id",
	"product",
	"X-Forwarded-For",
	"X-Forwarded-Proto",
	"Authorization",
}

func Timeout(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		// wrap the request context with a timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer func() {
			// check if context timeout was reached
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				// write response and abort the request
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				c.Abort()
			}
			//cancel to clear resources after finished
			cancel()
		}()
		// replace request with context wrapped request
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// 解决前端跨域问题
func Cors() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = AllowOrigins
	corsConfig.AllowHeaders = AllowHeaders

	return cors.New(corsConfig)
}

// Metric middleware
func Metric() gin.HandlerFunc {
	return func(c *gin.Context) {
		trlogger.Infof(c, "Metric start")
		tBegin := time.Now()
		c.Next()
		trlogger.Infof(c, "Metric end")
		duration := float64(time.Since(tBegin)) / float64(time.Second)
		path := c.Request.URL.Path
		trlogger.Infof(c, "uri=[%s] duration=[%f]", path, duration)
	}
}
