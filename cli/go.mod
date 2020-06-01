module github.com/nicolaferraro/boxnet/cli

go 1.13

require (
	github.com/nicolaferraro/boxnet/api v0.0.0
	github.com/nicolaferraro/boxnet/client v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.0.0
)

replace (
	github.com/nicolaferraro/boxnet/api => ../api
	github.com/nicolaferraro/boxnet/client => ../client
)
