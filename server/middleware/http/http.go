package http

import (
	"bytes"
	"gitee.com/idigpower/tros/context"
	trlogger "gitee.com/idigpower/tros/logx"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func AddTraceID(ctx *gin.Context) {
	body, err := ctx.GetRawData()
	if err != nil {
		trlogger.Error(ctx, err.Error())
	}
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	c := context.InsertTraceID(ctx)

	//输入参数写入日志
	trlogger.Infof(c, "smart_request_in: RequestURI[%s],RequestBody[%+v]", ctx.Request.Header, string(body))
}
