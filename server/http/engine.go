package http

import (
	"github.com/gin-gonic/gin"
	"github.com/woaijssss/tros/conf"
	http2 "github.com/woaijssss/tros/server/middleware/http"
)

func DefaultEngine() *gin.Engine {
	return NewEngine(http2.Recovery(), http2.Cors(), http2.Monitor())
}

func NewEngine(middlewares ...gin.HandlerFunc) *gin.Engine {
	if conf.IsProd() {
		gin.SetMode(gin.ReleaseMode)
	}
	e := gin.New()
	e.UseH2C = true
	e.MaxMultipartMemory = 8 << 20
	e.Use(middlewares...)

	return e
}

func AddMiddleWares(e *gin.Engine, middlewares ...gin.HandlerFunc) {
	e.Use(middlewares...)
}
