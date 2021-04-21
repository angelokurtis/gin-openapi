package openapi

import (
	"github.com/angelokurtis/gin-openapi/internal/document"
	"github.com/gin-gonic/gin"
)

type (
	RouteBuilder interface {
		Operation(operationID string) Operation
		BuildRoutes() (*gin.RouterGroup, error)
	}
	Operation interface {
		Handler(handler gin.HandlerFunc)
	}
	v3RouterBuilder struct {
		holder document.Holder
	}
)

func NewRouteBuilder(router gin.IRouter, uri string) (RouteBuilder, error) {
	return nil, nil
}
