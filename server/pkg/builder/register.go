package builder

import (
	"github.com/container-tools/boxit/server/pkg/builder/api"
	"github.com/container-tools/boxit/server/pkg/builder/local"
)

//var Default api.Builder = kubernetes.Builder
var Default api.Builder = local.Builder
