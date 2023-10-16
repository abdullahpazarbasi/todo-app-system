package driven_adapter_restful

import (
	"net/http"
	"time"
)

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

type cookie struct {
	name     string
	value    string
	path     string
	domain   string
	expiry   time.Time
	maxAge   int
	secure   bool
	httpOnly bool
	sameSite int
}

func NewCookieFromHttpCookie(httpCookie *http.Cookie) Cookie {
	return &cookie{
		name:     httpCookie.Name,
		value:    httpCookie.Value,
		path:     httpCookie.Path,
		domain:   httpCookie.Domain,
		expiry:   httpCookie.Expires,
		maxAge:   httpCookie.MaxAge,
		secure:   httpCookie.Secure,
		httpOnly: httpCookie.HttpOnly,
		sameSite: int(httpCookie.SameSite),
	}
}

func NewHttpCookieFromCookie(c Cookie) *http.Cookie {
	return &http.Cookie{
		Name:     c.Name(),
		Value:    c.Value(),
		Path:     c.Path(),
		Domain:   c.Domain(),
		Expires:  *c.Expiry(),
		MaxAge:   c.MaxAge(),
		Secure:   c.Secure(),
		HttpOnly: c.HttpOnly(),
		SameSite: http.SameSite(c.SameSite()),
	}
}

func (c *cookie) Name() string {
	return c.name
}

func (c *cookie) Value() string {
	return c.value
}

func (c *cookie) Path() string {
	return c.path
}

func (c *cookie) Domain() string {
	return c.domain
}

func (c *cookie) Expiry() *time.Time {
	return &c.expiry
}

func (c *cookie) MaxAge() int {
	return c.maxAge
}

func (c *cookie) Secure() bool {
	return c.secure
}

func (c *cookie) HttpOnly() bool {
	return c.httpOnly
}

func (c *cookie) SameSite() int {
	return c.sameSite
}
