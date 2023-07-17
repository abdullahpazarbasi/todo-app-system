package driven_app_ports_restful

type ResourcePathParameterOption interface {
	ClientOption
	Placeholder() string
	Value() string
}
