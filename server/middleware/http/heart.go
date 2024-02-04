package http

import (
	trlogger "gitee.com/idigpower/tros/logx"
	"github.com/gin-gonic/gin"
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
