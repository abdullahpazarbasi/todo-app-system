package domain_fault

type faultCauseOption struct {
	cause error
}

func (o *faultCauseOption) IsFaultOption() bool {
	return true
}

func (o *faultCauseOption) Cause() error {
	return o.cause
}
