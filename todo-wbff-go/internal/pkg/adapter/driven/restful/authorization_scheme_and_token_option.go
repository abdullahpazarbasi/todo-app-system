package driven_adapter_restful

type AuthorizationSchemeAndTokenOption interface {
	ClientOption
	Scheme() string
	Token() string
}
