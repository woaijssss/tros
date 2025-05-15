package feishu

import (
	"bytes"
	"context"
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
