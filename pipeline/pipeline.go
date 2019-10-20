package pipeline

import (
	"context"
	"fmt"
	"github.com/hervit0/aws-appsync-resolver/invocation"
	"github.com/hervit0/aws-appsync-resolver/resolver"
	"log"
)

type Pipeline map[string]resolver.Resolver

func (pipeline *Pipeline) Declare(resolverName string, handler interface{}) *Pipeline {
	resolverOp, err := resolver.ToResolver(handler)
	if err != nil {
		log.Printf("error resolver declaration: %v", err)
	}
	(*pipeline)[resolverName] = resolverOp
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

	return resolverHandler.Resolve(context, *payload)
}
