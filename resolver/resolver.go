package resolver

import (
	"context"
	"encoding/json"
)

type Request *json.RawMessage
type Response interface{}
type Handler interface{}

type Resolver struct {
	// TODO: Enforce the signature of an handler: (context, payload) -> (response, error)
	Handler Handler
}

func ToResolver(resolver interface{}) (Resolver, error) {
	// TODO: Validation, is it a resolver?
	return Resolver{}, nil
}

func (resolver Resolver) Resolve(context context.Context, request Request) (Response, error) {
	// TODO: Map the argument and perform the call
	//reflect.ValueOf(resolver.handler).Call(payload)
	var response Response
	return response, nil
}
