package document

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoteLoader_LoadOpenAPI(t *testing.T) {
	loader := mockRemoteLoader()
	swagger, err := loader.LoadOpenAPI("https://raw.githubusercontent.com/OAI/OpenAPI-Specification/master/examples/v3.0/petstore.yaml")
	assert.NoError(t, err)
	assert.NotNil(t, swagger)
	assert.Len(t, swagger.Paths, 2)
}
