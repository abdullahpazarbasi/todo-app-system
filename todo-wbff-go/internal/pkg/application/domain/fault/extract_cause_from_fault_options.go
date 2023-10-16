package domain_fault

import domainFaultPort "todo-app-wbff/internal/pkg/application/domain/fault/port"

func extractCauseFromFaultOptions(
	options *[]domainFaultPort.FaultOption,
) (
	error,
	*[]domainFaultPort.FaultOption,
) {
	var existent bool
	var cause error
	remainingOptions := make([]domainFaultPort.FaultOption, 0)
	for _, option := range *options {
		switch o := option.(type) {
		case domainFaultPort.FaultCauseOption:
			cause = o.Cause()
			existent = true
		default:
			remainingOptions = append(remainingOptions, option)
		}
	}
	if !existent {
		cause = nil
	}

	return cause, &remainingOptions
}
