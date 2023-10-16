package core_adapter

import (
	"fmt"
	"reflect"
)

func denormalizeModel(targetModelReference interface{}, sourceModelReference interface{}) error {
	targetModelReferenceValue := reflect.ValueOf(targetModelReference)
	if targetModelReferenceValue.Kind() != reflect.Ptr {
		return fmt.Errorf("target model reference must be a pointer")
	}
	targetModelValue := targetModelReferenceValue.Elem()

	var sourceModel interface{}
	switch smr := sourceModelReference.(type) {
	case *map[string]interface{}:
		sourceModel = *smr
	case *[]interface{}:
		sourceModel = *smr
	default:
		return fmt.Errorf("input must be a map of any or a slice of any, but %T given", sourceModelReference)
	}

	switch sm := sourceModel.(type) {
	case map[string]interface{}:
		for sourceModelFieldKey, sourceModelFieldContent := range sm {
			targetModelFieldValue, existent := findField(targetModelValue, sourceModelFieldKey)
			if !existent {
				continue
			}

			sourceModelFieldContentType := reflect.TypeOf(sourceModelFieldContent)
			if sourceModelFieldContentType.Kind() == reflect.Map {
				err := denormalizeModel(targetModelFieldValue.Addr().Interface(), sourceModelFieldContent)
				if err != nil {
					return err
				}
				continue
			}

			if !targetModelFieldValue.CanSet() {
				continue
			}

			if targetModelFieldValue.Type() != sourceModelFieldContentType {
				continue
			}

			targetModelFieldValue.Set(reflect.ValueOf(sourceModelFieldContent))
		}
	case []interface{}:
		if !targetModelValue.CanSet() {
			return nil
		}

		targetModelElementType := targetModelValue.Type().Elem()
		for _, sourceModelElementContent := range sm {
			targetModelElementReference := reflect.New(targetModelElementType)
			err := denormalizeModel(targetModelElementReference.Interface(), sourceModelElementContent)
			if err != nil {
				return err
			}
			targetModelValue.Set(reflect.Append(targetModelValue, targetModelElementReference.Elem()))
		}
	default:
		return fmt.Errorf("input must be a map of any or a slice of any, but %T given", sourceModel)
	}

	return nil
}
