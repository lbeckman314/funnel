---
title: Development
menu:
  main:
    identifier: development
    weight: -30
    
---
# Development

Funnel uses:

- [Go][go] for the majority of the code.
- [Task Execution Schemas][tes] for task APIs.
- [Protobuf][protobuf] + [gRPC][grpc] for RPC communication.
- [gRPC Gateway][gateway] for HTTP communication.
- [Angular][angular] and [SASS][sass] for the web dashboard.
- [GNU Make][make] for development tasks.
- [Docker][docker] for executing task containers.
- [Python][python] for end-to-end conformance tests.
- [vendetta][vendetta] for Go dependency vendoring.
- and more.

## Prerequisites

These are the tools you'll need to install:

- [Go 1.8][go]
- [Make][make]
- [Docker][docker] (tested with v1.12, v1.13)
- [Protocol Buffers][protobuf] if making changes to the schema.
- [NodeJS][node] and [NPM][npm] for web dashboard development.

## Build

Most development tasks are run through `make`.

|Command|Description|
|---|---|
|`make`              | Build the code.
|`make test`         | Run both unit and end-to-end tests.
|`make go-test`      | Run Go unit/integration tests.
|`make go-test-short`| Run only fast-running Go tests.
|`make proto`        | Regenerate code from protobuf schemas (requires protoc)
|`make tidy`         | Reformat code
|`make lint`         | Run code style and other checks.
|`make full`         | Run all steps needed to check the code before making a pull request.
|`make clean`        | Remove build/development files.
|`make add_deps`     | Add new vendored dependencies.
|`make prune_deps`   | Prune unused vendored dependencies.
|`make serve-doc`    | Serve API reference (godoc) docs on `localhost:6060`
|`make web`          | Build the web dashboard Javascript/CSS bundle.
|`make cross-compile`| Build binaries for all OS/Architectures.
|`make gce-installer`| Build the GCE image installer.
|`make gen-mocks`    | Generate mocks for testing.
|`make website-dev`   | Serve the Funnel website on localhost:1313
|`make upload-dev-release`| Upload a development release to GitHub.
|`make bundle-examples`| Bundle example task messages into Go code.

## Source

| Directory | Description |
|---|---|
|`cmd`         | Funnel command line interface.
|`config`      | Configuration parsing, loading, etc.
|`proto/tes`   | Generated GA4GH protobuf/gRPC files from [task-execution-schemas][tes].
|`proto/funnel`| The internal, Funnel-specific protobuf/gRPC files.
|`logger`      | Logging.
|`scheduler`   | Scheduler logic and backends.
|`server`      | Database and server implementing the [TES API][tes] and Scheduler RPC.
|`storage`     | Filesystem support, e.g. local, Google Cloud Storage, S3, etc.
|`worker`      | Worker process: task runner, docker, file mapper, etc.
|`web`         | Javascript, CSS, HTML for web dashboard.

## Go Tests

Run all tests: `make go-test`   
Run the scheduler tests: `go test ./scheduler/...`  
Run the worker tests with "Cancel" in the name: `go test ./worker -run Cancel`  

You get the idea. See the `go test` docs for more.

## Mocking

The [testify][testify] and [mockery][mockery] tools are used to generate and use
mock interfaces in test code, for example, to mock the Google Cloud APIs.

## Python Tests

There are end-to-end tests written in python.
These are heavyweight integration tests which start Funnel servers
and workers and submit tasks.

These are run with `make test`.

## Vendoring

Go dependencies are vendored under /vendor.  
Don't manually add new submodules, use `make add_deps`.

## Submodules

Funnel has git submodules. The Makefile usually handles this for you, but if needed,
`git submodule update --init --recursive` will get all the submodules.


[go]: https://golang.org
[angular]: https://angularjs.org/
[protobuf]: https://github.com/google/protobuf
[grpc]: http://www.grpc.io/
[sass]: http://sass-lang.com/
[make]: https://www.gnu.org/software/make/
[docker]: https://docker.io
[python]: https://www.python.org/
[vendetta]: https://github.com/dpw/vendetta
[node]: https://nodejs.org
[npm]: https://www.npmjs.com/
[gateway]: https://github.com/grpc-ecosystem/grpc-gateway
[tes]: https://github.com/ga4gh/task-execution-schemas
[testify]: https://github.com/stretchr/testify
[mockery]: https://github.com/vektra/mockery