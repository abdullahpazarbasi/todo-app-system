package driven_adapter_restful

import (
	"net"
	"net/http"
	"runtime"
	"time"
)

func createTransport(
	localAddr net.Addr,
	dialerTimeoutLimit time.Duration,
	keepAliveLifetime time.Duration,
	idleConnectionTimeoutLimit time.Duration,
	tlsHandshakeTimeoutLimit time.Duration,
	expectContinueTimeoutLimit time.Duration,
	maximumNumberOfIdleConnections int,
) *http.Transport {
	dialer := &net.Dialer{
		Timeout:   dialerTimeoutLimit,
		KeepAlive: keepAliveLifetime,
	}
	if localAddr != nil {
		dialer.LocalAddr = localAddr
	}

	return &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          maximumNumberOfIdleConnections,
		IdleConnTimeout:       idleConnectionTimeoutLimit,
		TLSHandshakeTimeout:   tlsHandshakeTimeoutLimit,
		ExpectContinueTimeout: expectContinueTimeoutLimit,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
	}
}
