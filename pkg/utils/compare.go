package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func CompareValueStructs[T any](old T, new T, ignoreFields []string, prefix string) (map[string]any, error) {
	changes := make(map[string]any)
	ignoreMap := make(map[string]bool)

	for _, field := range ignoreFields {
		ignoreMap[field] = true
	}

	oldVal := reflect.ValueOf(old)
	newVal := reflect.ValueOf(new)

	if oldVal.Kind() == reflect.Ptr {
		oldVal = oldVal.Elem()
	}
	if newVal.Kind() == reflect.Ptr {
		newVal = newVal.Elem()
	}

	if oldVal.Kind() != reflect.Struct || newVal.Kind() != reflect.Struct {
		return map[string]any{}, fmt.Errorf("only structs are supported")
	}

	for i := 0; i < oldVal.NumField(); i++ {
		field := oldVal.Type().Field(i)
		fieldName := field.Name
		fullFieldName := fieldName
		// TODO handle prefix

		if ignoreMap[fullFieldName] {
			continue
		}

		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			jsonTag = fieldName
		} else {
			jsonTag = strings.Split(jsonTag, ",")[0]
		}

		oldValue := oldVal.Field(i).Interface()
		newValue := newVal.Field(i).Interface()

		if field.Type.Kind() == reflect.Struct {

			nestedChanges, err := CompareValueStructs(oldValue, newValue, ignoreFields, fullFieldName)
			if err != nil {
				return map[string]any{}, err
			}
			if len(nestedChanges) > 0 {
				changes[jsonTag] = nestedChanges
			}
		} else {
			if !reflect.DeepEqual(oldValue, newValue) {
				changes[jsonTag] = map[string]any{
					"old_value": oldValue,
					"new_value": newValue,
				}
			}
		}
	}

	return changes, nil
}

func UniqMapBy[T any, U comparable, Slice ~[]T](collection Slice, iteratee func(item T, idx int) U) []U {
	result := make([]U, 0, len(collection))
	seen := make(map[U]struct{}, len(collection))

	for i := range collection {
		key := iteratee(collection[i], i)

		if _, ok := seen[key]; ok {
			continue
		}

		seen[key] = struct{}{}
		result = append(result, key)
	}

	return result
}
