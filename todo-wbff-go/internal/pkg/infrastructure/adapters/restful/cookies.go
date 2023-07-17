package infrastructure_adapters_restful

import (
	"net/http"
	drivenAppPortsRestful "todo-app-wbff/internal/pkg/app/ports/driven/restful"
)

func convertCookiesToHttpCookies(cookies *[]drivenAppPortsRestful.Cookie) []*http.Cookie {
	httpCookies := make([]*http.Cookie, 0)
	for _, naturalCookie := range *cookies {
		httpCookies = append(httpCookies, NewHttpCookieFromCookie(naturalCookie))
	}

	return httpCookies
}

func convertHttpCookiesToCookies(httpCookies []*http.Cookie) *[]drivenAppPortsRestful.Cookie {
	cookies := make([]drivenAppPortsRestful.Cookie, 0)
	for _, httpCookie := range httpCookies {
		cookies = append(cookies, NewCookieFromHttpCookie(httpCookie))
	}

	return &cookies
}
