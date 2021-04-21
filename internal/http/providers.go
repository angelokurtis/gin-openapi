package http

import (
	"github.com/google/wire"
	"net/http"
)

var Providers = wire.NewSet(
	wire.Struct(new(http.Client)),
)
