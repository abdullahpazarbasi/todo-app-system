package driven_adapter_restful

type HTTPErrorHandlingStrategyControllerOption interface {
	ClientOption
	Controller() func(lastExchange Exchange, exchangeError error) error
}

type httpErrorHandlingStrategyControllerOption struct {
	controller func(lastExchange Exchange, exchangeError error) error
}

func NewHTTPErrorHandlingStrategyControllerOption(
	controller func(lastExchange Exchange, exchangeError error) error,
) HTTPErrorHandlingStrategyControllerOption {
	return &httpErrorHandlingStrategyControllerOption{
		controller: controller,
	}
}

func (o *httpErrorHandlingStrategyControllerOption) IsRestfulClientOption() bool {
	return true
}

func (o *httpErrorHandlingStrategyControllerOption) Controller() func(lastExchange Exchange, exchangeError error) error {
	return o.controller
}
