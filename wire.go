//+build wireinject

package openapi

import (
	"github.com/angelokurtis/gin-openapi/internal/document"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func newV3RouterBuilder(router gin.IRouter) (*v3RouterBuilder, error) {
	wire.Build(
		providers,
		wire.Bind(new(document.Loader), new(*document.GenericLoader)),
		wire.Bind(new(document.Holder), new(*document.DefaultHolder)),
	)
	return nil, nil
}
