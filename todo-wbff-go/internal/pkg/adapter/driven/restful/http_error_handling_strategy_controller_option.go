package driven_adapter_restful

type HTTPErrorHandlingStrategyControllerOption interface {
	ClientOption
	Controller() func(lastExchange Exchange, cause error) error
}

type httpErrorHandlingStrategyControllerOption struct {
	controller func(lastExchange Exchange, cause error) error
}

func NewHTTPErrorHandlingStrategyControllerOption(
	controller func(lastExchange Exchange, cause error) error,
) HTTPErrorHandlingStrategyControllerOption {
	return &httpErrorHandlingStrategyControllerOption{
		controller: controller,
	}
}

func (o *httpErrorHandlingStrategyControllerOption) IsRestfulClientOption() bool {
	return true
}

func (o *httpErrorHandlingStrategyControllerOption) Controller() func(lastExchange Exchange, cause error) error {
	return o.controller
}
