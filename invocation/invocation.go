package invocation

import "encoding/json"

//https://docs.aws.amazon.com/appsync/latest/devguide/resolver-mapping-template-reference-lambda.html
type Invocation struct {
	Field   string  `json:"filed"`
	Context Context `json:"context"`
}

// https://docs.aws.amazon.com/appsync/latest/devguide/resolver-context-reference.html
// https://docs.aws.amazon.com/appsync/latest/devguide/tutorial-lambda-resolvers.html
type Context struct {
	Arguments *json.RawMessage `json:"arguments"` // AWS Doc: A map containing all GraphQL arguments for this field
	Source    *json.RawMessage `json:"source"`    // AWS Doc: A map containing the resolution of the parent field.
	Result    *json.RawMessage `json:"result"`
	Identity  *json.RawMessage `json:"identity"`
	Request   *json.RawMessage `json:"request"`
}

func (invocation Invocation) Payload() (*json.RawMessage, error) {
	if invocation.Context.Arguments != nil {
		return invocation.Context.Arguments, nil
	} else if invocation.Context.Source != nil {
		return invocation.Context.Source, nil
	}

	return nil, EmptyPayloadError(invocation.Context.Arguments, invocation.Context.Source)
}
