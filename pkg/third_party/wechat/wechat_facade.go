package wechat

import (
	"context"
	"gitee.com/idigpower/tros"
	"gitee.com/idigpower/tros/conf"
	"gitee.com/idigpower/tros/constants"
	trlogger "gitee.com/idigpower/tros/logx"
)

var Client = new(client)

const (
	UgcCheckSceneInfo      = 1
	UgcCheckSceneComment   = 2
	UgcCheckSceneForum     = 3
	UgcCheckSceneCommunity = 4
)

func (c *client) Init(atx tros.AppContext) error {
	Client.appId = conf.GetString(constants.WechatAppid)
	Client.appSecret = conf.GetString(constants.WechatAppSecret)
	Client.redisWxAccessTokenKey = "wxAccessToken"
	Client.redisWxAccessTokenTimeout = 6000

	return nil
}

type GetWechatAccessTokenResponse struct {
	AccessToken string
	ExpiresIn   int64
}

// GetWechatAccessTokenWithCache 带缓存的accesstoken
func (c *client) GetWechatAccessTokenWithCache(ctx context.Context) (string, error) {
	accessTokenObj, err := c.getAccessToken(ctx)
	if err != nil {
		trlogger.Errorf(ctx, "wechatTextCheck GetWechatAccessToken err: [%+v]", err)
		return "", err
	}
	return accessTokenObj.AccessToken, nil
}

type GetGenerateUrlLinkOption struct {
	Path       string `json:"path"`        // 通过 URL Link 进入的小程序页面路径，必须是已经发布的小程序存在的页面，不可携带 query 。path 为空时会跳转小程序主页
	Query      string `json:"query"`       // 通过 URL Link 进入小程序时的query，最大1024个字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~%
	ExpireType int32  `json:"expire_type"` // 默认值0.小程序 URL Link 失效类型，失效时间：0，失效间隔天数：1
	//expireTime     int32  `json:"expire_time"`     // 到期失效的 URL Link 的失效时间，为 Unix 时间戳。生成的到期失效 URL Link 在该时间前有效。最长有效期为30天。expire_type 为 0 必填
	ExpireInterval int32 `json:"expire_interval"` // 到期失效的URL Link的失效间隔天数。生成的到期失效URL Link在该间隔时间到达前有效。最长间隔天数为30天。expire_type 为 1 必填
}

func (c *client) GetGenerateUrlLink(ctx context.Context, opt *GetGenerateUrlLinkOption) (string, error) {
	accessTokenObj, err := c.getAccessToken(ctx)
	if err != nil {
		trlogger.Errorf(ctx, "GetGenerateUrlLink GetWechatAccessToken err: [%+v]", err)
		return "", err
	}
	return c.getGenerateUrlLink(ctx, accessTokenObj.AccessToken, opt)
}

type MessageCheckOption struct {
	//AccessToken string `json:"access_token"`
	Content string `json:"content"` // 实际需要检测的内容
	Scene   int32  `json:"scene"`   // 场景枚举值（1 资料；2 评论；3 论坛；4 社交日志）
	OpenId  string `json:"openid"`  // 用户的openid（用户需在近两小时访问过小程序）
	Version int32  `json:"version"`
}

// MessageCheck 微信平台的文字内容违规检测
func (c *client) MessageCheck(ctx context.Context, accessToken string, opt *MessageCheckOption) (bool, error) {
	if opt == nil {
		return true, nil
	}
	opt.Version = 2
	//return c.messageCheck(ctx, accessToken, opt)
	//for i := int32(1); i < 5; i++ {
	//opt.Scene = i
	ok, err := c.messageCheck(ctx, accessToken, opt)

	if err != nil {
		trlogger.Errorf(ctx, "MessageCheck err: [%+v]", err)
		return ok, err
	}
	//}
	return ok, nil
}

type ImageCheckOption struct {
	MediaUrl string `json:"media_url"` // 实际需要检测的图片地址
	Scene    int32  `json:"scene"`     // 场景枚举值（1 资料；2 评论；3 论坛；4 社交日志）
	OpenId   string `json:"openid"`    // 用户的openid（用户需在近两小时访问过小程序）

	MediaType int32 `json:"media_type"` // 1:音频;2:图片（这里选择2）
	Version   int32 `json:"version"`    // 默认为 2
}

// ImageCheck 微信平台的图片内容违规检测
func (c *client) ImageCheck(ctx context.Context, accessToken string, opt *ImageCheckOption) (string, error) {
	if opt == nil {
		return "", nil
	}
	opt.Version = 2
	opt.MediaType = 2
	traceId, err := c.imageCheck(ctx, accessToken, opt)
	if err != nil {
		return "", err
	}
	return traceId, nil
}
