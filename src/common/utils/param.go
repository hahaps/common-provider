package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

const (
	DefaultLimit int = 50
)

const (
	Float64 reflect.Kind = reflect.Float64
	Int     reflect.Kind = reflect.Int
	String  reflect.Kind = reflect.String
	Bool    reflect.Kind = reflect.Bool
	Array   reflect.Kind = reflect.Array
	Slice   reflect.Kind = reflect.Slice
	Map     reflect.Kind = reflect.Map
)

type Scheme struct {
	Param    string
	Required bool
	Type     reflect.Kind
	Default  interface{}
}

func CheckParam(params map[string]interface{}, schemes []Scheme) (map[string]interface{}, error) {
	var err error
	for _, scheme := range schemes {
		param, ok := params[scheme.Param]
		if !ok {
			params[scheme.Param] = scheme.Default
			param = scheme.Default
		}
		if scheme.Required && param == nil {
			return params, errors.New(
				fmt.Sprintf("Param[%v] should be specified",
					scheme.Param))
		}
		if scheme.Required && scheme.Type == reflect.String {
			if param == nil || param.(string) == "" {
				return params, errors.New(
					fmt.Sprintf("Param[%v] could not be empty",
						scheme.Param))
			}
		}
		pType := reflect.ValueOf(param).Kind()
		if param != nil && pType != scheme.Type {
			if pType == Float64 && scheme.Type == Int {
				params[scheme.Param], err = strconv.Atoi(fmt.Sprint(reflect.ValueOf(param)))
				if err != nil {
					return params, errors.New("bad params, need int, but float specified")
				}
			} else {
				return params, errors.New(
					fmt.Sprintf("Param[%v] should be %v, but %v specified",
						scheme.Param, scheme.Type, pType))
			}
		}
	}
	return params, nil
}
