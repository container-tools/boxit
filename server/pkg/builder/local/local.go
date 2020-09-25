package local

import (
	"encoding/json"
	"fmt"
	"github.com/container-tools/boxit/api"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/google/go-containerregistry/pkg/v1/remote/transport"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

const (
	registry   = "localhost:5000"
	repository = "boxit"

	descriptorAnnotation = "boxit.descriptor"
)

func Builder(img api.ImageRequest) (api.ImageResult, error) {
	existing, err := findExisting(img)
	if err != nil {
		return api.ImageResult{}, err
	} else if existing != nil {
		return *existing, nil
	}

	return buildNewImage(img)
}

func findExisting(img api.ImageRequest) (*api.ImageResult, error) {
	imageID := getImageID(img)
	ref, err := name.ParseReference(imageID)
	if err != nil {
		return nil, err
	}
	desc, err := remote.Get(ref)
	if err != nil {
		if te, ok := err.(*transport.Error); ok {
			if te.StatusCode >= 400 && te.StatusCode < 500 {
				fmt.Sprintf("Cannot get the remote image: %v", err)
				return nil, nil
			}
		}
		return nil, err
	}

	var annotation string
	remoteImage, err := desc.Image()
	if err != nil {
		return nil, err
	}
	manifest, err := remoteImage.Manifest()
	if err != nil {
		return nil, err
	}
	for _, layer := range manifest.Layers {
		if val, ok := layer.Annotations[descriptorAnnotation]; ok {
			annotation = val
		}
	}
	if annotation == "" {
		return nil, nil
	}

	var res api.ImageResult
	if err := json.Unmarshal([]byte(annotation), &res); err != nil {
		return nil, errors.Wrapf(err, "could not decode descriptor from %q annotation", descriptorAnnotation)
	}

	return &res, nil
}

func getImageID(img api.ImageRequest) string {
	hash := img.Hash()
	version := "latest"
	return fmt.Sprintf("%s/%s/%s:%s", registry, repository, hash, version)
}

func buildNewImage(img api.ImageRequest) (api.ImageResult, error) {
	root, path, err := mavenBuild(img)
	if err != nil {
		return api.ImageResult{}, err
	}
	defer os.RemoveAll(root)

	return publish(img, filepath.Join(root, path))
}
