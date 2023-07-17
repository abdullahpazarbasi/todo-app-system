package driven_app_ports_restful

type CustomHeaderOption interface {
	ClientOption
	Map() map[string][]string
}
