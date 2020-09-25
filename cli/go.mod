module github.com/container-tools/boxit/cli

go 1.13

require (
	github.com/container-tools/boxit/api v0.0.0
	github.com/container-tools/boxit/client v0.0.0
	github.com/spf13/cobra v1.0.0
)

replace (
	github.com/container-tools/boxit/api => ../api
	github.com/container-tools/boxit/client => ../client
)
