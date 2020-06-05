
default: build

build: build-api build-client build-server build-cli

build-api:
	cd api && go build ./...

build-client:
	cd client && go build ./...

build-server:
	cd server && go build -o ../boxit-server ./cmd/server

build-cli:
	cd cli && go build -o ../boxit ./cmd/cli
