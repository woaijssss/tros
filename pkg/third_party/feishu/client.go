package feishu

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/woaijssss/tros/client/http"
	trlogger "github.com/woaijssss/tros/logx"
	"github.com/woaijssss/tros/pkg/utils"
	"github.com/woaijssss/tros/pkg/utils/encrypt"
	"time"
)

type client struct {
	WebHookUrl string // 飞书机器人通知地址
	SignKey    string // 签名密钥

	// 飞书接口鉴权
	AppId     string
	AppSecret string
}

type RobotMessageContent struct {
	Text string `json:"text"`
}

type RobotTextMessage struct {
	MsgType string               `json:"msg_type"`
	Content *RobotMessageContent `json:"content"`

	Timestamp string `json:"timestamp"` // 时间戳
	Sign      string `json:"sign"`      // 得到的签名字符串
}

const (
	tenantTokenUrl = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal/"
)

func (c *client) businessFeiShuRobotTextMessage(ctx context.Context, content string) error {
	contentMsg := RobotMessageContent{
		Text: content,
	}
	requestParam := RobotTextMessage{
		MsgType: "text",
		Content: &contentMsg,
	}

	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	data := timestamp + "\n" + c.SignKey
	signature, err := encrypt.Sha256Encode(data)
	if err != nil {
		trlogger.Errorf(ctx, "businessFeiShuRobotTextMessage encrypt.Sha256Encode err: [%+v]", err)
		return err
	}

	requestParam.Timestamp = timestamp
	requestParam.Sign = signature

	b, err := utils.ToJsonByte(requestParam)
	resp, err := http.NewHttpClient().Post(ctx, c.WebHookUrl, bytes.NewReader(b))
	if err != nil {
		trlogger.Errorf(ctx, "businessFeiShuRobotTextMessage http post err: [%+v]", err)
		return err
	}
	resp.Body.Close()
	return nil
}

type getTenantTokenOption struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type GetTenantTokenResponse struct {
	Code              int32  `json:"code"`
	Msg               string `json:"msg"`
	AppAccessToken    string `json:"app_access_token"`
	Expire            int32  `json:"expire"`
	TenantAccessToken string `json:"tenant_access_token"`
}

func (c *client) getTenantToken(ctx context.Context) (string, error) {
	b, err := utils.ToJsonByte(&getTenantTokenOption{
		AppId:     c.AppId,
		AppSecret: c.AppSecret,
	})
	if err != nil {
		trlogger.Errorf(ctx, "getTenantToken marshal err: [%+v]", err)
		return "", err
	}
	resp, err := http.NewHttpClient().Post(ctx, tenantTokenUrl, bytes.NewReader(b))
	if err != nil {
		trlogger.Errorf(ctx, "getTenantToken send err: %v", err)
		return "", err
	}

	ret := new(GetTenantTokenResponse)
	err = http.ResToObj(resp, ret)
	if err != nil {
		trlogger.Errorf(ctx, "getTenantToken GetTenantTokenResponse utils.ResToObj err: %v", err)
		return "", err
	}

	if ret.Code != 0 {
		trlogger.Errorf(ctx, "getTenantToken http response [%d][%s]", ret.Code, ret.Msg)
		return "", errors.New(ret.Msg)
	}

	return ret.TenantAccessToken, nil
}
