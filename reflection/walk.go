package reflection

import "reflect"

func Walk(x interface{}, fn func(str string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.Chan:
		for value, ok := val.Recv(); ok; value, ok = val.Recv() {
			Walk(value.Interface(), fn)
		}
	case reflect.Func:
		funcResult := val.Call(nil)	
		for _, res := range funcResult {
			Walk(res.Interface(), fn)
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
		field := val.MapIndex(key)
			Walk(field.Interface(), fn)	
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			Walk(field.Interface(), fn)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			field := val.Index(i)
			Walk(field.Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}

