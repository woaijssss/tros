package wechat

import (
	"context"
	"fmt"
	"github.com/woaijssss/tros"
	"github.com/woaijssss/tros/conf"
	"github.com/woaijssss/tros/constants"
	trlogger "github.com/woaijssss/tros/logx"
	"github.com/woaijssss/tros/pkg/utils"
	"github.com/woaijssss/tros/pkg/utils/encrypt"
)

var Client = new(client)

const (
	UgcCheckSceneInfo      = 1 // 资料
	UgcCheckSceneComment   = 2 // 评论
	UgcCheckSceneForum     = 3 // 论坛
	UgcCheckSceneCommunity = 4 // 社交日志
)

func (c *client) Init(atx tros.AppContext) error {
	Client.appId = conf.GetString(constants.WechatAppid)
	Client.appSecret = conf.GetString(constants.WechatAppSecret)
	Client.mchId = conf.GetString(constants.WechatMchId)
	Client.wechatPayApiV2Key = conf.GetString(constants.WechatApiV2Key)
	Client.wechatPayApiV3Key = conf.GetString(constants.WechatApiV3Key)
	Client.redisWxAccessTokenKey = "wxAccessToken"
	Client.redisWxAccessTokenTimeout = 6000

	return nil
}

func (c *client) SetAppId(appId string) {
	c.appId = appId
}

func (c *client) SetAppSecret(appSecret string) {
	c.appSecret = appSecret
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

type MiniPayOrderOption struct {
	/* 以下为必填字段 */
	AppId    string `xml:"appid"`     // 小程序ID
	MchId    string `xml:"mch_id"`    // 微信支付分配的商户号
	NonceStr string `xml:"nonce_str"` // 随机字符串,长度要求在32位以内
	//Sign     string `xml:"sign"`      // 签名
	Body string `xml:"body"` // 商品描述，商品简单描述
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

type MiniPayOrderResponse struct {
	// 以下字段在return_code为SUCCESS的时候有效
	MchId    string `xml:"mch_id"`    // 调用接口提交的商户号
	PrepayId string `xml:"prepay_id"` // 预支付交易会话标识 微信生成的预支付会话标识，用于后续接口调用中使用，该值有效期为2小时
	/*
		二维码链接 trade_type=NATIVE时有返回，此url用于生成支付二维码，然后提供给用户进行扫码支付。
		注意：code_url的值并非固定，使用时按照URL格式转成二维码即可。时效性为2小时
	*/
	CodeUrl string `xml:"code_url"`

	// wx.requestPayment(OBJECT) required parameters
	WxObject WxRequestPaymentObject
	PaySign  string // 签名
}

type WxRequestPaymentObject struct {
	AppId     string `json:"appId"`     // 调用接口提交的小程序ID
	Timestamp string `json:"timeStamp"` // 时间戳从1970年1月1日00:00:00至今的秒数,即当前的时间
	NonceStr  string `json:"nonceStr"`  // 随机字符串，长度为32个字符以下。
	Package   string `json:"package"`   // 统一下单接口返回的 prepay_id 参数值，提交格式如：prepay_id=*
	SignType  string `json:"signType"`  // 签名类型，默认为MD5，支持HMAC-SHA256和MD5。注意此处需与统一下单的签名类型一致
}

// MiniPayOrder 微信小程序预支付订单接口
func (c *client) MiniPayOrder(ctx context.Context, opt *MiniPayOrderOption) (*MiniPayOrderResponse, error) {
	//signature := utils.StructToXMLKeyValueSorted(*opt)
	//stringSignTemp := fmt.Sprintf("%s&key=%s", signature, c.wechatPayApiV2Key)
	//sign := encrypt.EncodeMD5Upper(stringSignTemp)

	sign := c.Signature(*opt, []string{})

	response, err := c.miniPayOrderRequest(ctx, &miniPayOrderOption{
		AppId:          opt.AppId,
		MchId:          opt.MchId,
		NonceStr:       opt.NonceStr,
		Sign:           sign,
		Body:           opt.Body,
		OutTradeNo:     opt.OutTradeNo,
		TotalFee:       opt.TotalFee,
		SpbillCreateIp: opt.SpbillCreateIp,
		NotifyUrl:      opt.NotifyUrl,
		TradeType:      opt.TradeType,
		OpenId:         opt.OpenId,
	})

	if err != nil {
		trlogger.Errorf(ctx, "MiniPayOrder miniPayOrderRequest fail, err: [%+v]", err)
		return nil, err
	}

	wxObject := WxRequestPaymentObject{
		AppId:     response.AppId,
		Timestamp: utils.GetCurrentTimestampString(),
		NonceStr:  GenerateNonceStr(),
		Package:   fmt.Sprintf("prepay_id=%s", response.PrepayId),
		SignType:  "MD5", // 固定MD5
	}
	wxObjectSignature := utils.StructToJSONKeyValueSorted(wxObject)
	wxObjectSignTemp := fmt.Sprintf("%s&key=%s", wxObjectSignature, c.wechatPayApiV2Key)
	wxObjectSign := encrypt.EncodeMD5Upper(wxObjectSignTemp)

	return &MiniPayOrderResponse{
		MchId:    response.MchId,
		PrepayId: response.PrepayId,
		CodeUrl:  response.CodeUrl,
		WxObject: wxObject,
		PaySign:  wxObjectSign,
	}, nil
}
