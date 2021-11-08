package utils

import (
	"reflect"
	"strings"
)

func JoinStringPtr(elems []*string, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return *elems[0]
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(*elems[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(*elems[0])
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(*s)
	}
	return b.String()
}

func SafeString(ptr *string)string {
	if ptr == nil {
		return ""
	} else {
		return *ptr
	}
}

func SafeInt32(ptr *int32) int32 {
	if ptr == nil {
		return 0
	} else {
		return *ptr
	}
}

func SafeFloat32(ptr *float32) float32 {
	if ptr == nil {
		return 0
	} else {
		return *ptr
	}
}

func SafeBool(ptr *bool, def bool) bool {
	if ptr == nil {
		return def
	} else {
		return *ptr
	}
}

func CheckQueryKeys(query map[string]interface{}, obj interface{}) bool {
	if reflect.ValueOf(obj).Kind() != reflect.Struct {
		return false
	}
	val := reflect.TypeOf(obj)
	fLen := val.NumField()
	for key, _ := range query {
		exist := false
		for i := 0; i < fLen; i++ {
			if val.Field(i).Name == key {
				exist = true
				break
			}
		}
		if !exist {
			return false
		}
	}
	return true
}
