package domain_fault

import domainFaultPort "todo-app-wbff/internal/pkg/application/domain/fault/port"

func extractMessageFromFaultOptions(
	options *[]domainFaultPort.FaultOption,
	defaultValue string,
) (
	string,
	*[]domainFaultPort.FaultOption,
) {
	var existent bool
	var message string
	remainingOptions := make([]domainFaultPort.FaultOption, 0)
	for _, option := range *options {
		switch o := option.(type) {
		case domainFaultPort.FaultMessageOption:
			message = o.Message()
			existent = true
		default:
			remainingOptions = append(remainingOptions, option)
		}
	}
	if !existent {
		message = defaultValue
	}

	return message, &remainingOptions
}
