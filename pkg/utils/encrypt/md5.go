package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"strings"
)

// EncodeMD5AsEmpty md5 encryption
func EncodeMD5AsEmpty() string {
	value := uuid.Must(uuid.NewV4(), nil).String()
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

// EncodeMD5 md5 encryption
func EncodeMD5Byte(value string) []byte {
	m := md5.New()
	m.Write([]byte(value))
	b := m.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(dst, b)
	fmt.Println("string(dst): ", string(dst))
	return dst
}

// EncodeMD5 md5 encryption
func EncodeMD5Upper(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	md5Str := hex.EncodeToString(m.Sum(nil))
	return strings.ToUpper(md5Str)
}

func GetMd5Sign(id, key, sign string, timeStamp int64) string {
	s := strconv.FormatInt(timeStamp, 10)
	str := id + key + s
	str = str + "_" + sign
	return EncodeMD5(str)
}
