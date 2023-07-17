package driven_app_ports_restful

type AuthorizationSchemeAndTokenOption interface {
	ClientOption
	Scheme() string
	Token() string
}
