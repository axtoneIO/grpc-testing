# gRPC-testing project
## _Implementing unit tests for our microservices_

[![N|Solid](https://cldup.com/dTxpPi9lDf.thumb.png)](https://nodesource.com/products/nsolid)

[![Build Status](https://github.com/axtoneIO)](https://github.com/axtoneIO)

This is a gRPC micro-services project that is interacting with a postgresql database and is responsible for managing an inventory of Rockets 🚀🚀 I started from scratch and incrementally added new features, the whole purpose of the project was creating a proper configuration for gRPC unit tests..

The project covers the following tools:
## Tech

Dillinger uses a number of open source projects to work properly:

- [gRPC](https://grpc.io/) - An open source universal RPC framework
- [Docker](https://www.docker.com/) - Developers Love Docker
- [Golang](https://go.dev/) - An open source language supported by Google

## Installation

Dillinger requires [Golang](https://go.dev/dl/) v1.16+ to run.

Install the dependencies:

```sh
go mod tidy
go mod vendor
```
## Docker

grpc-testing is very easy to install and deploy in a Docker container.

On both Dockerfile and docker-compose.yml you will find the whole configuration
for this project to work

```sh
docker-compose up --build
```
This will create the database and grpc-microservice containers
and pull in the necessary dependencies.

## License

MIT

**Free Software, Hell Yeah!**