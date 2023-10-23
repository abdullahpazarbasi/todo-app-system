package driven_adapter_restful

func extractPathParametersFromClientOptions(options *[]ClientOption) (map[string]string, *[]ClientOption) {
	pathParameters := make(map[string]string)
	remainingOptions := make([]ClientOption, 0)
	for _, option := range *options {
		switch o := option.(type) {
		case ResourcePathParameterOption:
			pathParameters[o.Placeholder()] = o.Value()
		case ResourcePathParametersOption:
			for p, v := range o.Map() {
				pathParameters[p] = v
			}
		default:
			remainingOptions = append(remainingOptions, option)
		}
	}

	return pathParameters, &remainingOptions
}
