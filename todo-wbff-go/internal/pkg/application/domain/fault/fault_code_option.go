package domain_fault

type faultCodeOption struct {
	code string
}

func (o *faultCodeOption) IsFaultOption() bool {
	return true
}

func (o *faultCodeOption) Code() string {
	return o.code
}
