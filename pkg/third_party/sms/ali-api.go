package sms

import (
	"context"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	trlogger "github.com/woaijssss/tros/logx"
	"github.com/woaijssss/tros/trerror"
)

type aliSmsService struct{}

var AliSmsService = new(aliSmsService)

type AliSendSmsOption struct {
	RegionId        string
	AccessKeyId     string // oss访问key id
	AccessKeySecret string // oss访问key secret

	PhoneNumber   string
	SignName      string
	TemplateCode  string
	TemplateParam string
}

func (ass *aliSmsService) Send(ctx context.Context, opt *AliSendSmsOption) error {
	client, err := dysmsapi.NewClientWithAccessKey(opt.RegionId, opt.AccessKeyId, opt.AccessKeySecret)
	if err != nil {
		trlogger.Errorf(ctx, "aliSmsService send sms dysmsapi.NewClientWithAccessKey err: [%+v]", err)
		return err
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https" // 使用HTTPS协议
	request.PhoneNumbers = opt.PhoneNumber
	request.SignName = opt.SignName
	request.TemplateCode = opt.TemplateCode
	request.TemplateParam = opt.TemplateParam
	response, err := client.SendSms(request)
	if err != nil {
		trlogger.Errorf(ctx, "aliSmsService send sms client.SendSms err: [%+v]", err)
		return err
	}

	if response == nil {
		trlogger.Errorf(ctx, "aliSmsService send sms response == nil")
		return trerror.TR_ERROR
	}

	if response.Code != "OK" {
		trlogger.Errorf(ctx, "aliSmsService send sms response err: [%+v]", *response)
		return trerror.TR_ERROR
	}

	return nil
}
