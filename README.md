# gokita
go-kit-a[pp] - A [go-kit](https://gokit.io/) boiler plate

- [Introduction](#introduction)
- [Motivation](#motivation)
- [What's Included](#whats-included)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [Epilogue](#epilogue)

### Introduction
[Go kit] is a toolkit for microservices. It provides guidance and solutions for
most of the common operational and infrastructural concerns. Allowing you to
focus your mental energy on your business logic. It provides the building blocks
for separating transports from business domains; making it easy to switch one
transport for the other. Or even service multiple transports for a service at
once.

### Motivation
One major gripe users of [Go kit] have is that When using the go-kit, the overhead 
of adding API you your service is very high. You need to add a lot of code, which is 
mostly copy-paste of other APIs and there are too many places to make mistakes.

This has scared of some developers from using it and even though there has been some 
work done around generating much of the boilerplate automatically, the tools listed below 
don't get the job done as one would hope.

Also most of them haven't had a commit done to them in a while which means they might not be 
compatible with some of the packages go-kit uses.

- https://github.com/digimuza/go-kit-seed-microservice-generator
- https://github.com/kujtimiihoxha/kit
- https://github.com/kujtimiihoxha/gk
- https://github.com/go-kit/kit/tree/master/cmd/kitgen

Once we decided to use [Go kit] for our microservices we knew we could not rely on any of the 
above "generators" instead we used a combination of them to create a [Go kit] base app that will 
serve as a starting point for each microservice we create.

### What's Included

We've tried to make this boilerplate cover our most basic needs for we define as a microservice 
that can run in production.

#### Core

- Transport layer (grpc)
- Endpoint layer
- Service layer

#### Middleware

- Endpoint Logging Middleware
- Endpoint Instrumenting Middleware
- Service Logging Middleware
- Service Instrumenting Middleware

#### Observability & Tracing

- Prometheus
- Zipkin

#### Misc

- Docker Integration (Dockerfile + docker-compose.yml)

### Project Structure

```shell script
.
├── Gopkg.lock
├── Gopkg.toml
├── README.md
├── build
│   └── Dockerfile
├── cmd
│   ├── main.go
│   └── service
│       ├── service.go
│       └── service_gen.go
├── docker-compose.yml
├── pkg
│   ├── endpoint
│   │   ├── endpoint.go
│   │   ├── endpoint_gen.go
│   │   └── middleware.go
│   ├── service
│   │   ├── middleware.go
│   │   └── service.go
│   └── transport
│       ├── grpc
│       │   ├── handler.go
│       │   └── handler_gen.go
│       └── http # TODO
├── proto
│   ├── compile.sh
│   ├── gokita.pb.go
│   └── gokita.proto
```

### Getting Started

To get gokita up and running you'll need to have a few things installed beforehand:

1. [Go](https://golang.org/doc/install)
2. [Go Dep](https://golang.github.io/dep/docs/installation.html)
3. [Protobuf](https://github.com/google/protobuf)

With the prerequisites out of the way;

```shell script
git clone https://github.com/SeamPay/gokita.git
cd gokita
docker-compose up --build
```

After you run docker-compose up your services will start up and any change you make to 
your code will automatically rebuild and restart your service.

You can use [BloomRPC](https://github.com/uw-labs/bloomrpc) to validate/test.

### Epilogue

This boilerplate is meant to be a template for hardened, production ready Go microservices. 
By no means is this a "one size fits all" template. There may be very small services that 
don't require the structure defined here.

Additionally, we don't claim to have all of the answers! If you have a better idea or practice 
please open an Issue against this repository with your suggestion and we will do our best to 
facilitate a discussion about your concern or suggestion.

Thanks!
