package resolver

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response interface{}

type Resolver struct {
	Handler lambda.Handler
}

func ToResolver(handler interface{}) (Resolver, error) {
	resolverHandler := lambda.NewHandler(handler)
	return Resolver{Handler: resolverHandler}, nil
}

func (resolver Resolver) Resolve(context context.Context, request []byte) (Response, error) {
	return resolver.Handler.Invoke(context, request)
}
