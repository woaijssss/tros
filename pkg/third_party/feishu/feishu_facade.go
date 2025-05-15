package feishu

import (
	"context"
	"github.com/woaijssss/tros"
	"github.com/woaijssss/tros/conf"
	"github.com/woaijssss/tros/constants"
)

var Client = new(client)

func (c *client) Init(atx tros.AppContext) error {
	c.WebHookUrl = conf.GetString(constants.FeiShuWebHookUrlKey)
	c.SignKey = conf.GetString(constants.FeiShuSignKey)

	return nil
}

func (c *client) BusinessFeiShuRobotTextMessage(ctx context.Context, content string) error {
	if len(c.WebHookUrl) <= 0 { // 如果未设置url，则不推送消息
		return nil
	}

	return c.businessFeiShuRobotTextMessage(ctx, content)
}
