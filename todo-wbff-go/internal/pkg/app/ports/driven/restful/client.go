package driven_app_ports_restful

import "context"

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
