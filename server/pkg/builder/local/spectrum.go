package local

import (
	"encoding/json"
	"github.com/container-tools/boxit/api"
	builderapi "github.com/container-tools/boxit/server/pkg/builder/api"
	spectrum "github.com/container-tools/spectrum/pkg/builder"
	"os"
	"path/filepath"
	"sort"
)

func publish(options builderapi.BuilderOptions, img api.ImageRequest, libDir string) (api.ImageResult, error) {
	imageID := getImageID(options, img)
	res := api.ImageResult{
		ID: imageID,
	}
	var files []string
	filepath.Walk(libDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	sort.Strings(files)

	targetDestination := "/deployment/dependencies"
	for _, f := range files {
		name := filepath.Base(f)
		a := api.Artifact{
			Location: filepath.Join(targetDestination, name),
		}
		res.Artifacts = append(res.Artifacts, a)
	}

	serial, err := json.Marshal(res)
	if err != nil {
		return res, err
	}

	opts := spectrum.Options{
		PushInsecure: options.Insecure,
		Base:         builderapi.BaseImage(img.Platform),
		Target:       imageID,
		Stdout:       os.Stdout,
		Stderr:       os.Stderr,
		Annotations: map[string]string{
			descriptorAnnotation: string(serial),
		},
	}
	_, err = spectrum.Build(opts, libDir+":"+targetDestination)
	if err != nil {
		return res, err
	}
	return res, nil
}
