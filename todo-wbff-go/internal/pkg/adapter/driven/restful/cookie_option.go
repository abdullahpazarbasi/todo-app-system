package driven_adapter_restful

type CookieOption interface {
	ClientOption
	Cookie() Cookie
}
