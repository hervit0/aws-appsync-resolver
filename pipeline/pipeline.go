package pipeline

import (
	"context"
	"fmt"
	"github.com/hervit0/aws-appsync-resolver/invocation"
	"github.com/hervit0/aws-appsync-resolver/resolver"
)

type Pipeline map[string]resolver.Resolver

func (pipeline *Pipeline) Declare(resolverName string, handler resolver.Handler) *Pipeline {
	(*pipeline)[resolverName] = resolver.Resolver{Handler: handler}
	return pipeline
}

func (pipeline Pipeline) Handle(context context.Context, invocation invocation.Invocation) (resolver.Response, error) {
	resolverHandler, ok := pipeline[invocation.Field]
	if !ok {
		return nil, fmt.Errorf("resolver not found: %v", invocation.Field)
	}

	payload, errParsingPayload := invocation.Payload()
	if errParsingPayload != nil {
		return nil, errParsingPayload
	}

	return resolverHandler.Resolve(context, payload)
}
