package main

import (
	"fmt"
	"reflect"
)

func walk(x interface{}, fn func(input string)) error {
	val := reflect.ValueOf(x)
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}
	switch val.Kind() {
	// base case		-> call callback with string
	case reflect.String:
		fn(val.String())
	// pointer case -> convert pointer to element, recurse
	case reflect.Pointer:
		walkValue(val.Elem())
	// slice case   -> recurse over every element in slice
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	// struct case   -> recurse over every element in slice
	case reflect.Struct:
		fieldCount := val.NumField()
		for i := 0; i < fieldCount; i++ {
			walkValue(val.Field(i))
		}
	// error case
	default:
		return fmt.Errorf("Received interface is not supported")
	}
	return nil
}
