package core_adapter

import "reflect"

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
