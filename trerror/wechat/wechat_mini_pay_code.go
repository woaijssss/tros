package wechat

import "github.com/woaijssss/tros/trerror"

func GetErrCodeDes(errCode string) (*WechatMiniPayCodeValue, error) {
	if errCode == "" {
		return nil, nil
	}

	v, ok := wechatMiniPayCode2Msg[errCode]
	if !ok {
		return nil, trerror.TR_INVALID_ERROR
	}
	return v, nil
}

type WechatMiniPayCodeValue struct {
	Desc     string // 错误描述
	Reason   string // 错误原因
	Solution string // 解决方案
}

var wechatMiniPayCode2Msg = map[string]*WechatMiniPayCodeValue{
	"INVALID_REQUEST": {
		Desc:     "参数错误",
		Reason:   "参数格式有误或者未按规则上传",
		Solution: "订单重入时，要求参数值与原请求一致，请确认参数问题",
	},
	"NOAUTH": {
		Desc:     "商户无此接口权限",
		Reason:   "商户未开通此接口权限",
		Solution: "请商户前往申请此接口权限",
	},
	"ORDERPAID": {
		Desc:     "商户订单已支付",
		Reason:   "商户订单已支付，无需重复操作",
		Solution: "商户订单已支付，无需更多操作",
	},
	"ORDERCLOSED": {
		Desc:     "订单已关闭",
		Reason:   "当前订单已关闭，无法支付",
		Solution: "当前订单已关闭，请重新下单",
	},
	"SYSTEMERROR": {
		Desc:     "系统错误",
		Reason:   "系统超时",
		Solution: "系统异常，请用相同参数重新调用",
	},
	"APPID_NOT_EXIST": {
		Desc:     "APPID不存在",
		Reason:   "参数中缺少APPID",
		Solution: "请检查APPID是否正确",
	},
	"MCHID_NOT_EXIST": {
		Desc:     "MCHID不存在",
		Reason:   "参数中缺少MCHID",
		Solution: "请检查MCHID是否正确",
	},
	"APPID_MCHID_NOT_MATCH": {
		Desc:     "appid和mch_id不匹配",
		Reason:   "appid和mch_id不匹配",
		Solution: "请确认appid和mch_id是否匹配",
	},
	"LACK_PARAMS": {
		Desc:     "缺少参数",
		Reason:   "缺少必要的请求参数",
		Solution: "请检查参数是否齐全",
	},
	"OUT_TRADE_NO_USED": {
		Desc:     "商户订单号重复",
		Reason:   "同一笔交易不能多次提交",
		Solution: "请核实商户订单号是否重复提交",
	},
	"SIGNERROR": {
		Desc:     "签名错误",
		Reason:   "参数签名结果不正确",
		Solution: "请检查签名参数和方法是否都符合签名算法要求",
	},
	"XML_FORMAT_ERROR": {
		Desc:     "XML格式错误",
		Reason:   "XML格式错误",
		Solution: "请检查XML参数格式是否正确",
	},
	"REQUIRE_POST_METHOD": {
		Desc:     "请使用post方法",
		Reason:   "未使用post传递参数",
		Solution: "请检查请求参数是否通过post方法提交",
	},
	"POST_DATA_EMPTY": {
		Desc:     "post数据为空",
		Reason:   "post数据不能为空",
		Solution: "请检查post数据是否为空",
	},
	"NOT_UTF8": {
		Desc:     "编码格式错误",
		Reason:   "未使用指定编码格式",
		Solution: "请使用UTF-8编码格式",
	},
}
