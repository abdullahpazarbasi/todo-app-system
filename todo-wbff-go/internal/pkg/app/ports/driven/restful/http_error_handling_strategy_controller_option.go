package driven_app_ports_restful

type HTTPErrorHandlingStrategyControllerOption interface {
	ClientOption
	Controller() func(lastExchange Exchange, cause error) Error
}
