package utils

import (
	"errors"
	"fmt"
	"reflect"
)

const (
	DefaultLimit int = 50
)

type Scheme struct {
	Param    string
	Required bool
	Type     reflect.Kind
	Default  interface{}
}

func CheckParam(params map[string]interface{}, schemes []Scheme) (map[string]interface{}, error) {
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
		if reflect.ValueOf(param).Kind() != scheme.Type {
			return params, errors.New(
				fmt.Sprintf("Param[%v] should be %v",
					scheme.Param, scheme.Type))
		}
	}
	return params, nil
}
