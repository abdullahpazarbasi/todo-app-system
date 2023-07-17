package driven_app_ports_restful

type ClientFactoryProvider interface {
	Provide() ClientFactory
}
