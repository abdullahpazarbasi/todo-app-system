package infrastructure_adapters_restful

import drivenAppPortsRestful "todo-app-wbff/internal/pkg/app/ports/driven/restful"

func extractHTTPErrorHandlingStrategyControllerFromClientOptions(options *[]drivenAppPortsRestful.ClientOption) (func(lastExchange drivenAppPortsRestful.Exchange, cause error) drivenAppPortsRestful.Error, *[]drivenAppPortsRestful.ClientOption) {
	var controller func(lastExchange drivenAppPortsRestful.Exchange, cause error) drivenAppPortsRestful.Error
	remainingOptions := make([]drivenAppPortsRestful.ClientOption, 0)
	for _, option := range *options {
		switch o := option.(type) {
		case drivenAppPortsRestful.HTTPErrorHandlingStrategyControllerOption:
			controller = o.Controller()
		default:
			remainingOptions = append(remainingOptions, option)
		}
	}
	if controller == nil {
		controller = func(lastExchange drivenAppPortsRestful.Exchange, cause error) drivenAppPortsRestful.Error {
			rs := lastExchange.Response()
			if rs.IsStatusError() {
				e, err := rs.DecodeModel()
				e["status_code"] = rs.StatusCode()
				if err != nil {
					e["message"] = err.Error()
				}

				return denormalize(e)
			}

			return nil
		}
	}

	return controller, &remainingOptions
}
