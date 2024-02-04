package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(myRecover)
}

func myRecover(c *gin.Context, err any) {
	//	// todo 打印一条日志
	//
	// 封装通用json结果返回
	c.JSON(http.StatusOK, err)
	// 终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
	c.Abort()
}
