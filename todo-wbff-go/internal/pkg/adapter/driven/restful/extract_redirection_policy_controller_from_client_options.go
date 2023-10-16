package driven_adapter_restful

func extractRedirectionPolicyControllerFromClientOptions(options *[]ClientOption) (func(int, string, map[string][]string) bool, *[]ClientOption) {
	var controller func(int, string, map[string][]string) bool
	remainingOptions := make([]ClientOption, 0)
	for _, option := range *options {
		switch o := option.(type) {
		case RedirectionPolicyControllerOption:
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
