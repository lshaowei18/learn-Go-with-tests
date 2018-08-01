package reflection

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	values := getValue(x)
	for _, value := range values {
		if value.Kind() == reflect.String {
			fn(value.String())
		}
	}
}

func getValue(x interface{}) []reflect.Value {
	var slice []reflect.Value
	values := reflect.ValueOf(x)
	if values.Kind() == reflect.Ptr {
		values = values.Elem()
	}
	switch valType := values.Kind(); valType {
	case reflect.Struct:
		for i := 0; i < values.NumField(); i++ {
			val := values.Field(i)
			slice = append(slice, val)
		}
	case reflect.Slice:
		for i := 0; i < values.Len(); i++ {
			val := values.Index(i)
			for j := 0; j < val.NumField(); j++ {
				v := val.Field(j)
				slice = append(slice, v)
			}
		}
	}
	return slice
}

func Walk1(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	walkValue := func(value reflect.Value) {
		Walk1(value.Interface(), fn)
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
	}
}
