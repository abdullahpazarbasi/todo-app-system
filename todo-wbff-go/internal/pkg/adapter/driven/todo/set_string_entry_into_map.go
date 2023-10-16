package driven_adapter_todo

func setStringEntryIntoMap(target *map[string]interface{}, key string, value string) {
	if value == "" {
		return
	}
	if target == nil {
		return
	}
	(*target)[key] = value
}
