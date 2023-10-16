package domain_fault

import domainFaultPort "todo-app-service/internal/pkg/application/domain/fault/port"

var availableFaultTypes = []domainFaultPort.FaultType{
	domainFaultPort.FaultTypeUnknown,
	domainFaultPort.FaultTypeConnectionFailure,
	domainFaultPort.FaultTypeInaccessibleServer,
	domainFaultPort.FaultTypeDatabaseNotPrivileged,
	domainFaultPort.FaultTypeCollectionNotFound,
	domainFaultPort.FaultTypeDuplicatedEntry,
	domainFaultPort.FaultTypeAssociationViolation,
}

func translateFaultTypeFromString(faultTypeCandidate string) domainFaultPort.FaultType {
	for _, currentFaultType := range availableFaultTypes {
		if currentFaultType == domainFaultPort.FaultType(faultTypeCandidate) {
			return currentFaultType
		}
	}

	panic("unknown fault type")
}
