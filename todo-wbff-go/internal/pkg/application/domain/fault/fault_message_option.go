package domain_fault

type faultMessageOption struct {
	message string
}

func (o *faultMessageOption) IsFaultOption() bool {
	return true
}

func (o *faultMessageOption) Message() string {
	return o.message
}
