package driven_app_ports_restful

type ExtraHeaderLineOption interface {
	ClientOption
	Name() string
	Value() string
}
