package driven_adapter_restful

import "net/url"

type formData map[string][]string

func (fd *formData) Export() map[string][]string {
	return *fd
}

func (fd *formData) Encode() string {
	return url.Values(*fd).Encode()
}
