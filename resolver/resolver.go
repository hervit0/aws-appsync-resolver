package resolver

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

type Response interface{}

type Resolver struct {
	Handler lambda.Handler
}

func ToResolver(handler interface{}) (Resolver, error) {
	resolverHandler := lambda.NewHandler(handler)
	return Resolver{Handler: resolverHandler}, nil
}

func (r Resolver) Resolve(context context.Context, request []byte) (Response, error) {
	function := lambda.NewFunction(r.Handler)
	function.Invoke()
	l, e := r.Handler.Invoke(context, request)
	log.Printf("yo: %v", l)
	//return resolver.Handler.Invoke(context, request)
	return l, e
}
