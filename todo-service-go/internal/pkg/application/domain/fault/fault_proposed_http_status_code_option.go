package domain_fault

type faultProposedHTTPStatusCodeOption struct {
	proposedHTTPStatusCode int
}

func (o *faultProposedHTTPStatusCodeOption) IsFaultOption() bool {
	return true
}

func (o *faultProposedHTTPStatusCodeOption) ProposedHTTPStatusCode() int {
	return o.proposedHTTPStatusCode
}
