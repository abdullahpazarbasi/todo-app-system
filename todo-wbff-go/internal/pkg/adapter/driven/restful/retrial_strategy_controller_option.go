package driven_adapter_restful

type RetrialStrategyControllerOption interface {
	ClientOption
	Controller() func(lastExchange Exchange) (breaking bool)
}
