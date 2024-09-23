package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func StructToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	val := reflect.ValueOf(obj)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()

	fmt.Println(typ)

	for i := 0; i < val.NumField(); i++ {
		fieldName := typ.Field(i).Name
		fieldValueKind := val.Field(i).Kind()
		var fieldValue interface{}

		if fieldValueKind == reflect.Struct {
			fieldValue = StructToMap(val.Field(i).Interface())
		} else {
			fieldValue = val.Field(i).Interface()
		}

		result[fieldName] = fieldValue
	}

	return result
}

func TypeConverter[T any](data any) (T, error) {
	var result T
	dataJson, err := json.Marshal(&data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(dataJson, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
