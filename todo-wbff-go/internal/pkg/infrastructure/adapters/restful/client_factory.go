package infrastructure_adapters_restful

import (
	drivenAppPortsRestful "todo-app-wbff/internal/pkg/app/ports/driven/restful"
)

type clientFactory struct{}

func NewClientFactory() drivenAppPortsRestful.ClientFactory {
	return &clientFactory{}
}

func (cf *clientFactory) Create(serverBaseURL string) drivenAppPortsRestful.Client {
	return &client{
		serverBaseURL: serverBaseURL,
	}
}
