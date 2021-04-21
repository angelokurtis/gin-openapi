package document

import (
	"github.com/angelokurtis/gin-openapi/test"
	"github.com/stretchr/testify/assert"
	"path"
	"testing"
)

func TestLocalLoader_LoadOpenAPI(t *testing.T) {
	loader := mockLocalLoader()
	swagger, err := loader.LoadOpenAPI(path.Join(test.Dir(), "petstore.yaml"))
	assert.NoError(t, err)
	assert.NotNil(t, swagger)
	assert.Len(t, swagger.Paths, 2)
}
