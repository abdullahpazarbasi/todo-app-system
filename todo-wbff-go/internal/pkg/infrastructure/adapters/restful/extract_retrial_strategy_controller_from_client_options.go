package infrastructure_adapters_restful

import drivenAppPortsRestful "todo-app-wbff/internal/pkg/app/ports/driven/restful"

func extractRetrialStrategyControllerFromClientOptions(options *[]drivenAppPortsRestful.ClientOption) (func(drivenAppPortsRestful.Exchange) bool, *[]drivenAppPortsRestful.ClientOption) {
	var controller func(drivenAppPortsRestful.Exchange) bool
	remainingOptions := make([]drivenAppPortsRestful.ClientOption, 0)
	for _, option := range *options {
		switch o := option.(type) {
		case drivenAppPortsRestful.RetrialStrategyControllerOption:
			controller = o.Controller()
		default:
			remainingOptions = append(remainingOptions, option)
		}
	}

	return controller, &remainingOptions
}
