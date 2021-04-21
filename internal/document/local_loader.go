package document

import (
	"context"
	"github.com/getkin/kin-openapi/openapi3"
	"io/ioutil"
)

type LocalLoader struct {
	loader *openapi3.SwaggerLoader
}

func NewLocalLoader(loader *openapi3.SwaggerLoader) *LocalLoader {
	return &LocalLoader{loader: loader}
}

func (l *LocalLoader) LoadOpenAPI(uri string) (*openapi3.Swagger, error) {
	b, err := ioutil.ReadFile(uri)
	if err != nil {
		return nil, err
	}

	swagger, err := l.loader.LoadSwaggerFromData(b)
	if err != nil {
		return nil, err
	}

	err = swagger.Validate(context.TODO())
	if err != nil {
		return nil, err
	}

	return swagger, nil
}
