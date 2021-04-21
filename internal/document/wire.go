//+build wireinject

package document

import (
	"github.com/angelokurtis/gin-openapi/internal/http"
	"github.com/google/wire"
)

func mockRemoteLoader() Loader {
	wire.Build(
		Providers,
		http.Providers,
		wire.Bind(new(Loader), new(*RemoteLoader)),
	)
	return nil
}

func mockLocalLoader() Loader {
	wire.Build(
		Providers,
		wire.Bind(new(Loader), new(*LocalLoader)),
	)
	return nil
}
