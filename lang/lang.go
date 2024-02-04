// Package lang for globalization
package lang

// Lang type of language
type Lang string

const (
	// CN means Chinese
	CN Lang = "cn"
	// EN means English
	EN Lang = "en"
)

// Default Lang, config by cfgkey.GlobalLang
func Default() Lang {
	//str := conf.GetString(cfgkey.GlobalLang)
	//switch str {
	//case "cn":
	//	return CN
	//case "en":
	//	return EN
	//}
	return CN
}
