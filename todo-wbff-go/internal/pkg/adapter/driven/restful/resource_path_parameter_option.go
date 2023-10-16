package driven_adapter_restful

type ResourcePathParameterOption interface {
	ClientOption
	Placeholder() string
	Value() string
}

type resourcePathParameterOption struct {
	placeholder string
	value       string
}

func NewResourcePathParameterOption(placeholder, value string) ResourcePathParameterOption {
	return &resourcePathParameterOption{
		placeholder: placeholder,
		value:       value,
	}
}

func (o *resourcePathParameterOption) IsRestfulClientOption() bool {
	return true
}

func (o *resourcePathParameterOption) Placeholder() string {
	return o.placeholder
}

func (o *resourcePathParameterOption) Value() string {
	return o.value
}
