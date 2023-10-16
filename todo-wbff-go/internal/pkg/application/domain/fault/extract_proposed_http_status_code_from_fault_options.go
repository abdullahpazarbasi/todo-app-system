package domain_fault

import domainFaultPort "todo-app-wbff/internal/pkg/application/domain/fault/port"

func extractProposedHTTPStatusCodeFromFaultOptions(
	options *[]domainFaultPort.FaultOption,
	defaultValue int,
) (
	int,
	*[]domainFaultPort.FaultOption,
) {
	var existent bool
	var proposedHTTPStatusCode int
	remainingOptions := make([]domainFaultPort.FaultOption, 0)
	for _, option := range *options {
		switch o := option.(type) {
		case domainFaultPort.FaultProposedHTTPStatusCodeOption:
			proposedHTTPStatusCode = o.ProposedHTTPStatusCode()
			existent = true
		default:
			remainingOptions = append(remainingOptions, option)
		}
	}
	if !existent {
		proposedHTTPStatusCode = defaultValue
	}

	return proposedHTTPStatusCode, &remainingOptions
}
