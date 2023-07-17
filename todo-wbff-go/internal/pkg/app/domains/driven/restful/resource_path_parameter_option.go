package driven_app_domains_restful

import drivenAppPortsRestful "todo-app-wbff/internal/pkg/app/ports/driven/restful"

type resourcePathParameterOption struct {
	placeholder string
	value       string
}

func NewResourcePathParameterOption(placeholder, value string) drivenAppPortsRestful.ResourcePathParameterOption {
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
