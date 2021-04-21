package document

import (
	"github.com/google/wire"
)
import "github.com/getkin/kin-openapi/openapi3"

var Providers = wire.NewSet(
	NewDefaultHolder,
	NewGenericLoader,
	NewLocalLoader,
	NewRemoteLoader,
	openapi3.NewSwaggerLoader,
)
