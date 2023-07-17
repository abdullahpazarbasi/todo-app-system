package infrastructure_adapters_restful

import (
	"time"
	drivenAppPortsRestful "todo-app-wbff/internal/pkg/app/ports/driven/restful"
)

type exchange struct {
	request     drivenAppPortsRestful.Request
	response    drivenAppPortsRestful.Response
	cookies     *[]drivenAppPortsRestful.Cookie
	elapsedTime time.Duration
	previous    drivenAppPortsRestful.Exchange
}

func (ec *exchange) Request() drivenAppPortsRestful.Request {
	return ec.request
}

func (ec *exchange) Response() drivenAppPortsRestful.Response {
	return ec.response
}

func (ec *exchange) Cookies() *[]drivenAppPortsRestful.Cookie {
	return ec.cookies
}

func (ec *exchange) ElapsedTime() time.Duration {
	return ec.elapsedTime
}

func (ec *exchange) Previous() drivenAppPortsRestful.Exchange {
	return ec.previous
}
