package invocation

//https://docs.aws.amazon.com/appsync/latest/devguide/resolver-mapping-template-reference-lambda.html
type Invocation struct {
	Field   string  `json:"filed"`
	Context Context `json:"context"`
}

//https://docs.aws.amazon.com/appsync/latest/devguide/resolver-context-reference.html
//"arguments" : { ... },
//"source" : { ... },
//"result" : { ... },
//"identity" : { ... },
//"request" : { ... }
type Context struct {
	Arguments Arguments `json:"arguments"`
	Source    Source    `json:"source"`
}

// AWS Documentation: A map containing all GraphQL arguments for this field.
type Arguments struct {
}

// AWS Documentation: A map containing the resolution of the parent field.
type Source struct {
}
