package core_adapter

import "reflect"

func resolveFieldName(f *reflect.StructField) string {
	tagValue, tagExistence := f.Tag.Lookup("field")
	if tagExistence && len(tagValue) > 0 {
		return tagValue
	}

	return f.Name
}
