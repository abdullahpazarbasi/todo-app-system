package infrastructure_adapters_restful

import drivenAppPortsRestful "todo-app-wbff/internal/pkg/app/ports/driven/restful"

func extractPathParametersFromClientOptions(options *[]drivenAppPortsRestful.ClientOption) (*map[string]string, *[]drivenAppPortsRestful.ClientOption) {
	pathParameters := make(map[string]string)
	remainingOptions := make([]drivenAppPortsRestful.ClientOption, 0)
	for _, option := range *options {
		switch o := option.(type) {
		case drivenAppPortsRestful.ResourcePathParameterOption:
			pathParameters[o.Placeholder()] = o.Value()
		case drivenAppPortsRestful.ResourcePathParametersOption:
			for p, v := range o.Map() {
				pathParameters[p] = v
			}
		default:
			remainingOptions = append(remainingOptions, option)
		}
	}

	return &pathParameters, &remainingOptions
}
