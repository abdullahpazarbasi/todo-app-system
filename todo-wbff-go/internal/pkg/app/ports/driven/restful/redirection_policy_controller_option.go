package driven_app_ports_restful

type RedirectionPolicyControllerOption interface {
	Controller() func(statusCode int, targetURL string, header map[string][]string) (redirectability bool)
}
