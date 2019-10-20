package invocation

//https://docs.aws.amazon.com/appsync/latest/devguide/resolver-mapping-template-reference-lambda.html
type Invocation struct {
	Field   string  `json:"filed"`
	Context Context `json:"context"`
}

// https://docs.aws.amazon.com/appsync/latest/devguide/resolver-context-reference.html
// https://docs.aws.amazon.com/appsync/latest/devguide/tutorial-lambda-resolvers.html
type Context struct {
	Arguments *[]byte `json:"arguments"` // AWS Doc: A map containing all GraphQL arguments for this field
	Source    *[]byte `json:"source"`    // AWS Doc: A map containing the resolution of the parent field.
	Result    *[]byte `json:"result"`
	Identity  *[]byte `json:"identity"`
	Request   *[]byte `json:"request"`
}

func (invocation Invocation) Payload() (*[]byte, error) {
	if invocation.Context.Arguments != nil {
		return invocation.Context.Arguments, nil
	} else if invocation.Context.Source != nil {
		return invocation.Context.Source, nil
	}

	return nil, EmptyPayloadError(invocation.Context.Arguments, invocation.Context.Source)
}
