package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"math/rand"
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

// GenerateRandomBool Generate boolean type values with a specified probability param prob,such as 0.5
func GenerateRandomBool(prob float64) bool {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randomNumber := r.Float64()
	return randomNumber >= prob
}

// GenerateCommonUniqueIdOnlyNumber General method to generate a unique ID (numeric only) based on length.
func GenerateCommonUniqueIdOnlyNumber(length int) string {
	rand.Seed(time.Now().UnixNano())
	// 用于存储生成的用户编号
	var uniqueID strings.Builder
	// 循环生成12位字符
	for i := 0; i < length; i++ {
		// 随机选择字符集合中的一个字符
		randomIndex := rand.Intn(len(uuidCommonChar))
		uniqueID.WriteRune(rune(uuidCommonChar[randomIndex]))
	}
	return uniqueID.String()
}

// GenerateCommonUniqueIdOnlyNumberReturnInteger General method to generate a unique ID (numeric only) based on length.
func GenerateCommonUniqueIdOnlyNumberReturnInteger(length int) int64 {
	rand.Seed(time.Now().UnixNano())
	// 用于存储生成的用户编号
	var uniqueID strings.Builder
	// 循环生成12位字符
	for i := 0; i < length; i++ {
		// 随机选择字符集合中的一个字符
		randomIndex := rand.Intn(len(uuidCommonChar))
		uniqueID.WriteRune(rune(uuidCommonChar[randomIndex]))
	}

	return String2Int64(uniqueID.String())
}

// GenerateUniqueIdByUniqueStr Generate an uuid of type int64 based on the specified unique string.
func GenerateUniqueIdByUniqueStr(s string) (int64, error) {
	// Calculate the MD5 hash value of a string
	hash := md5.Sum([]byte(s))
	hashBytes := hash[:]

	// Convert byte slices to hexadecimal strings
	hexStr := fmt.Sprintf("%x", hashBytes)

	// Convert hexadecimal string to integer in base 25 format
	num, ok := new(big.Int).SetString(hexStr, 16)
	if !ok {
		return -1, errors.New("Unable to convert hexadecimal string to integer")
	}
	base25Num := new(big.Int).Set(num)
	base := big.NewInt(25)
	zero := big.NewInt(0)
	result := big.NewInt(0)
	for base25Num.Cmp(zero) > 0 {
		remainder := new(big.Int).Mod(base25Num, base)
		result.Mul(result, base).Add(result, remainder)
		base25Num.Div(base25Num, base)
	}

	// Take a mold for 2 * * 25-1
	mod := big.NewInt(1 << 25)
	mod.Sub(mod, big.NewInt(1))
	finalResult := new(big.Int).Mod(result, mod)

	// Convert the final result to int64 type and return it
	return finalResult.Int64(), nil
}

// GenerateFullGlobalUuid A universal method for generating unique IDs (combinations of numbers and uppercase letters) based on length.
func GenerateFullGlobalUuid(length int) string {
	rand.Seed(time.Now().UnixNano())
	// Used to store generated user IDs
	var uniqueID strings.Builder
	// Generate characters of a specified number of digits in a loop
	for i := 0; i < length; i++ {
		// Randomly select a character from the character set
		randomIndex := rand.Intn(len(uuidFullChar))
		uniqueID.WriteRune(rune(uuidFullChar[randomIndex]))
	}
	return uniqueID.String()
}
