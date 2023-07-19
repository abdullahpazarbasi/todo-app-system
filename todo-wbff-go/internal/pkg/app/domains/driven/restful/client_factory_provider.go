package driven_app_domains_restful

import drivenAppPortsRestful "todo-app-wbff/internal/pkg/app/ports/driven/restful"

type clientFactoryProvider struct {
	clientFactory drivenAppPortsRestful.ClientFactory
}

func NewClientFactoryProvider(clientFactory drivenAppPortsRestful.ClientFactory) *clientFactoryProvider {
	return &clientFactoryProvider{clientFactory: clientFactory}
}

func (cfp *clientFactoryProvider) Provide() drivenAppPortsRestful.ClientFactory {
	return cfp.clientFactory
}
