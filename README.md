
# Local

## Server
```
$ go mod tidy
$ go run main.go
```

## Client
```
$ brew install evans
$ evans proto/hello-world.proto
```

# Docker

## Build Docker Image
```
$ docker pull ubuntu:latest
$ docker build --build-arg GO_VERSION=1.20.3 -t grpctest/helloworld .
```

## Server
```
$ docker run -p 50051:50051 grpctest/helloworld
```

## Client
```
$ brew install evans
$ evans proto/hello-world.proto
```

# Kubernetes Istio

## Server
```
$ kubectl apply -f manifests/
```

## Client
```
$ evans proto/hello-world.proto -p 80
```