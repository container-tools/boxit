package api

import (
	"github.com/container-tools/boxit/api"
)

type Builder func(img api.ImageRequest) (api.ImageResult, error)

const BuildFailure = "image build failed"
