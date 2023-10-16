package domain_fault

import domainFaultPort "todo-app-wbff/internal/pkg/application/domain/fault/port"

type faultTypeOption struct {
	tipe domainFaultPort.FaultType
}

func (o *faultTypeOption) IsFaultOption() bool {
	return true
}

func (o *faultTypeOption) Type() domainFaultPort.FaultType {
	return o.tipe
}
