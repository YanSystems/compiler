# Compiler

![Tests](https://github.com/YanSystems/compiler/actions/workflows/tests.yml/badge.svg) [![codecov](https://codecov.io/gh/YanSystems/compiler/graph/badge.svg?token=JFN7EBA8GD)](https://codecov.io/gh/YanSystems/compiler) [![Go Report](https://goreportcard.com/badge/YanSystems/compiler)](https://goreportcard.com/report/YanSystems/compiler) [![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/YanSystems/compiler/blob/main/LICENSE)

A compiler microservice that exposes a REST API.

## Running the service locally

First, make sure you have go version `1.22.4` (check by running `go version`). Then, run the following make script,
```
make run
```

Alternatively, you can run the service in a container. First, build the docker image,
```
make image
```

To start the container, run `make up`. To stop it, run `make down`

## License

This compiler microservice is [MIT licensed.](https://github.com/YanSystems/compiler/blob/main/LICENSE)
