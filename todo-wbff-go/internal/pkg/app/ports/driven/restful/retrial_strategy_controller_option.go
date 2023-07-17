package driven_app_ports_restful

type RetrialStrategyControllerOption interface {
	ClientOption
	Controller() func(lastExchange Exchange) (breaking bool)
}
