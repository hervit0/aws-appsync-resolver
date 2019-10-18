package resolver

type Resolver struct {
}

func ToResolver(resolver interface{}) (Resolver, error) {
	return Resolver{}, nil
}
