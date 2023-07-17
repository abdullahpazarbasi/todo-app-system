package driven_app_domains_restful

import "net/url"

type queryParameters map[string][]string

func (qp *queryParameters) Export() map[string][]string {
	return *qp
}

func (qp *queryParameters) Encode() string {
	return url.Values(*qp).Encode()
}
