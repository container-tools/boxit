package api

import (
	"fmt"
	"github.com/container-tools/boxit/api"
)

type Builder func(img api.ImageRequest) (api.ImageResult, error)

const BuildFailure = "image build failed"

func BaseImage(platform api.Platform) string {
	switch platform {
	case api.PlatformJVM:
		return "adoptopenjdk/openjdk11:slim"
	}
	panic(fmt.Sprintf("no base image defined for platform %q", platform))
}
