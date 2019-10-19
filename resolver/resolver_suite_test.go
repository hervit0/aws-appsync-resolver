package resolver_test

import (
	"context"
	"github.com/hervit0/aws-appsync-resolver/fixture"
	"github.com/hervit0/aws-appsync-resolver/resolver"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestResolver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Resolver Suite")
}

var _ = Describe("Resolver", func() {
	var r resolver.Resolver
	handler := func(ctx context.Context, request interface{}) (resolver.Response, error) {
		return nil, nil
	}

	BeforeEach(func() {
		r.Handler = handler
	})

	Describe(".Resolve", func() {
		It("call the resolver handler and return response", func() {
			request := fixture.RawRequestFixture()
			response, err := r.Resolve(context.TODO(), &request)

			Expect(response).To(BeNil())
			Expect(err).To(BeNil())
		})
	})
})
