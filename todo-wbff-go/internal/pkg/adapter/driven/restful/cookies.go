package driven_adapter_restful

import (
	"net/http"
)

func convertCookiesToHttpCookies(cookies *[]Cookie) []*http.Cookie {
	httpCookies := make([]*http.Cookie, 0)
	if cookies == nil {
		return httpCookies
	}
	for _, naturalCookie := range *cookies {
		httpCookies = append(httpCookies, NewHttpCookieFromCookie(naturalCookie))
	}

	return httpCookies
}

func convertHttpCookiesToCookies(httpCookies []*http.Cookie) *[]Cookie {
	cookies := make([]Cookie, 0)
	for _, httpCookie := range httpCookies {
		cookies = append(cookies, NewCookieFromHttpCookie(httpCookie))
	}

	return &cookies
}
