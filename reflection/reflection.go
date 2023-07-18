package reflection

import (
	"fmt"
	"reflect"
)

func Walk(x interface{}, fn func(input string)) {
	val := getValue(x)
	fmt.Printf("Reflection Valueof %#v\n", val) // only for debug purpose

	walkValue := func(value reflect.Value) {
		Walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for value, ok := val.Recv(); ok; value, ok = val.Recv() {
			walkValue(value)
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		// extract the underlying value before using Elem()
		val = val.Elem()
	}

	return val
}
