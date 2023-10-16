package driven_adapter_restful

type CookiesOption interface {
	ClientOption
	Cookies() *[]Cookie
}
