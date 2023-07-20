# Comet

Comet is a generalized key-value REST service on top of RocksDB databases. It is meant as a cheap and simple KV datastore for applications that serve only a few terabytes of data but cannot fit all of them in Redis RAM or do not need a full database cluster and its attendant cost. The project develop using Fiber framework on API side and Gin framework on Puller with applying clean architecture.

## Features

* *Speed* - A single instance of Comet can handle ~63K `get` operations per second at an average latency of 1 ms or less.
This was on a medium-sized GCP compute engine (`n2-standard-8`, 8 cores, 32 GB RAM, 1 local NVMe SSD) hosting 110 GB of
application data. Higher throughput can be achieved by hosting on larger instances as necessary.

* *Binary format* - Everything that's stored in RocksDB are just uninterpreted byte arrays and Comet
makes use of this in order to store and serve any content type -- plain text, JSON documents, images, Protobuf, etc.

* *Bulk loading* - For use cases where a REST endpoint just acts as a read-only serving layer for data that is prepared
elsewhere, SST files can be generated offline and ingested in Comet.

* *Authentication* - Comet uses configuration-based bearer tokens to secure access to both the KV and operator endpoints.
endpoints.

* *Instrumentation* - A Prometheus endpoint is provided at `/metrics`. There's also `/healthcheck` for liveness checks.

## Template Structure

### API

* [Fiber](https://github.com/gofiber/fiber) is an Express inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go. Designed to ease things up for fast development with zero memory allocation and performance in mind.
* [Wire](https://github.com/google/wire) is a code generation tool that automates connecting components using dependency injection.
* [grocksdb](https://github.com/linxGnu/grocksdb) is a RocksDB wrapper for Go.
* [validator](github.com/go-playground/validator) is a package validator implements value validations for structs and individual fields based on tags.
* [zap](https://github.com/uber-go/zap) is a Blazing fast, structured, leveled logging in Go.
* [Viper](https://github.com/spf13/viper) is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats.

### Puller

* [Gin](github.com/gin-gonic/gin) is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.
* [Wire](https://github.com/google/wire) is a code generation tool that automates connecting components using dependency injection.
* [gocron](https://github.com/go-co-op/gocron) is a job scheduling package which lets you run Go functions at pre-determined intervals using a simple, human-friendly syntax.
* [validator](github.com/go-playground/validator) is a package validator implements value validations for structs and individual fields based on tags.
* [zap](https://github.com/uber-go/zap) is a Blazing fast, structured, leveled logging in Go.
* [Viper](https://github.com/spf13/viper) is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats.

## Using `comet` project

### Prerequisite

1. Follow rocksdb install guide [rocksdb](https://github.com/facebook/rocksdb/blob/main/INSTALL.md)

2. Install ALL dependencies compression libraries and tools

* You can link RocksDB with following compression libraries:
  * [zlib](http://www.zlib.net/) - a library for data compression.
  * [bzip2](http://www.bzip.org/) - a library for data compression.
  * [lz4](https://github.com/lz4/lz4) - a library for extremely fast data compression.
  * [snappy](http://google.github.io/snappy/) - a library for fast
        data compression.
  * [zstandard](http://www.zstd.net) - Fast real-time compression
        algorithm.

* All our tools depend on:
  * [gflags](https://gflags.github.io/gflags/) - a library that handles
        command line flags processing. You can compile rocksdb library even
        if you don't have gflags installed.

3. Export environment

```bash
export CGO_CFLAGS="-I/opt/homebrew/opt/rocksdb/include"
export CGO_LDFLAGS="-L/opt/homebrew/opt/rocksdb/lib -lrocksdb -lstdc++ -lm -lz -lsnappy -llz4 -lzstd" 
export LIBRARY_PATH="/opt/homebrew/opt/snappy/lib"
```

### Run application

To use `comet` project, follow these steps:

```bash
# Navigate into the project
cd ./comet/api or ./comet/puller 

# Install dependencies
make deps

# Generate wire_gen.go for dependency injection
# Please make sure you are export the env for GOPATH
make wire

# Run the project in Development Mode
make run
```

Additional commands:

```bash
➔ make help
build                          Compile the code, build Executable File
run                            Start application
test                           Run tests
test-coverage                  Run tests and generate coverage file
deps                           Install dependencies
deps-cleancache                Clear cache in Go module
wire                           Generate wire_gen.go
help                           Display this help screen
```

## Endpoint

### User endpoint

Main user-facing endpoint that provides basic Key-Value functionality

#### `PUT /databases/{db}/keys/{key}`

* Put a value to database, the value can be any format specified by the body.

```
curl -v --request PUT http://localhost:8000/databases/database-a/keys/key1 \
--data 'this is a value'
```

#### `GET /databases/{db}/keys/{key}`

* Get a value by key.

```
curl -v --request GET http://localhost:8000/databases/database-a/keys/key1
```

#### `DELETE /databases/{db}/keys/{key}`

* Delete a key.

```
curl -v --request DELETE http://localhost:8000/databases/database-a/keys/key1
```

### Operator endpoint

Operator endpoint requires admin access and every request must pass the JWT Bearer token in header.

#### `POST /admin/databases/{db}/catch-up-with-primary`

* Update database to match the referenced database. The selected database must be opened as secondary.

#### `GET /admin/databases/{db}/properties/{property-name}`

* Get property of the database, where `property-name` is one of those listed [here](https://github.com/facebook/rocksdb/blob/08809f5e6cd9cc4bc3958dd4d59457ae78c76660/include/rocksdb/db.h#L428-L634).

#### `PUT /admin/databases/{db}/checkpoint/{directory}`

* Create checkpoint of selected database at `directory` where the directory is relative to the configured `app.backup-path`.

#### `PUT /admin/databases/{db}/ingest/{directory}`

* Ingest sst files into selected database where `directory` is relative to the configured `app.ingest-path`.

## Folder Structure

This project design by using clean architecture and hexagonal architecture so folder of project will organize base on
clean architecture below

Ref: <https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html>

Here below is how folder map to layer and component in clean architecture

* domain -> Entity
* usecase -> Usecase
* repository -> Repository
* api -> Handler
* driver -> remote call
