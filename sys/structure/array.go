package structure

import (
	"fmt"
	"strings"
)

type ArrayElemType interface {
	byte | int8 | int16 | int32 | int64 | int | string
}

// Array 列表
type Array[T ArrayElemType] interface {
	// RemoveDuplicates 去重
	RemoveDuplicates()
	// Array 返回实例
	Array() []T
	// Join 逗号分隔（默认是逗号分隔）
	Join(sep string) string
	//Exist(target T) bool
	//AddOne(elem T)
	//Add(elems ...T)
	//Len() int
	//Cap() int
	//AllElements() []T
}

func NewArray[T ArrayElemType]() Array[T] {
	return &array[T]{
		list: []T{},
	}
}

func NewFromArray[T ArrayElemType](arr []T) Array[T] {
	return &array[T]{
		list: arr,
	}
}

type array[T ArrayElemType] struct {
	list []T
}

func (a *array[T]) Join(sep string) string {
	strList := make([]string, len(a.list))
	for i, v := range a.list {
		strList[i] = fmt.Sprintf("%d", v)
	}
	if sep == "" {
		sep = "," // 默认是逗号分隔
	}
	return strings.Join(strList, ",")
}

func (a *array[T]) Array() []T {
	return a.list
}

func (a *array[T]) RemoveDuplicates() {
	uniqueMap := make(map[T]bool)
	var result []T

	for _, value := range a.list {
		if v, _ := uniqueMap[value]; !v {
			uniqueMap[value] = true
			result = append(result, value)
		}
	}

	fmt.Println("a.list: ", a.list)
	a.list = result
	fmt.Println("a.list: ", a.list)
}
