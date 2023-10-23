package driven_adapter_restful

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"
)

func exchangeMessage(
	ctx context.Context,
	method string,
	resourceURL string,
	body io.Reader,
	header map[string][]string,
	cookies *[]Cookie,
	timeOutLimit time.Duration,
	dialerTimeoutLimit time.Duration,
	keepAliveLifetime time.Duration,
	idleConnectionTimeoutLimit time.Duration,
	tlsHandshakeTimeoutLimit time.Duration,
	expectContinueTimeoutLimit time.Duration,
	maximumNumberOfIdleConnections int,
	redirectionPolicyController func(
		statusCode int,
		targetURL string,
		header map[string][]string,
	) (
		redirectability bool,
	),
	previousExchange Exchange,
) (
	Exchange,
	error,
) {
	var err error
	ec := exchange{
		cookies:  &[]Cookie{},
		previous: previousExchange,
	}
	rq := request{}
	ec.request = &rq
	rs := response{}
	ec.response = &rs
	rq.raw, err = http.NewRequestWithContext(ctx, method, resourceURL, body)
	if err != nil {
		return &ec, err
	}
	rq.raw.Header = header
	host := rq.raw.Header.Get("Host")
	if host != "" {
		rq.raw.Host = host
	}
	cookieJar := createCookieJar()
	cookieJar.SetCookies(rq.raw.URL, convertCookiesToHttpCookies(cookies))
	c := http.Client{
		Transport: createTransport(
			nil,
			dialerTimeoutLimit,
			keepAliveLifetime,
			idleConnectionTimeoutLimit,
			tlsHandshakeTimeoutLimit,
			expectContinueTimeoutLimit,
			maximumNumberOfIdleConnections,
		),
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
	rq.sentAt = time.Now()
	rs.raw, err = c.Do(rq.raw)
	rs.receivedAt = time.Now()
	if rs.raw != nil {
		//defer silentlyClose(rs.raw.Body)
		ec.cookies = convertHttpCookiesToCookies(rs.raw.Cookies())
	}
	ec.elapsedTime = rs.ReceivedAt().Sub(*rq.SentAt())
	if err != nil {
		return &ec, err
	}

	return &ec, nil
}
