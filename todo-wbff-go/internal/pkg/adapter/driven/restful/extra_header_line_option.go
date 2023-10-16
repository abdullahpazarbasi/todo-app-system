package driven_adapter_restful

type ExtraHeaderLineOption interface {
	ClientOption
	Name() string
	Value() string
}
