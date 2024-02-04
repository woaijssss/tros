package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/siddontang/go/bson"
	"reflect"
	"strconv"
)

func Interface2String(inter interface{}, param ...int) string {
	switch inter.(type) {
	case string:
		return inter.(string)
	case int, int32:
		return strconv.FormatInt(int64(inter.(int)), 10)
	case int64:
		return strconv.FormatInt(inter.(int64), 10)
	case float32:
		return strconv.FormatFloat(float64(inter.(float32)), 'f', param[0], 32)
	case float64:
		perc := 2
		if param != nil {
			perc = param[0]
		}
		return strconv.FormatFloat(inter.(float64), 'f', perc, 64)
	}
	return ""
}

// IsNil 判断interface是否为nil
func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

// StructToMapUseJson 结构体转KV
func StructToMapUseJson(st interface{}) map[string]interface{} {
	var ret map[string]interface{}
	jsonStr, _ := json.Marshal(st)
	json.Unmarshal(jsonStr, &ret)
	return ret
}

// StructToMap Deprecated：未经过实际测试，可能存在无法预知的问题
func StructToMap(st interface{}) map[string]interface{} {
	types := reflect.TypeOf(st)
	values := reflect.ValueOf(st)

	var data = make(map[string]interface{})
	for i := 0; i < types.NumField(); i++ {
		k := types.Field(i).Name
		var v = values.Field(i).Interface()
		if !IsNil(v) {
			switch v.(type) {
			case int:
				data[k] = v
			case *string:
				vv := v.(*string)
				data[k] = *vv
			case string:
				data[k] = v
			}
		}
	}

	return data
}

// DeepCopy Deprecated：未经过实际测试，可能存在无法预知的问题
func DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v)
		}

		return newSlice
	} else if valueMap, ok := value.(bson.M); ok {
		newMap := make(bson.M)
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}
	}
	return value
}

// StrictStructCopy 严格的struct拷贝
// Deprecated：未经过实际测试，可能存在无法预知的问题
func StrictStructCopy(src interface{}, dst interface{}) {
	sval := reflect.ValueOf(src).Elem()
	dval := reflect.ValueOf(dst).Elem()

	for i := 0; i < sval.NumField(); i++ {
		value := sval.Field(i)
		name := sval.Type().Field(i).Name

		dvalue := dval.FieldByName(name)
		if dvalue.IsValid() == false {
			continue
		}
		dvalue.Set(value)
	}
}

// StructCopyUseJson 利用json做中间值进行拷贝
func StructCopyUseJson(src interface{}, dst interface{}) {
	//srcJson, _ := json.Marshal(src)
	//_ = json.Unmarshal(srcJson, dst)

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonBytes, _ := json.Marshal(src)
	_ = json.Unmarshal(jsonBytes, dst)
}

// StructCopyUseReflect 利用反射进行拷贝
// 已知问题：对组合型结构体拷贝存在问题
func StructCopyUseReflect(dst interface{}, src interface{}) (err error) {
	// 防止意外panic
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%v", e))
		}
	}()

	dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	srcType, srcValue := reflect.TypeOf(src), reflect.ValueOf(src)

	// dst必须结构体指针类型
	if dstType.Kind() != reflect.Ptr || dstType.Elem().Kind() != reflect.Struct {
		return errors.New("dst type should be a struct pointer")
	}

	// src必须为结构体或者结构体指针
	if srcType.Kind() == reflect.Ptr {
		srcType, srcValue = srcType.Elem(), srcValue.Elem()
	}
	if srcType.Kind() != reflect.Struct {
		return errors.New("src type should be a struct or a struct pointer")
	}

	// 取具体内容
	dstType, dstValue = dstType.Elem(), dstValue.Elem()

	// 属性个数
	propertyNums := dstType.NumField()

	for i := 0; i < propertyNums; i++ {
		// 属性
		property := dstType.Field(i)
		// 待填充属性值
		propertyValue := srcValue.FieldByName(property.Name)

		// 无效，说明src没有这个属性 || 属性同名但类型不同
		if !propertyValue.IsValid() || property.Type != propertyValue.Type() {
			continue
		}

		if dstValue.Field(i).CanSet() {
			dstValue.Field(i).Set(propertyValue)
		}
	}

	return nil
}
