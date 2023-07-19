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
	if controller == nil {
		controller = func(statusCode int, targetURL string, header map[string][]string) (redirectability bool) {
			if targetURL == "" {
				redirectability = false
				return
			}
			if statusCode < 300 && statusCode > 399 {
				redirectability = false
				return
			}

			redirectability = true
			return
		}
	}

	return controller, &remainingOptions
}
