package driven_adapter_restful

type ResourcePathParametersOption interface {
	ClientOption
	Map() map[string]string
}
