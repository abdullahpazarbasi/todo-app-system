package domain_fault

import domainFaultPort "todo-app-wbff/internal/pkg/application/domain/fault/port"

func extractCodeFromFaultOptions(
	options *[]domainFaultPort.FaultOption,
	defaultValue string,
) (
	string,
	*[]domainFaultPort.FaultOption,
) {
	var existent bool
	var code string
	remainingOptions := make([]domainFaultPort.FaultOption, 0)
	for _, option := range *options {
		switch o := option.(type) {
		case domainFaultPort.FaultCodeOption:
			code = o.Code()
			existent = true
		default:
			remainingOptions = append(remainingOptions, option)
		}
	}
	if !existent {
		code = defaultValue
	}

	return code, &remainingOptions
}
