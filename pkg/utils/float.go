package utils

import "math"

// FloatRetain 浮点型保留几位小数
func FloatRetain(f float64, point int32) float64 {
	// 将float64值乘以10^6
	scaled := f * math.Pow(10, float64(point))

	// 使用math.Round进行舍入
	rounded := math.Round(scaled)

	// 将舍入后的值除以10^6，以恢复原始的比例
	return rounded / math.Pow(10, float64(point))
}
