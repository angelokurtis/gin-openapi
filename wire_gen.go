// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package openapi

import (
	"github.com/angelokurtis/gin-openapi/internal/document"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Injectors from wire.go:

func newV3RouterBuilder(router gin.IRouter) (*v3RouterBuilder, error) {
	swaggerLoader := openapi3.NewSwaggerLoader()
	localLoader := document.NewLocalLoader(swaggerLoader)
	client := &http.Client{}
	remoteLoader := document.NewRemoteLoader(client, swaggerLoader)
	genericLoader := document.NewGenericLoader(localLoader, remoteLoader)
	defaultHolder := document.NewDefaultHolder(genericLoader)
	openapiV3RouterBuilder := &v3RouterBuilder{
		holder: defaultHolder,
	}
	return openapiV3RouterBuilder, nil
}
