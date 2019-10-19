package resolver

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
)

type Request *json.RawMessage

type Response interface{}

type Handler func(context context.Context, request Request) (Response, error)

type Resolver struct {
	Handler Handler
}

func (resolver Resolver) Resolve(context context.Context, request Request) (Response, error) {
	handlerRequestType := reflect.TypeOf(resolver.Handler).In(1)
	requestArguments, err := parse(request, &handlerRequestType)
	if err != nil {
		return nil, err
	}
	requestArgumentsWithContext := append([]reflect.Value{reflect.ValueOf(context)}, requestArguments...)

	response := reflect.ValueOf(resolver.Handler).Call(requestArgumentsWithContext)
	return response[0].Interface(), response[1].Interface().(error)
}

func parse(rawPayload Request, requestType *reflect.Type) ([]reflect.Value, error) {
	args := reflect.New(*requestType)

	if err := json.Unmarshal(*rawPayload, args.Interface()); err != nil {
		return nil, fmt.Errorf("error parsing payload: %v", err)
	}

	return append([]reflect.Value{}, args.Elem()), nil
}
