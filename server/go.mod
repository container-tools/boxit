module github.com/container-tools/boxit/server

go 1.13

require (
	github.com/container-tools/boxit/api v0.0.0
	github.com/container-tools/spectrum v0.3.3
	github.com/docker/cli v0.0.0-20200210162036-a4bedce16568 // indirect
	github.com/docker/docker v1.13.1 // indirect
	github.com/google/go-containerregistry v0.1.2
	github.com/opencontainers/runc v1.0.0-rc2.0.20190611121236-6cc515888830 // indirect
	github.com/pkg/errors v0.9.1
	github.com/scylladb/go-set v1.0.2
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.5.1
)

replace github.com/container-tools/boxit/api => ../api
