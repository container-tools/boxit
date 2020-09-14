package builder

import (
	"github.com/container-tools/boxit/server/pkg/builder/api"
	"github.com/container-tools/boxit/server/pkg/builder/kubernetes"
)

var Default api.Builder = kubernetes.Builder
