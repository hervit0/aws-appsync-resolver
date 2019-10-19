package pipeline_test

import (
	"context"
	"github.com/hervit0/aws-appsync-resolver/pipeline"
	"github.com/hervit0/aws-appsync-resolver/resolver"
	"reflect"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPipeline(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pipeline Suite")
}

var _ = Describe("Pipeline", func() {
	var pipeline pipeline.Pipeline
	handler := func(ctx context.Context, request resolver.Request) (resolver.Response, error) {
		return nil, nil
	}

	BeforeEach(func() {
		pipeline = make(map[string]resolver.Resolver)
	})

	Describe(".Declare", func() {
		It("add a unique resolver", func() {
			pipeline.Declare("handler", handler)

			storedResolver, ok := pipeline["handler"]
			Expect(ok).To(BeTrue())
			Expect(reflect.ValueOf(storedResolver.Handler).Pointer()).To(Equal(reflect.ValueOf(handler).Pointer()))
		})

		It("add multiple resolvers", func() {
			pipeline.Declare("handler A", handler).Declare("handler B", handler)

			storedResolverA, ok := pipeline["handler A"]
			Expect(ok).To(BeTrue())
			Expect(reflect.ValueOf(storedResolverA.Handler).Pointer()).To(Equal(reflect.ValueOf(handler).Pointer()))

			storedResolverB, ok := pipeline["handler B"]
			Expect(ok).To(BeTrue())
			Expect(reflect.ValueOf(storedResolverB.Handler).Pointer()).To(Equal(reflect.ValueOf(handler).Pointer()))
		})
	})
})
