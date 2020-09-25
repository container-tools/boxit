package local

import (
	"github.com/container-tools/boxit/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocalBuild(t *testing.T) {
	_, err := Builder(api.ImageRequest{
		Platform: api.PlatformJVM,
		Dependencies: []api.Dependency{
			"mvn:org.apache.camel:camel-telegram:3.5.0",
		},
	})
	assert.NoError(t, err)
}
