package resolver

import (
	"context"
	"encoding/json"
)

type Request *json.RawMessage

type Response interface{}

type Handler func(context context.Context, request Request) (Response, error)

type Resolver struct {
	Handler Handler
}

func ToResolver(handler Handler) (Resolver, error) {
	// TODO: Validation, is it a resolver?
	// Too much indirection here, it's strongly type, basta

	return Resolver{
		Handler: handler,
	}, nil
}

func (resolver Resolver) Resolve(context context.Context, request Request) (Response, error) {
	// TODO: Map the argument and perform the call
	//reflect.ValueOf(resolver.handler).Call(payload)
	var response Response
	return response, nil
}
