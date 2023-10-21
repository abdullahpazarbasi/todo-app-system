package domain_fault

import domainFaultPort "todo-app-service/internal/pkg/application/domain/fault/port"

var availableFaultTypes = []domainFaultPort.FaultType{
	domainFaultPort.FaultTypeUnknown,
	domainFaultPort.FaultTypeConnectionFailure,
	domainFaultPort.FaultTypeInaccessibleServer,
	domainFaultPort.FaultTypeDatabaseNotPrivileged,
	domainFaultPort.FaultTypeCollectionNotFound,
	domainFaultPort.FaultTypeItemNotFound,
	domainFaultPort.FaultTypeDuplicatedEntry,
	domainFaultPort.FaultTypeAssociationViolation,
	domainFaultPort.FaultTypeTimeout,
	domainFaultPort.FaultTypeRaceCondition,
	domainFaultPort.FaultTypeStructuralIncompatibility,
}

func translateFaultTypeFromString(faultTypeCandidate string) domainFaultPort.FaultType {
	for _, currentFaultType := range availableFaultTypes {
		if currentFaultType == domainFaultPort.FaultType(faultTypeCandidate) {
			return currentFaultType
		}
	}

	return domainFaultPort.FaultTypeUnknown
}
