package wechat

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/woaijssss/tros/client/http"
	trlogger "github.com/woaijssss/tros/logx"
	"github.com/woaijssss/tros/pkg/utils"
	"github.com/woaijssss/tros/trerror/wechat"
)

// WeChat public client
type client struct {
	appId     string
	appSecret string

	// WeChat's payment field
	mchId             string
	wechatPayApiV2Key string
	wechatPayApiV3Key string

	redisWxAccessTokenKey     string
	redisWxAccessTokenTimeout int64 // expire time
}

// WeChat pay client
type payClient struct {
	client
}

const (
	getAccessTokenUrl     = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	getGenerateUrlLinkUrl = "https://api.weixin.qq.com/wxa/generate_urllink?access_token=%s"
	messageCheckUrl       = "https://api.weixin.qq.com/wxa/msg_sec_check?access_token=%s"
	imageCheckUrl         = "https://api.weixin.qq.com/wxa/media_check_async?access_token=%s"
	payOrderGatewayUrl    = "https://api.mch.weixin.qq.com/pay/unifiedorder"
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

func (c *client) getAccessTokenWithInput(ctx context.Context, appid, appSecret string) (*GetWechatAccessTokenResponse, error) {
	url := fmt.Sprintf(getAccessTokenUrl, appid, appSecret)
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

type miniPayOrderOption struct {
	/* 以下为必填字段 */
	AppId    string `xml:"appid"`     // 小程序ID
	MchId    string `xml:"mch_id"`    // 微信支付分配的商户号
	NonceStr string `xml:"nonce_str"` // 随机字符串,长度要求在32位以内
	Sign     string `xml:"sign"`      // 签名
	Body     string `xml:"body"`      // 商品描述，商品简单描述
	/* OutTradeNo
	商户订单号
	商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*且在同一个商户号下唯一
	*/
	OutTradeNo     string `xml:"out_trade_no"`
	TotalFee       int32  `xml:"total_fee"`        // 标价金额，订单总金额，单位为分
	SpbillCreateIp string `xml:"spbill_create_ip"` // 终端IP，即：发起人IP地址 支持IPV4和IPV6两种格式的IP地址。调用微信支付API的机器IP
	NotifyUrl      string `xml:"notify_url"`       // 通知地址 异步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。公网域名必须为https
	/* TradeType
	交易类型 取值范围如下：
		JSAPI--JSAPI支付（或小程序支付）、NATIVE--Native支付、APP--app支付，MWEB--H5支付，不同trade_type决定了调起支付的方式，请根据支付产品正确上传
		MICROPAY--付款码支付，付款码支付有单独的支付接口，所以接口不需要上传，该字段在对账单中会出现
	*/
	TradeType string `xml:"trade_type"`

	/* 以下为选填字段 */
	//SignType string `xml:"sign_type"` // 签名类型 默认为MD5，支持HMAC-SHA256和MD5
	OpenId string `xml:"openid"` // 用户标识（trade_type=JSAPI，此参数必传，用户在商户appid下的唯一标识）
}

type miniPayOrderResponse struct {
	ReturnCode string `xml:"return_code"` // 返回状态码 SUCCESS/FAIL 此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
	ReturnMsg  string `xml:"return_msg"`  // 返回信息 返回信息，如非空，为错误原因 签名失败、参数格式校验错误等

	// 以下字段在return_code为SUCCESS的时候有效
	AppId      string `xml:"appid"`        // 调用接口提交的小程序ID
	MchId      string `xml:"mch_id"`       // 调用接口提交的商户号
	DeviceInfo string `xml:"device_info"`  // 自定义参数，可以为请求支付的终端设备号等
	NonceStr   string `xml:"nonce_str"`    // 微信返回的随机字符串
	Sign       string `xml:"sign"`         // 微信返回的签名值
	ResultCode string `xml:"result_code"`  // 业务结果 SUCCESS/FAIL
	ErrCode    string `xml:"err_code"`     // 错误代码 详细参见微信小程序支付的接口错误列表
	ErrCodeDes string `xml:"err_code_des"` // 错误信息描述

	// 以下字段在return_code 和result_code都为SUCCESS的时候有返回
	TradeType string `xml:"trade_type"` // 交易类型，取值为：JSAPI，NATIVE，APP等
	PrepayId  string `xml:"prepay_id"`  // 预支付交易会话标识 微信生成的预支付会话标识，用于后续接口调用中使用，该值有效期为2小时
	/*
		二维码链接 trade_type=NATIVE时有返回，此url用于生成支付二维码，然后提供给用户进行扫码支付。
		注意：code_url的值并非固定，使用时按照URL格式转成二维码即可。时效性为2小时
	*/
	CodeUrl string `xml:"code_url"`
}

// MiniPayOrder 微信小程序预支付订单接口
func (c *client) miniPayOrderRequest(ctx context.Context, opt *miniPayOrderOption) (*miniPayOrderResponse, error) {
	resp, err := http.NewHttpClient().PostXml(ctx, payOrderGatewayUrl, opt)
	if err != nil {
		trlogger.Errorf(ctx, "miniPayOrder PostXml fail, err: [%+v]", err)
		return nil, err
	}

	var result miniPayOrderResponse
	err = http.ResXmlToObj(resp, &result)
	if err != nil {
		trlogger.Errorf(ctx, "miniPayOrder unmarshal xml err: %v", err)
		return nil, err
	}

	if result.ResultCode != "SUCCESS" {
		err = errors.New(result.ReturnMsg)
		trlogger.Errorf(ctx, "miniPayOrder response fail, err: [%+v]", err)
		return nil, err
	}

	if result.ResultCode != "SUCCESS" {
		msg, err := wechat.GetErrCodeDes(result.ErrCode)
		if err != nil {
			trlogger.Errorf(ctx, "miniPayOrder GetErrCodeDes fail, err: [%+v]", err)
			return nil, err
		}

		err = errors.New(msg.Desc)
		trlogger.Errorf(ctx, "miniPayOrder ResultCode fail, err: [%+v]", err)
		return nil, err
	}

	// True request successful, return result parameters
	return &result, nil
}
