package core_adapter

import (
	"fmt"
	"reflect"
)

func normalizeModel(sourceModel interface{}) interface{} {
	sourceModelValue := reflect.ValueOf(sourceModel)
	if sourceModelValue.Kind() == reflect.Ptr {
		sourceModelValue = sourceModelValue.Elem()
	}

	sourceModelKind := sourceModelValue.Kind()
	switch sourceModelKind {
	case reflect.Struct:
		targetModel := make(map[string]interface{})

		sourceModelType := sourceModelValue.Type()
		for i := 0; i < sourceModelType.NumField(); i++ {
			sourceModelFieldValue := sourceModelValue.Field(i)
			if !sourceModelFieldValue.CanInterface() {
				continue
			}
			sourceModelField := sourceModelType.Field(i)
			sourceModelFieldName := resolveFieldName(&sourceModelField)
			sourceModelFieldKind := sourceModelField.Type.Kind()
			switch sourceModelFieldKind {
			case reflect.Struct, reflect.Slice:
				targetModel[sourceModelFieldName] = normalizeModel(sourceModelFieldValue.Interface())
			default:
				targetModel[sourceModelFieldName] = sourceModelFieldValue.Interface()
			}
		}

		return targetModel
	case reflect.Slice:
		targetModel := make([]interface{}, 0)

		for i := 0; i < sourceModelValue.Len(); i++ {
			sourceModelElementValue := sourceModelValue.Index(i)
			targetModel = append(targetModel, normalizeModel(sourceModelElementValue.Interface()))
		}

		return targetModel
	default:
		panic(fmt.Errorf("kind of source model must be struct or slice, but %s given", sourceModelKind))
	}
}

func resolveFieldName(f *reflect.StructField) string {
	tagValue, tagExistence := f.Tag.Lookup("field")
	if tagExistence && len(tagValue) > 0 {
		return tagValue
	}

	return f.Name
}
