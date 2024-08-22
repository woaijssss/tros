package utils

import (
	"math"
	"strconv"
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
