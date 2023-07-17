package driven_app_ports_restful

import "time"

type Cookie interface {
	Name() string
	Value() string

	Path() string       // optional
	Domain() string     // optional
	Expiry() *time.Time // optional

	// MaxAge = 0 means no 'Max-Age' attribute specified. < 0 means delete cookie now, equivalently 'Max-Age: 0'. > 0 means Max-Age attribute present and given in seconds
	MaxAge() int
	Secure() bool
	HttpOnly() bool
	SameSite() int
}
