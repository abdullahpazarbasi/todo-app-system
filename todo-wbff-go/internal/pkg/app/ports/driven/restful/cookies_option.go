package driven_app_ports_restful

type CookiesOption interface {
	ClientOption
	Cookies() *[]Cookie
}
