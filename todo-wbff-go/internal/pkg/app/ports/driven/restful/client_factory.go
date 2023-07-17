package driven_app_ports_restful

type ClientFactory interface {
	Create(serverBaseURL string) Client
}
