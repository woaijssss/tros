package utils

import (
	"math"
	"strconv"
	"unicode"
)

func String2Int64(input string) int64 {
	n, err := strconv.Atoi(input)
	if err != nil {
		return math.MaxInt64 // 错误值
	}
	return int64(n)
}

// SetEmpty 不规则空字符串，设置为标准空字符串
func SetEmpty(s string) string {
	if s == "<nil>" || // "<nil>"针对于 daog 仓库，date类型的NULL值
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
