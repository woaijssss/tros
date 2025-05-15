package feishu

import (
	"context"
	"github.com/woaijssss/tros"
	"github.com/woaijssss/tros/conf"
	"github.com/woaijssss/tros/constants"
)

var Client = new(client)

func (c *client) Init(atx tros.AppContext) error {
	// 默认使用测试环境的通知地址
	c.WebHookUrl = conf.GetString(constants.FeiShuWebHookUrlKeyForTest)
	c.SignKey = conf.GetString(constants.FeiShuSignKeyForTest)

	return nil
}

// SetEnv 设置通知地址（需要主动调用！！！）
func (c *client) SetEnv(systemEnv constants.SystemEnv) {
	if systemEnv == constants.Prod {
		c.WebHookUrl = conf.GetString(constants.FeiShuWebHookUrlKeyForProd)
		c.SignKey = conf.GetString(constants.FeiShuSignKeyForProd)
	} else if systemEnv == constants.Test {
		c.WebHookUrl = conf.GetString(constants.FeiShuWebHookUrlKeyForTest)
		c.SignKey = conf.GetString(constants.FeiShuSignKeyForTest)
	} else { // 如果环境错误，则不通知！！！
		c.WebHookUrl = ""
		c.SignKey = ""
	}
}

func (c *client) BusinessFeiShuRobotTextMessage(ctx context.Context, content string) error {
	if len(c.WebHookUrl) <= 0 { // 如果未设置url，则不推送消息
		return nil
	}

	return c.businessFeiShuRobotTextMessage(ctx, content)
}
