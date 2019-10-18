package pipeline

import (
	"context"
	"fmt"
	"github.com/hervit0/aws-appsync-resolver/invocation"
	"github.com/hervit0/aws-appsync-resolver/resolver"
)

type Pipeline map[string]resolver.Resolver

// TODO: Here, interface{} represents the handler method of the resolver
func (pipeline Pipeline) Declare(resolverName string, handler interface{}) error {
	resolverHandler, errorCreateResolver := resolver.ToResolver(handler)
	if errorCreateResolver != nil {
		pipeline[resolverName] = resolverHandler
	}
	return errorCreateResolver
}

// TODO: Here, interface{} represents the return type of the handler response
func (pipeline Pipeline) Handle(context context.Context, invocation invocation.Invocation) (interface{}, error) {
	resolverHandler, ok := pipeline[invocation.Field]
	if !ok {
		return nil, fmt.Errorf("resolver not found: %v", invocation.Field)
	}

	payload, errParsingPayload := invocation.Payload()
	if errParsingPayload != nil {
		return nil, errParsingPayload
	}

	return resolverHandler.Resolve(context, payload), nil
}
