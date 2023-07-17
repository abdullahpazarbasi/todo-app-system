package driven_app_ports_restful

type ResourcePathParametersOption interface {
	ClientOption
	Map() map[string]string
}
