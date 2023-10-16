package driven_adapter_restful

type ParameterMap interface {
	Export() map[string][]string
}
