package wechat

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"gitee.com/idigpower/tros/client/http"
	trlogger "gitee.com/idigpower/tros/logx"
	"gitee.com/idigpower/tros/pkg/utils"
)

type client struct {
	appId                     string
	appSecret                 string
	redisWxAccessTokenKey     string
	redisWxAccessTokenTimeout int64 // 过期时间
}

const (
	getAccessTokenUrl     = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	getGenerateUrlLinkUrl = "https://api.weixin.qq.com/wxa/generate_urllink?access_token=%s"
	messageCheckUrl       = "https://api.weixin.qq.com/wxa/msg_sec_check?access_token=%s"
	imageCheckUrl         = "https://api.weixin.qq.com/wxa/media_check_async?access_token=%s"
)

type getAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func (c *client) getAccessToken(ctx context.Context) (*GetWechatAccessTokenResponse, error) {
	url := fmt.Sprintf(getAccessTokenUrl, c.appId, c.appSecret)
	resp, _ := http.NewHttpClient().Get(ctx, url)
	ret := new(getAccessTokenResponse)
	err := http.ResToObj(resp, ret)
	if err != nil {
		trlogger.Errorf(ctx, "getAccessToken unmarshal err: %v", err)
		return nil, err
	}

	return &GetWechatAccessTokenResponse{
		AccessToken: ret.AccessToken,
		ExpiresIn:   ret.ExpiresIn,
	}, nil
}

type getGenerateUrlLinkResponse struct {
	Errcode int32  `json:"errcode"`  // 错误码
	Errmsg  string `json:"errmsg"`   // 错误信息
	UrlLink string `json:"url_link"` // 生成的小程序 URL Link
}

func (c *client) getGenerateUrlLink(ctx context.Context, accessToken string, opt *GetGenerateUrlLinkOption) (string, error) {
	url := fmt.Sprintf(getGenerateUrlLinkUrl, accessToken)
	b, err := utils.ToJsonByte(opt)
	if err != nil {
		trlogger.Errorf(ctx, "getGenerateUrlLink marshal err: [%+v]", err)
		return "", err
	}

	resp, err := http.NewHttpClient().Post(ctx, url, bytes.NewReader(b))
	if err != nil {
		trlogger.Errorf(ctx, "getGenerateUrlLink send err: %v", err)
		return "", err
	}
	ret := new(getGenerateUrlLinkResponse)
	err = http.ResToObj(resp, ret)
	if err != nil {
		trlogger.Errorf(ctx, "getGenerateUrlLink unmarshal err: %v", err)
		return "", err
	}

	if ret.Errcode != 0 {
		err = errors.New(ret.Errmsg)
		trlogger.Errorf(ctx, "getGenerateUrlLink response err: %v", err)
		return "", err
	}

	return ret.UrlLink, nil
}

type messageCheckResponse struct {
	Errcode int32                 `json:"errcode"` // 错误码
	Errmsg  string                `json:"errmsg"`  // 错误信息
	Detail  []*messageCheckDetail `json:"detail"`  // 详细检测结果
	Result  *messageCheckResult   `json:"result"`  // 综合检测结果
}
type messageCheckResult struct {
	Suggest string `json:"suggest"` // 建议，有risky、pass、review三种值
	Label   int32  `json:"label"`   // 命中标签枚举值，100 正常；10001 广告；20001 时政；20002 色情；20003 辱骂；20006 违法犯罪；20008 欺诈；20012 低俗；20013 版权；21000 其他
	Keyword string `json:"keyword"` // 命中的自定义关键词 todo 补充后续显示逻辑
}

type messageCheckDetail struct {
	Strategy string `json:"strategy"` // 策略类型
	Errcode  int32  `json:"errcode"`  // 错误码，仅当该值为0时，该项结果有效
	Suggest  string `json:"suggest"`  // 建议，有risky、pass、review三种值
	Label    int32  `json:"label"`    // 命中标签枚举值，100 正常；10001 广告；20001 时政；20002 色情；20003 辱骂；20006 违法犯罪；20008 欺诈；20012 低俗；20013 版权；21000 其他
	Keyword  string `json:"keyword"`  // 命中的自定义关键词
	Prob     int32  `json:"prob"`     // 0-100，代表置信度，越高代表越有可能属于当前返回的标签（label）
}

func (c *client) messageCheck(ctx context.Context, accessToken string, opt *MessageCheckOption) (bool, error) {
	url := fmt.Sprintf(messageCheckUrl, accessToken)
	b, err := utils.ToJsonByte(opt)
	if err != nil {
		trlogger.Errorf(ctx, "messageCheck marshal err: [%+v]", err)
		return true, err
	}

	resp, err := http.NewHttpClient().Post(ctx, url, bytes.NewReader(b))
	if err != nil {
		trlogger.Errorf(ctx, "messageCheck send err: %v", err)
		return true, err
	}
	trlogger.Infof(ctx, "send messageCheck success")
	//if resp.StatusCode != http.StatusOK {
	//	trlogger.Errorf(ctx, "send llm  http.Post statusCode: %d", resp.StatusCode)
	//	return "", errors.New("send llm failed")
	//}

	ret := new(messageCheckResponse)
	err = http.ResToObj(resp, ret)
	if err != nil {
		trlogger.Errorf(ctx, "send llm d utils.ResToObj err: %v", err)
		return true, err
	}

	if ret.Errcode != 0 {
		trlogger.Errorf(ctx, "messageCheck http response err: %+v", ret.Errmsg)
		return true, errors.New(ret.Errmsg)
	}

	return msgResult(ctx, ret.Result)
}

type imageCheckResponse struct {
	Errcode int32  `json:"errcode"`  // 错误码
	Errmsg  string `json:"errmsg"`   // 错误信息
	TraceId string `json:"trace_id"` // 唯一请求标识，标记单次请求，用于匹配异步推送结果
}

func (c *client) imageCheck(ctx context.Context, accessToken string, opt *ImageCheckOption) (string, error) {
	url := fmt.Sprintf(imageCheckUrl, accessToken)
	fmt.Println("url: ", url)
	b, err := utils.ToJsonByte(opt)
	if err != nil {
		trlogger.Errorf(ctx, "imageCheck marshal err: [%+v]", err)
		return "", err
	}

	resp, err := http.NewHttpClient().Post(ctx, url, bytes.NewReader(b))
	if err != nil {
		trlogger.Errorf(ctx, "imageCheck send err: %v", err)
		return "", err
	}
	trlogger.Infof(ctx, "send imageCheck success")

	ret := new(imageCheckResponse)
	err = http.ResToObj(resp, ret)
	if err != nil {
		trlogger.Errorf(ctx, "send llm d utils.ResToObj err: %v", err)
		return "", err
	}

	//return msgResult(ctx, ret.Result)
	// todo 增加图片校验方法
	return ret.TraceId, nil
}

func msgDetail(ctx context.Context, detail []*messageCheckDetail) error {
	for _, d := range detail {
		fmt.Println("d: ", *d)
		if d.Label != 100 {
			err := fmt.Sprintf("命中违规词标签: [%d]", d.Label)
			trlogger.Warnf(ctx, "msgDetail err: [%+v]", err)
			return errors.New(err)
		}
	}
	return nil
}

func msgResult(ctx context.Context, result *messageCheckResult) (bool, error) {
	if result.Label != 100 {
		err := fmt.Sprintf("命中综合检测记过的违规词标签: [%d]", result.Label)
		trlogger.Warnf(ctx, "msgResult err: [%+v]", err)
		return false, nil // 成功检测不返回error，外部通过ok值判断
	}
	return true, nil
}
