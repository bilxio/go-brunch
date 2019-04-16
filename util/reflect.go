package util

import (
	"reflect"
)

// GetPrototype 从一个 interface{} 中得到其原始 Type 信息
func GetPrototype(v interface{}) reflect.Type {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Slice ||
		t.Kind() == reflect.Array ||
		t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}

// GetTypeName 从一个 interface{} 中得到其 Type Name
func GetTypeName(v interface{}) string {
	if t, ok := v.(reflect.Type); ok {
		return t.Name()
	}

	var t2 = reflect.TypeOf(v)

	// 如果是指针或者切片，则取其元素的类型
	for t2.Kind() == reflect.Ptr || t2.Kind() == reflect.Slice {
		t2 = t2.Elem()
	}
	return t2.Name()
}

// NewSliceValue 根据一个 struct 类型，新建一个该类型的 Slice 类型的 Value
func NewSliceValue(t reflect.Type) reflect.Value {
	var sliceType reflect.Type
	if t.Kind() == reflect.Slice ||
		t.Kind() == reflect.Array ||
		t.Kind() == reflect.Ptr {
		sliceType = reflect.SliceOf(t.Elem())
	} else {
		sliceType = reflect.SliceOf(t)
	}
	return reflect.New(sliceType)
}
