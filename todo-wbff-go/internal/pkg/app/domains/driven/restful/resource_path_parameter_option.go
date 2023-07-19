package driven_app_domains_restful

type resourcePathParameterOption struct {
	placeholder string
	value       string
}

func NewResourcePathParameterOption(placeholder, value string) *resourcePathParameterOption {
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
