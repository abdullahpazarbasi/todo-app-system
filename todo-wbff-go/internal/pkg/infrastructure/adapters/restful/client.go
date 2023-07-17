package infrastructure_adapters_restful

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strings"
	"time"
	drivenAppPortsRestful "todo-app-wbff/internal/pkg/app/ports/driven/restful"
)

type client struct {
	serverBaseURL string
	exchange      drivenAppPortsRestful.Exchange
	cookies       *[]drivenAppPortsRestful.Cookie
}

func (c *client) Post(
	ctx context.Context,
	resourcePathPattern string,
	stateRepresentation interface{},
	options ...drivenAppPortsRestful.ClientOption,
) (
	drivenAppPortsRestful.Exchange,
	error,
) {
	resty.New()
	return c.ExchangeMessage(
		ctx,
		"POST",
		resourcePathPattern,
		nil,
		stateRepresentation,
		&options,
	)
}

func (c *client) Get(
	ctx context.Context,
	resourcePathPattern string,
	queryParameters drivenAppPortsRestful.ParameterMap,
	options ...drivenAppPortsRestful.ClientOption,
) (
	drivenAppPortsRestful.Exchange,
	error,
) {
	return c.ExchangeMessage(
		ctx,
		"GET",
		resourcePathPattern,
		queryParameters,
		nil,
		&options,
	)
}

func (c *client) Put(
	ctx context.Context,
	resourcePathPattern string,
	stateRepresentation interface{},
	options ...drivenAppPortsRestful.ClientOption,
) (
	drivenAppPortsRestful.Exchange,
	error,
) {
	return c.ExchangeMessage(
		ctx,
		"PUT",
		resourcePathPattern,
		nil,
		stateRepresentation,
		&options,
	)
}

func (c *client) Patch(
	ctx context.Context,
	resourcePathPattern string,
	stateRepresentation interface{},
	options ...drivenAppPortsRestful.ClientOption,
) (
	drivenAppPortsRestful.Exchange,
	error,
) {
	return c.ExchangeMessage(
		ctx,
		"PATCH",
		resourcePathPattern,
		nil,
		stateRepresentation,
		&options,
	)
}

func (c *client) Delete(
	ctx context.Context,
	resourcePathPattern string,
	options ...drivenAppPortsRestful.ClientOption,
) (
	drivenAppPortsRestful.Exchange,
	error,
) {
	return c.ExchangeMessage(
		ctx,
		"DELETE",
		resourcePathPattern,
		nil,
		nil,
		&options,
	)
}

func (c *client) ExchangeMessage(
	ctx context.Context,
	method string,
	resourcePathPattern string,
	queryParameters drivenAppPortsRestful.ParameterMap,
	stateRepresentation interface{},
	options *[]drivenAppPortsRestful.ClientOption,
) (
	drivenAppPortsRestful.Exchange,
	error,
) {
	resourceURLPattern := fmt.Sprintf(
		"%s/%s",
		strings.TrimRight(c.serverBaseURL, "/"),
		strings.TrimLeft(resourcePathPattern, "/"),
	)
	var pathParameters *map[string]string
	pathParameters, options = extractPathParametersFromClientOptions(options)
	var resourceURL string
	resourceURL = applyPathParametersToURL(resourceURLPattern, pathParameters)
	resourceURL = applyQueryParametersToURL(resourceURL, queryParameters)
	var header *map[string][]string
	header, options = extractHeaderFromClientOptions(options)
	var additionalCookies *[]drivenAppPortsRestful.Cookie
	additionalCookies, options = extractCookiesFromClientOptions(options)
	c.appendCookies(additionalCookies)
	bodyReader := resolveBodyReader(stateRepresentation)
	var timeOutLimit time.Duration
	timeOutLimit, options = extractTimeOutLimitFromClientOptions(options)
	var redirectionPolicyController func(int, string, map[string][]string) bool
	redirectionPolicyController, options = extractRedirectionPolicyControllerFromClientOptions(options)
	var retrialStrategyController func(drivenAppPortsRestful.Exchange) bool
	retrialStrategyController, options = extractRetrialStrategyControllerFromClientOptions(options)
	var ec drivenAppPortsRestful.Exchange
	var err error
	for i := 0; i < 100; i++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			ec, err = exchangeMessage(
				ctx,
				method,
				resourceURL,
				bodyReader,
				header,
				c.cookies,
				timeOutLimit,
				redirectionPolicyController,
				ec,
			)
			if err != nil {
				if retrialStrategyController != nil {
					if retrialStrategyController(ec) {
						return ec, err
					}
				}
			}
			c.cookies = ec.Cookies()
		}
	}

	return ec, err
}

func (c *client) appendCookies(additionalCookies *[]drivenAppPortsRestful.Cookie) {
	for _, nc := range *additionalCookies {
		*c.cookies = append(*c.cookies, nc)
	}
}
