package driven_adapter_restful

import (
	"time"
)

type Exchange interface {
	Request() Request
	Response() Response
	ElapsedTime() time.Duration
	Cookies() *[]Cookie
	Previous() Exchange
}

type exchange struct {
	request     Request
	response    Response
	cookies     *[]Cookie
	elapsedTime time.Duration
	previous    Exchange
}

func (ec *exchange) Request() Request {
	return ec.request
}

func (ec *exchange) Response() Response {
	return ec.response
}

func (ec *exchange) Cookies() *[]Cookie {
	return ec.cookies
}

func (ec *exchange) ElapsedTime() time.Duration {
	return ec.elapsedTime
}

func (ec *exchange) Previous() Exchange {
	return ec.previous
}
