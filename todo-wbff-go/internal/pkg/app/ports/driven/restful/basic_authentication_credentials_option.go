package driven_app_ports_restful

type BasicAuthenticationCredentialsOption interface {
	ClientOption
	Username() string
	Password() string
}
