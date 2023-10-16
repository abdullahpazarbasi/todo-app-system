package driven_adapter_restful

type BasicAuthenticationCredentialsOption interface {
	ClientOption
	Username() string
	Password() string
}
