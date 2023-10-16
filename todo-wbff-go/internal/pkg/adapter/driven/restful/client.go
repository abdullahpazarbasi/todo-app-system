package driven_adapter_restful

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type Client interface {
	Post(
		ctx context.Context,
		resourcePathPattern string,
		stateRepresentation interface{},
		options ...ClientOption,
	) (
		Exchange,
		error,
	)
	Get(
		ctx context.Context,
		resourcePathPattern string,
		queryParameters ParameterMap,
		options ...ClientOption,
	) (
		Exchange,
		error,
	)
	Put(
		ctx context.Context,
		resourcePathPattern string,
		stateRepresentation interface{},
		options ...ClientOption,
	) (
		Exchange,
		error,
	)
	Patch(
		ctx context.Context,
		resourcePathPattern string,
		stateRepresentation interface{},
		options ...ClientOption,
	) (
		Exchange,
		error,
	)
	Delete(
		ctx context.Context,
		resourcePathPattern string,
		options ...ClientOption,
	) (
		Exchange,
		error,
	)
}

type client struct {
	serverBaseURL string
	exchange      Exchange
	cookies       *[]Cookie
}

func NewClient(serverBaseURL string) Client {
	return &client{
		serverBaseURL: serverBaseURL,
	}
}

func (c *client) Post(
	ctx context.Context,
	resourcePathPattern string,
	stateRepresentation interface{},
	options ...ClientOption,
) (
	Exchange,
	error,
) {
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
	queryParameters ParameterMap,
	options ...ClientOption,
) (
	Exchange,
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
	options ...ClientOption,
) (
	Exchange,
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
	options ...ClientOption,
) (
	Exchange,
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
	options ...ClientOption,
) (
	Exchange,
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
	queryParameters ParameterMap,
	stateRepresentation interface{},
	options *[]ClientOption,
) (
	Exchange,
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
	var additionalCookies *[]Cookie
	additionalCookies, options = extractCookiesFromClientOptions(options)
	c.appendCookies(additionalCookies)
	bodyReader := resolveBodyReader(stateRepresentation)
	var timeOutLimit time.Duration
	timeOutLimit, options = extractTimeOutLimitFromClientOptions(options)
	var redirectionPolicyController func(int, string, map[string][]string) bool
	redirectionPolicyController, options = extractRedirectionPolicyControllerFromClientOptions(options)
	var retrialStrategyController func(Exchange) bool
	retrialStrategyController, options = extractRetrialStrategyControllerFromClientOptions(options)
	var httpErrorHandlingStrategyController func(lastExchange Exchange, cause error) error
	httpErrorHandlingStrategyController, options = extractHTTPErrorHandlingStrategyControllerFromClientOptions(options)
	var xchng Exchange
	var err error
	var lastTrial bool
	for i := 0; i < 100; i++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			xchng, err = exchangeMessage(
				ctx,
				method,
				resourceURL,
				bodyReader,
				header,
				c.cookies,
				timeOutLimit,
				redirectionPolicyController,
				xchng,
			)
			if err != nil {
				lastTrial = retrialStrategyController(xchng)
			}
			c.cookies = xchng.Cookies()
		}
		if lastTrial {
			break
		}
	}

	httpError := httpErrorHandlingStrategyController(xchng, err)
	if httpError != nil {
		return xchng, httpError
	}

	return xchng, err
}

func (c *client) appendCookies(additionalCookies *[]Cookie) {
	for _, nc := range *additionalCookies {
		*c.cookies = append(*c.cookies, nc)
	}
}
