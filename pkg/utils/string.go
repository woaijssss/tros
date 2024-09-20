package utils

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

func String2Int32(s string) int32 {
	n, err := strconv.Atoi(s)
	if err != nil {
		return math.MaxInt32 // 错误值
	}
	return int32(n)
}

func String2Int64(s string) int64 {
	n, err := strconv.Atoi(s)
	if err != nil {
		return math.MaxInt64 // 错误值
	}
	return int64(n)
}

func String2Float64WithDefaultError(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return -1 // 错误值
	}
	return f
}

func String2Float64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return math.MaxFloat64 // 错误值
	}
	return f
}

// SetEmpty 不规则空字符串，设置为标准空字符串
func SetEmpty(s string) string {
	if s == "<nil>" || // "<nil>"针对于 godbx 仓库，date类型的NULL值
		s == "" {
		return ""
	}

	return s
}

// IsAllWhiteSpace 检查是否是空白符或只包含空格
func IsAllWhiteSpace(s string) bool {
	for _, c := range s {
		if !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// JoinStringArray 将字符串数组，用指定分隔符连接
func JoinStringArray(arr []string, sep string) string {
	return strings.Join(arr, sep)
}
