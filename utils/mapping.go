package utils

import (
	"reflect"
)

func Mapping(s interface{}, v interface{}) {

}

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func StructToSlice(obj interface{}, tag string, filterKey ...string) []interface{} {
	slice := make([]interface{}, 0)
	v := reflect.ValueOf(obj)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	typ := v.Type()
re:
	for i := 0; i < v.NumField(); i++ {
		fi := typ.Field(i)
		if fi.Tag.Get(tag) == "" {
			continue
		}
		if len(filterKey) > 0 {
			for j := 0; j < len(filterKey); j++ {
				if fi.Tag.Get(tag) == filterKey[j] {
					continue re
				}
			}
		}

		slice = append(slice, v.Field(i).Interface())
	}
	return slice
}
