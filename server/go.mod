module github.com/container-tools/boxit/server

go 1.13

require (
	github.com/apache/camel-k/pkg/apis/camel v1.1.1
	github.com/apache/camel-k/pkg/client/camel v1.1.1
	github.com/container-tools/boxit/api v0.0.0
	github.com/mitchellh/go-homedir v1.1.0
	k8s.io/apimachinery v0.18.2
    k8s.io/client-go v0.18.2
	sigs.k8s.io/controller-runtime v0.5.2
)

replace (
 	k8s.io/apimachinery => k8s.io/apimachinery v0.17.6
 	k8s.io/client-go => k8s.io/client-go v0.17.6

  	github.com/container-tools/boxit/api => ../api
)
