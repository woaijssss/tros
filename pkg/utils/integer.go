package utils

import (
	"fmt"
	"github.com/unknwon/com"
	"strconv"
	"strings"
)

func InArray(num int, slice []int) bool {
	for _, v := range slice {
		if num == v {
			return true
		}
	}
	return false
}

// Intersect 交集
func Intersect(sliceA []int, sliceB []int) []int {
	var result []int

	for _, a := range sliceA {
		for _, b := range sliceB {
			if a == b {
				result = append(result, a)
			}
		}
	}

	return result
}

// Subtract 差集
func Subtract(sa []int, sb []int) []int {
	var result []int
	for _, a := range sa {
		isExist := false
		for _, b := range sb {
			if a == b {
				isExist = true
				break
			}
		}
		if isExist == false {
			result = append(result, a)
		}
	}
	return result
}

func ToInt(i interface{}) int {
	switch i.(type) {
	case int:
		return i.(int)
	case int8:
		return int(i.(int8))
	case int16:
		return int(i.(int16))
	case int32:
		return int(i.(int32))
	case int64:
		return int(i.(int64))
	case float32:
		return int(i.(float32))
	case float64:
		return int(i.(float64))
	}
	panic(fmt.Sprintf("存在无法转换的数据类型: %v", i))
}

func FindIndex(ids []int, x int) int {
	for i, j := range ids {
		if j == x {
			return i
		}
	}
	return -1
}

func Join(arr []int, split string) string {
	var strArr []string
	for _, i := range arr {
		strArr = append(strArr, strconv.Itoa(i))
	}
	return strings.Join(strArr, split)
}

func ToIntSlice(str string, separator string) (rst []int) {
	if str == "" {
		return []int{}
	}
	strArr := strings.Split(str, separator)
	for _, str := range strArr {
		rst = append(rst, com.StrTo(str).MustInt())
	}
	return
}

func GetIntOrDefault(i *int, dft int) int {
	if i == nil {
		return dft
	}
	return *i
}

func GetIntNotZeroOrDefault(i *int, dft int) int {
	if i == nil {
		return dft
	}
	if *i == 0 {
		return dft
	}
	return *i
}
