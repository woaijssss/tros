package feishu

import (
	"context"
	"fmt"
	"github.com/woaijssss/tros/constants"
	context2 "github.com/woaijssss/tros/context"
)

/*
	内置的飞书消息推送模板
*/

type templateService struct {
}

var TemplateService = new(templateService)

// GetApiExecuteTimeTemplate 接口执行时间模板
func (ts *templateService) GetApiExecuteTimeTemplate(ctx context.Context, tm int64) string {
	result := fmt.Sprintf("接口地址: [%s]\n", context2.GetRequestUrlFromCtx(ctx))
	traceId, ok := ctx.Value(constants.TraceId).(string)
	if ok {
		result += fmt.Sprintf("日志id: [%s]\n", traceId)
	}

	result += fmt.Sprintf("执行时间: %d ms\n", tm)
	return result
}

type GetApiCommonErrTemplateOption struct {
	Url        string // 接口地址
	TraceId    string // 日志id
	ErrMessage string // 错误信息
}

// GetApiCommonErrTemplate 接口执行错误通用模板
func (ts *templateService) GetApiCommonErrTemplate(ctx context.Context, opt *GetApiCommonErrTemplateOption) string {
	var result string

	result += fmt.Sprintf("接口地址: [%s]\n", opt.Url)
	if len(opt.TraceId) > 0 {
		result += fmt.Sprintf("日志id: [%s]\n", opt.TraceId)
	}

	result += fmt.Sprintf("错误信息: [%s]\n", opt.ErrMessage)
	return result
}
