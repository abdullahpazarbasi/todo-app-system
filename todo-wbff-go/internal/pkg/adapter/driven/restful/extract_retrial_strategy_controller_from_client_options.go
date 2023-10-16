package driven_adapter_restful

func extractRetrialStrategyControllerFromClientOptions(options *[]ClientOption) (func(Exchange) bool, *[]ClientOption) {
	var controller func(Exchange) bool
	remainingOptions := make([]ClientOption, 0)
	for _, option := range *options {
		switch o := option.(type) {
		case RetrialStrategyControllerOption:
			controller = o.Controller()
		default:
			remainingOptions = append(remainingOptions, option)
		}
	}
	if controller == nil {
		controller = func(lastExchange Exchange) (breaking bool) {
			if lastExchange.Response().StatusCode() == 429 {
				breaking = false
				return
			}

			breaking = true
			return
		}
	}

	return controller, &remainingOptions
}
