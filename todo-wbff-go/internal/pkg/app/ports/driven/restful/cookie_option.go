package driven_app_ports_restful

type CookieOption interface {
	ClientOption
	Cookie() Cookie
}
