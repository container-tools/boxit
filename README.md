# Boxit

Boxit is a distributed container image builder that can dynamically produce
container images from a set of dependencies given as input.

Supported platforms are currently:
- jvm

## 1. Compiling

Source code can be compiled with Go 1.14, using the command:

```shell
make
```

This will produce two binaries:
- `boxit-server`: the server part of the distributed builder
- `boxit`: the client part

## 2. Running

### 2.1 Running the registry

Boxit needs a registry to host the produced images.
To run a local registry using docker, you can run:

```shell
docker run -d -p "5000:5000" --name "registry" registry:2
```

### 2.2 Running the server

The server can be simply executed via:

```shell
./boxit-server
```

It runs on port `8080` and points by default to the registry on `localhost:5000` (insecure).

### 2.3 Creating images with the client

The `boxit` client can be now used to ask the server to create new images (uses server on `http://localhost:8080` by default):

```shell
./boxit create -d mvn:org.apache.camel:camel-core:3.5.0./boxit create -d mvn:org.apache.camel:camel-core:3.5.0 -d mvn:org.apache.camel:camel-timer:3.5.0
```

A possible output may be: `localhost:5000/boxit/vb7uwws6n2mwbc3h6fbctwuhkmmmh2cl6a2v45r5g23rbcsndcuha:latest`.
A new image will be created if not already present. The ID of the image containing the dependencies is always printed.

The first time the command is executed, it will take several seconds to have
an image ready. Once the image is build, subsequent calls will be completed immediately.
 
