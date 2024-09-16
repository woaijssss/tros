package http

import (
	"github.com/gin-gonic/gin"
	trlogger "github.com/woaijssss/tros/logx"
	"net/http"
)

func HeartCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.RequestURI == "/heart" {
			trlogger.Infof(c, "health check")
			c.AbortWithStatusJSON(http.StatusOK, "ok")
		}
	}
}
