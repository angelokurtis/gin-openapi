package document

import (
	"context"
	"github.com/getkin/kin-openapi/openapi3"
	"io/ioutil"
	"net/http"
)

type RemoteLoader struct {
	client *http.Client
	loader *openapi3.SwaggerLoader
}

func NewRemoteLoader(client *http.Client, loader *openapi3.SwaggerLoader) *RemoteLoader {
	return &RemoteLoader{client: client, loader: loader}
}

func (l *RemoteLoader) LoadOpenAPI(uri string) (*openapi3.Swagger, error) {
	b, err := l.download(uri)
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

func (l *RemoteLoader) download(uri string) ([]byte, error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json, application/yaml, application/x-yaml")

	res, err := l.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
