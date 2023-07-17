package driven_app_ports_restful

type ParameterMap interface {
	Export() map[string][]string
}
