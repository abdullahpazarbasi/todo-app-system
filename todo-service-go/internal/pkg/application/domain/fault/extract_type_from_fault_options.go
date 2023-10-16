package domain_fault

import domainFaultPort "todo-app-service/internal/pkg/application/domain/fault/port"

func extractTypeFromFaultOptions(
	options *[]domainFaultPort.FaultOption,
) (
	domainFaultPort.FaultType,
	*[]domainFaultPort.FaultOption,
) {
	var existent bool
	var tipe domainFaultPort.FaultType
	remainingOptions := make([]domainFaultPort.FaultOption, 0)
	for _, option := range *options {
		switch o := option.(type) {
		case domainFaultPort.FaultTypeOption:
			tipe = o.Type()
			existent = true
		default:
			remainingOptions = append(remainingOptions, option)
		}
	}
	if !existent {
		tipe = ""
	}

	return tipe, &remainingOptions
}
