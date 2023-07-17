package infrastructure_adapters_restful

import drivenAppPortsRestful "todo-app-wbff/internal/pkg/app/ports/driven/restful"

func extractRedirectionPolicyControllerFromClientOptions(options *[]drivenAppPortsRestful.ClientOption) (func(int, string, map[string][]string) bool, *[]drivenAppPortsRestful.ClientOption) {
	var controller func(int, string, map[string][]string) bool
	remainingOptions := make([]drivenAppPortsRestful.ClientOption, 0)
	for _, option := range *options {
		switch o := option.(type) {
		case drivenAppPortsRestful.RedirectionPolicyControllerOption:
			controller = o.Controller()
		default:
			remainingOptions = append(remainingOptions, option)
		}
	}

	return controller, &remainingOptions
}
