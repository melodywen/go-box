package collections

import (
	"reflect"
)

// ArrFun struct arr fun
type ArrFun struct {
}

// Wrap interface to slice
func (_ ArrFun) Wrap(value interface{}) []interface{} {
	if kind := reflect.TypeOf(value).Kind(); kind == reflect.Slice || kind == reflect.Array {
		return value.([]interface{})
	}
	return []interface{}{value}
}
