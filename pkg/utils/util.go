package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
	"time"
)

func Decode(input []byte, output interface{}) error {
	buf := bytes.NewBuffer(input)
	decoder := json.NewDecoder(buf)
	decoder.UseNumber()
	err := decoder.Decode(output)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func GetHost(ctx *gin.Context) (host string) {
	hostUrl := ctx.Request.Host
	hostUrlList := strings.Split(hostUrl, ":")
	if len(hostUrlList) == 0 {
		host = ""
	} else {
		host = hostUrlList[0]
	}
	return host
}

func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func VerifyCarLicenseFormat(license string) bool {
	regular := "^[京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领A-Z]{1}[A-Z]{1}[A-Z0-9]{3,6}[A-Z0-9挂学警港澳]{1}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(license)
}

func VerifyIPFormat(ip string) bool {
	regular := "^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(ip)
}

func GetDateAndSec() string {
	strDate := time.Unix(time.Now().Unix(), 0).Format("20060102150405")
	return strDate
}
