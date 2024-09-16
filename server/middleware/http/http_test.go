package http

import (
	trlogger "github.com/woaijssss/tros/logx"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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

func TestHTTP(t *testing.T) {
	adapter := &traceLogAdapter{}

	router := gin.New()
	router.Use(HTTPLoggerWithConfig(HTTPConfig{
		Logger: trlogger.DefaultTrLogger(),
	}))
	router.GET("/healthz", func(c *gin.Context) {
		c.String(200, "echo")
	})

	req, _ := http.NewRequest(http.MethodGet, "/healthz", nil)
	router.ServeHTTP(httptest.NewRecorder(), req)

	assert.Equal(t, 1, len(adapter.records))

	record := adapter.records[0]
	assert.Equal(t, trlogger.LevelInfo, record.level)
}

func TestHTTPExclude(t *testing.T) {
	adapter := &traceLogAdapter{}

	router := gin.New()
	router.Use(HTTPLoggerWithConfig(HTTPConfig{
		Logger:   trlogger.DefaultTrLogger(),
		Excludes: []string{"/echo"},
	}))
	router.GET("/echo", func(c *gin.Context) {
		c.String(200, "echo")
	})

	req, _ := http.NewRequest(http.MethodGet, "/echo", nil)
	router.ServeHTTP(httptest.NewRecorder(), req)

	assert.Equal(t, 0, len(adapter.records))
}
