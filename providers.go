package openapi

import (
	"github.com/angelokurtis/gin-openapi/internal/document"
	"github.com/angelokurtis/gin-openapi/internal/http"
	"github.com/google/wire"
)

var providers = wire.NewSet(
	document.Providers,
	http.Providers,
	wire.Struct(new(v3RouterBuilder), "*"),
)
