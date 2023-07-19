package infrastructure_adapters_model

import (
	"context"
	"fmt"
	"reflect"
	drivenAppPortsModel "todo-app-wbff/internal/pkg/app/ports/driven/model"
)

type modelNormalizer struct{}

func NewModelNormalizer() *modelNormalizer {
	return &modelNormalizer{}
}

func (n *modelNormalizer) NewContextWith(parentContext context.Context) context.Context {
	return context.WithValue(parentContext, drivenAppPortsModel.ModelNormalizerKey{}, n)
}

func (n *modelNormalizer) Normalize(sourceModel interface{}) interface{} {
	return normalize(sourceModel)
}

func (n *modelNormalizer) Denormalize(targetModelReference interface{}, sourceModelReference interface{}) error {
	return denormalize(targetModelReference, sourceModelReference)
}

func normalize(sourceModel interface{}) interface{} {
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
				targetModel[sourceModelFieldName] = normalize(sourceModelFieldValue.Interface())
			default:
				targetModel[sourceModelFieldName] = sourceModelFieldValue.Interface()
			}
		}

		return targetModel
	case reflect.Slice:
		targetModel := make([]interface{}, 0)

		for i := 0; i < sourceModelValue.Len(); i++ {
			sourceModelElementValue := sourceModelValue.Index(i)
			targetModel = append(targetModel, normalize(sourceModelElementValue.Interface()))
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

func denormalize(targetModelReference interface{}, sourceModelReference interface{}) error {
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
				err := denormalize(targetModelFieldValue.Addr().Interface(), sourceModelFieldContent)
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
			err := denormalize(targetModelElementReference.Interface(), sourceModelElementContent)
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

func findField(model reflect.Value, fieldName string) (reflect.Value, bool) {
	var modelField reflect.StructField
	var fallback reflect.Value
	for i := 0; i < model.NumField(); i++ {
		modelField = model.Type().Field(i)
		if modelField.Tag.Get("field") == fieldName {
			return model.Field(i), true
		} else if modelField.Name == fieldName {
			fallback = model.Field(i)
		}
	}

	return fallback, fallback.IsValid()
}
