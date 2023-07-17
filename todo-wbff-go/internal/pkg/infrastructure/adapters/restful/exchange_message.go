package infrastructure_adapters_restful

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"
	drivenAppPortsRestful "todo-app-wbff/internal/pkg/app/ports/driven/restful"
)

func exchangeMessage(
	ctx context.Context,
	method string,
	resourceURL string,
	body io.Reader,
	header *map[string][]string,
	cookies *[]drivenAppPortsRestful.Cookie,
	timeOutLimit time.Duration,
	redirectionPolicyController func(
		statusCode int,
		targetURL string,
		header map[string][]string,
	) (
		redirectability bool,
	),
	previousExchange drivenAppPortsRestful.Exchange,
) (
	drivenAppPortsRestful.Exchange,
	error,
) {
	var err error
	ec := exchange{
		previous: previousExchange,
	}
	rq := request{}
	ec.request = &rq
	rq.raw, err = http.NewRequestWithContext(ctx, method, resourceURL, body)
	if err != nil {
		return nil, err
	}
	rq.raw.Header = *header
	host := rq.raw.Header.Get("Host")
	if host != "" {
		rq.raw.Host = host
	}
	cookieJar := createCookieJar()
	cookieJar.SetCookies(rq.raw.URL, convertCookiesToHttpCookies(cookies))
	c := http.Client{
		Transport: createTransport(nil),
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if redirectionPolicyController != nil {
				res := req.Response
				var targetURL *url.URL
				targetURL, err = res.Location()
				if err != nil {
					return http.ErrUseLastResponse
				}
				if redirectionPolicyController(res.StatusCode, targetURL.String(), res.Header) {
					return nil
				}
			}

			return http.ErrUseLastResponse
		},
		Jar:     cookieJar,
		Timeout: timeOutLimit,
	}
	rs := response{}
	ec.response = &rs
	rq.sentAt = time.Now()
	rs.raw, err = c.Do(rq.raw)
	defer silentlyClose(rs.raw.Body)
	rs.receivedAt = time.Now()
	ec.cookies = convertHttpCookiesToCookies(rs.raw.Cookies())
	ec.elapsedTime = rs.ReceivedAt().Sub(*rq.SentAt())
	if err != nil {
		return nil, err
	}

	return &ec, nil
}
