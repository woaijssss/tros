package currency

import (
	"errors"
	"fmt"
)

type Currency struct {
	Code   string
	CnName string
}

var (
	AUD = &Currency{Code: "AUD", CnName: "澳币"}
	CAD = &Currency{Code: "CAD", CnName: "加拿大元／加币"}
	CHF = &Currency{Code: "CHF", CnName: "瑞士法郎"}
	CNY = &Currency{Code: "CNY", CnName: "人民币"}
	CZK = &Currency{Code: "CZK", CnName: "捷克克朗"}
	DKK = &Currency{Code: "DKK", CnName: "丹麦克朗"}
	EUR = &Currency{Code: "EUR", CnName: "欧元"}
	GBP = &Currency{Code: "GBP", CnName: "英镑"}
	HKD = &Currency{Code: "HKD", CnName: "港币"}
	HUF = &Currency{Code: "HUF", CnName: "匈牙利福林"}
	ILS = &Currency{Code: "ILS", CnName: "以色列谢克尔"}
	JPY = &Currency{Code: "JPY", CnName: "日元"}
	KRW = &Currency{Code: "KRW", CnName: "韩元"}
	MXN = &Currency{Code: "MXN", CnName: "墨西哥比索"}
	NOK = &Currency{Code: "NOK", CnName: "挪威克朗"}
	NZD = &Currency{Code: "NZD", CnName: "新西兰元"}
	PLN = &Currency{Code: "PLN", CnName: "波兰兹罗提"}
	RON = &Currency{Code: "RON", CnName: "罗马尼亚列伊"}
	SEK = &Currency{Code: "SEK", CnName: "瑞典克朗"}
	SGD = &Currency{Code: "SGD", CnName: "新加坡元/新币"}
	THB = &Currency{Code: "THB", CnName: "泰铢"}
	USD = &Currency{Code: "USD", CnName: "美元"}
	ZAR = &Currency{Code: "ZAR", CnName: "兰特"}
)

var allCurrencyMap = map[string]*Currency{
	"AUD": AUD,
	"CAD": CAD,
	"CHF": CHF,
	"CNY": CNY,
	"CZK": CZK,
	"DKK": DKK,
	"EUR": EUR,
	"GBP": GBP,
	"HKD": HKD,
	"HUF": HUF,
	"ILS": ILS,
	"JPY": JPY,
	"KRW": KRW,
	"MXN": MXN,
	"NOK": NOK,
	"NZD": NZD,
	"PLN": PLN,
	"RON": RON,
	"SEK": SEK,
	"SGD": SGD,
	"THB": THB,
	"USD": USD,
	"ZAR": ZAR,
}

func GetByCode(code string) (*Currency, error) {
	cy := allCurrencyMap[code]
	if cy == nil {
		return nil, errors.New(fmt.Sprintf("illegal argument currency code: %s", code))
	}
	return cy, nil
}
