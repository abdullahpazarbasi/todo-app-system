package infrastructure_adapters_restful

import (
	"golang.org/x/net/publicsuffix"
	"net/http"
	"net/http/cookiejar"
)

func createCookieJar() http.CookieJar {
	j, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

	return j
}
