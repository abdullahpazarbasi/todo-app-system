package driven_adapter_restful

type ExtraHeaderLineOption interface {
	ClientOption
	Name() string
	Value() string
}

type extraHeaderLineOption struct {
	name  string
	value string
}

func NewExtraHeaderLineOption(name string, value string) ExtraHeaderLineOption {
	return &extraHeaderLineOption{
		name:  name,
		value: value,
	}
}

func (o *extraHeaderLineOption) IsRestfulClientOption() bool {
	return true
}

func (o *extraHeaderLineOption) Name() string {
	return o.name
}

func (o *extraHeaderLineOption) Value() string {
	return o.value
}
