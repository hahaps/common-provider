// @Author Bruce<lixipengs@qq.com>

package model

import (
	"crypto/md5"
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

type BaseModel interface {
	SetIndex()
	SetChecksum()
	GetIndex() string
	GetChecksum() string
	CheckRequired() (bool, string)
}

func String2Bytes(s *string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Checksum(val *string) string {
	return string(md5.New().Sum(String2Bytes(val)))
}

func GetValFromKeys(obj interface{}, k string) (str string) {
	keys := strings.Split(k, ",")
	val := reflect.ValueOf(obj).Elem()
	vLen := len(keys)
	for i, key := range keys {
		key = strings.TrimSpace(key)
		if i + 1 < vLen {
			str += fmt.Sprint(val.FieldByName(key)) + "-"
		} else {
			str += fmt.Sprint(val.FieldByName(key))
		}
	}
	return str
}
