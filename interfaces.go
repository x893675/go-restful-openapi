package restfulspec

import "github.com/go-openapi/spec"

// If an (embedded) type of a Request or Response implements this
// then the definition builder will call this hook.
type SchemaBuilder interface {
	PostBuildOpenAPISchema(*spec.Schema)
}

// Documented is use to allow customisation of Swagger doc.
type Documented interface {
	SwaggerDoc() map[string]string
}
