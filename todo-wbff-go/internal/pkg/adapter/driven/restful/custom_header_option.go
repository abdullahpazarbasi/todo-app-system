package driven_adapter_restful

type CustomHeaderOption interface {
	ClientOption
	Map() map[string][]string
}
