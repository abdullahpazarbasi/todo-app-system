package driven_adapter_restful

import (
	"fmt"
)

func extractHTTPErrorHandlingStrategyControllerFromClientOptions(
	options *[]ClientOption,
) (
	func(lastExchange Exchange, exchangeError error) error,
	*[]ClientOption,
) {
	var controller func(lastExchange Exchange, cause error) error
	remainingOptions := make([]ClientOption, 0)
	for _, option := range *options {
		switch o := option.(type) {
		case HTTPErrorHandlingStrategyControllerOption:
			controller = o.Controller()
		default:
			remainingOptions = append(remainingOptions, option)
		}
	}
	if controller == nil {
		controller = func(lastExchange Exchange, exchangeError error) error {
			if exchangeError != nil {
				return nil
			}
			lastResponse := lastExchange.Response()
			if lastResponse.IsStatusError() {
				responseModel, err := lastResponse.DecodeModel()
				if err != nil {
					return err
				}

				return fmt.Errorf(
					"http error %d: %v",
					lastResponse.StatusCode(),
					responseModel,
				)
			}

			return nil
		}
	}

	return controller, &remainingOptions
}
