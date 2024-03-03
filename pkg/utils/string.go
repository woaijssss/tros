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
