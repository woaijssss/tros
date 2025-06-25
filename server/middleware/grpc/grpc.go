package grpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/woaijssss/tros/conf"
	"github.com/woaijssss/tros/constants"
	context2 "github.com/woaijssss/tros/context"
	trlogger "github.com/woaijssss/tros/logx"
	"github.com/woaijssss/tros/pkg/third_party/feishu"
	"github.com/woaijssss/tros/pkg/utils"
	"github.com/woaijssss/tros/trerror"
	"google.golang.org/grpc"
	"time"
)

// UnaryServerInterceptor returns unary gRpc tracing middleware
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx = context2.InsertTraceID(ctx) // set log id
		//ctx = context2.InsertRemoteIp(ctx) // set remote ip
		ctx = context2.InsertAllInfo(ctx)

		t1 := time.Now()
		resp, err = handler(ctx, req)
		t2 := time.Now()

		go asyncSendFeiShuAlarm(ctx, t1, t2, err)

		return
	}
}

func asyncSendFeiShuAlarm(ctx context.Context, t1, t2 time.Time, err error) {
	uri := context2.GetRequestUrlFromCtx(ctx)
	fmt.Println(uri)

	ignoreAlarmUris := conf.GetStringSlice(constants.IgnoreAlarmUris)

	if !utils.In(uri, ignoreAlarmUris) { // 只有不在忽略列表里的 接口报错，才需要推送
		// 异步推送消息
		go apiExecuteTimeAlarmNotify(ctx, t1, t2)
		go apiCommonAlarmNotify(ctx, err)
	}
}

func apiExecuteTimeAlarmNotify(ctx context.Context, t1, t2 time.Time) {
	elapsed := utils.CalcMillisecondsBetween(t1, t2)
	trlogger.Infof(ctx, "[%s] apiExecuteTimeAlarm api execute time: [%+v] ms\n", context2.GetRequestUrlFromCtx(ctx), elapsed)
	if elapsed > constants.MaxApiExecuteTimeMs {
		executeTimeContent := feishu.TemplateService.GetApiExecuteTimeTemplate(ctx, elapsed)
		feishu.Client.BusinessFeiShuRobotTextMessage(ctx, executeTimeContent)
	}
}

func apiCommonAlarmNotify(ctx context.Context, err error) {
	if err == nil { // 无错误不推送
		return
	}

	if errors.Is(err, trerror.TR_NO_PERMISSION) { // 特殊错误不推送
		return
	}

	opt := feishu.GetApiCommonErrTemplateOption{
		Url:        context2.GetRequestUrlFromCtx(ctx),
		ErrMessage: err.Error(),
	}

	traceId, ok := ctx.Value(constants.TraceId).(string)
	if ok {
		opt.TraceId = traceId
	}

	commonErrContent := feishu.TemplateService.GetApiCommonErrTemplate(ctx, &opt)
	feishu.Client.BusinessFeiShuRobotTextMessage(ctx, commonErrContent)
}
