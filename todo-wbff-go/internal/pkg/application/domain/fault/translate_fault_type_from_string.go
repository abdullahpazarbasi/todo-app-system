package domain_fault

import domainFaultPort "todo-app-wbff/internal/pkg/application/domain/fault/port"

var availableFaultTypes = []domainFaultPort.FaultType{
	domainFaultPort.FaultTypeUnknown,
	domainFaultPort.FaultTypeConnectionFailure,
}

func translateFaultTypeFromString(faultTypeCandidate string) domainFaultPort.FaultType {
	for _, currentFaultType := range availableFaultTypes {
		if currentFaultType == domainFaultPort.FaultType(faultTypeCandidate) {
			return currentFaultType
		}
	}

	panic("unknown fault type")
}
