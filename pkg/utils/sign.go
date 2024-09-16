package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/woaijssss/tros/pkg/utils/encrypt"
	"net/http"
	"sort"
	"strings"
)

const SIGNKEY = "5506eb75f447d18a2ff9cfdd7d9820fe"

type SignPrams struct {
	m    map[string]interface{} // 参数
	sign string                 // 签名
}

// 签名算法
func (p *SignPrams) InitSign() string {
	var keys []string
	for k := range p.m {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	builder := strings.Builder{}
	for _, v := range keys {
		if p.m[v] == nil {
			continue
		}
		if p.m[v].(string) != "" {
			builder.WriteString(v)
			builder.WriteString("=")
			builder.WriteString(fmt.Sprint(p.m[v]))
			builder.WriteString("&")
		} else if p.m[v].(string) != "" {
		}
	}
	builder.WriteString("key=" + SIGNKEY)
	//p.sign = strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(builder.String()))))
	p.sign = strings.ToUpper(fmt.Sprintf("%x", encrypt.EncodeMD5(builder.String())))

	return p.sign
}

func (p *SignPrams) GetSign() string {
	p.InitSign()
	return p.sign
}

func CheckSign() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		buf := make([]byte, 1024)
		n, _ := c.Request.Body.Read(buf)
		paramMap := make(map[string]interface{})
		_ = json.Unmarshal(buf[0:n], &paramMap)
		SignPrams := &SignPrams{}
		SignPrams.m = paramMap
		sign := SignPrams.GetSign()

		if strings.EqualFold(sign, paramMap["sign"].(string)) {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				//"code": errcode.ERROR_AUTH_CHECK_SIGN_FAIL,
				//"msg":  errcode.GetMsg(errcode.ERROR_AUTH_CHECK_SIGN_FAIL),
				"data": data,
			})
			c.Abort()
			return
		}
	}
}

func CreateSign(param []byte) string {
	var paramMap map[string]interface{}
	_ = json.Unmarshal(param, &paramMap)
	SignPrams := &SignPrams{}
	SignPrams.m = paramMap
	sign := SignPrams.GetSign()
	return sign
}
