package pipeline_test

import (
	"context"
	"errors"
	"github.com/hervit0/aws-appsync-resolver/fixture"
	"github.com/hervit0/aws-appsync-resolver/invocation"
	"github.com/hervit0/aws-appsync-resolver/pipeline"
	"github.com/hervit0/aws-appsync-resolver/resolver"
	"log"
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

	BeforeEach(func() {
		pipeline = make(map[string]resolver.Resolver)
	})

	//Describe(".Declare", func() {
	//	It("add a unique resolver", func() {
	//		pipeline.Declare("handler", fixture.Handler{Error: "yo"}.Invoke)
	//
	//		storedResolver, ok := pipeline["handler"]
	//		Expect(ok).To(BeTrue())
	//
	//		_, err := storedResolver.Handler.Invoke(context.TODO(), fixture.RawRequestFixture())
	//		Expect(err).To(Equal(errors.New("yo")))
	//	})
	//
	//	It("add multiple resolvers", func() {
	//		pipeline.Declare("handler A", fixture.Handler{Error: "A"}.Invoke)
	//		pipeline.Declare("handler B", fixture.Handler{Error: "B"}.Invoke)
	//
	//		_, ok := pipeline["handler A"]
	//		Expect(ok).To(BeTrue())
	//
	//		_, ok = pipeline["handler B"]
	//		Expect(ok).To(BeTrue())
	//	})
	//
	//	It("map correctly the handlers", func() {
	//		pipeline.Declare("handler A", fixture.Handler{Error: "A"}.Invoke)
	//		pipeline.Declare("handler B", fixture.Handler{Error: "B"}.Invoke)
	//
	//		storedResolverA, _ := pipeline["handler A"]
	//		_, err := storedResolverA.Handler.Invoke(context.TODO(), fixture.RawRequestFixture())
	//		Expect(err).To(Equal(errors.New("A")))
	//
	//		storedResolverB, _ := pipeline["handler B"]
	//		_, err = storedResolverB.Handler.Invoke(context.TODO(), fixture.RawRequestFixture())
	//		Expect(err).To(Equal(errors.New("B")))
	//	})
	//})

	Describe(".Resolve", func() {
		It("call a resolver", func() {
			pipeline.Declare("handler", fixture.Handler{Error: "yo"}.Handle)
			request := fixture.RawRequestFixture()
			invocationRequest := invocation.Invocation{
				Field: "handler",
				Context: invocation.Context{
					Arguments: &request,
				},
			}

			response, err := pipeline.Handle(context.TODO(), invocationRequest)

			log.Printf("fixture: %v", fixture.RawResponseFixture())
			log.Printf("response: %v", response)
			//Expect(response).To(Equal(fixture.RawResponseFixture()))
			Expect(err).To(Equal(errors.New("yo")))
		})
	})
})
