module github.com/nicolaferraro/boxnet/server

go 1.13

require (
	github.com/apache/camel-k/pkg/apis/camel v0.0.0
	github.com/apache/camel-k/pkg/client/camel v0.0.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/nicolaferraro/boxnet/api v0.0.0
	k8s.io/apimachinery v0.16.4
	k8s.io/client-go v0.16.4
	sigs.k8s.io/controller-runtime v0.4.0
)

replace (
	github.com/apache/camel-k/pkg/apis/camel => github.com/apache/camel-k/pkg/apis/camel v0.0.0-20200529095204-31c39d23e408
	github.com/nicolaferraro/boxnet/api => ../api
)
