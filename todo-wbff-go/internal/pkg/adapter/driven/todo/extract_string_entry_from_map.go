package driven_adapter_todo

func extractStringEntryFromMap(source map[string]interface{}, key string) string {
	rawValue, existent := source[key]
	if !existent {
		return ""
	}
	value, ok := rawValue.(string)
	if !ok {
		return ""
	}

	return value
}
