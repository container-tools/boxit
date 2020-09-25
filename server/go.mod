module github.com/container-tools/boxit/server

go 1.13

require (
	github.com/apache/camel-k v1.1.1
	github.com/apache/camel-k/pkg/apis/camel v1.1.1
	github.com/apache/camel-k/pkg/client/camel v1.1.1
	github.com/container-tools/boxit/api v0.0.0
	github.com/google/go-containerregistry v0.1.2
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/scylladb/go-set v1.0.2
	github.com/stretchr/testify v1.5.1
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.2
	k8s.io/client-go v12.0.0+incompatible
	sigs.k8s.io/controller-runtime v0.5.2
	github.com/container-tools/spectrum v0.3.3
)

replace (

	github.com/container-tools/boxit/api => ../api
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.6
	k8s.io/client-go => k8s.io/client-go v0.17.6
)
