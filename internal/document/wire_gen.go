// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package document

import (
	"github.com/getkin/kin-openapi/openapi3"
	"net/http"
)

// Injectors from wire.go:

func mockRemoteLoader() Loader {
	client := &http.Client{}
	swaggerLoader := openapi3.NewSwaggerLoader()
	remoteLoader := NewRemoteLoader(client, swaggerLoader)
	return remoteLoader
}

func mockLocalLoader() Loader {
	swaggerLoader := openapi3.NewSwaggerLoader()
	localLoader := NewLocalLoader(swaggerLoader)
	return localLoader
}
