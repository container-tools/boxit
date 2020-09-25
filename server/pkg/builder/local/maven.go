package local

import (
	"fmt"
	"github.com/container-tools/boxit/api"
	"github.com/container-tools/boxit/server/pkg/util/maven"
	"io/ioutil"
	"strings"
	"time"
)

func mavenBuild(img api.ImageRequest) (root string, path string, err error) {
	root, err = ioutil.TempDir("", "boxit-*")
	if err != nil {
		return "", "", err
	}
	path = "target/dependencies"

	p := maven.NewProjectWithGAV("io.container-tools", "project", "1.0")

	p.Dependencies = make([]maven.Dependency, 0)
	p.DependencyManagement = &maven.DependencyManagement{Dependencies: make([]maven.Dependency, 0)}

	for _, depStr := range img.Dependencies {
		if strings.HasPrefix(string(depStr), "mvn:") {
			mid := strings.TrimPrefix(string(depStr), "mvn:")
			gav := strings.Replace(mid, "/", ":", -1)

			p.AddEncodedDependencyGAV(gav)
		} else {
			return root, path, fmt.Errorf("unsupported dependency type %q", depStr)
		}
	}

	mc := maven.NewContext(root, p)
	mc.AdditionalArguments = append(mc.AdditionalArguments, "dependency:copy-dependencies", "-DoutputDirectory="+path)
	mc.Timeout = 5 * time.Minute

	if err := maven.Run(mc); err != nil {
		return root, path, err
	}

	return root, path, nil
}
