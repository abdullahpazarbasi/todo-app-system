package driven_app_ports_restful

import "time"

type Exchange interface {
	Request() Request
	Response() Response
	ElapsedTime() time.Duration
	Cookies() *[]Cookie
	Previous() Exchange
}
