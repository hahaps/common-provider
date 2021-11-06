package utils

import (
	"errors"
	"fmt"
	"reflect"
)

type Scheme struct {
	Param    string
	Required bool
	Type reflect.Kind
}

func CheckParam(params map[string]interface{}, schemes []Scheme) error {
	for _, scheme := range schemes {
		param, ok := params[scheme.Param]
		if !ok {
			return errors.New(
				fmt.Sprintf("Param[%v] should be specifid",
					scheme.Param))
		}
		if reflect.ValueOf(param).Kind() != scheme.Type {
			return errors.New(
				fmt.Sprintf("Param[%v] should be %v",
					scheme.Param, scheme.Type))
		}
		if !scheme.Required {
			return nil
		}
		if scheme.Type == reflect.String {
			if param.(string) == "" {
				return errors.New(
					fmt.Sprintf("Param[%v] could not be empty",
						scheme.Param))
			}
		}
	}
	return nil
}
