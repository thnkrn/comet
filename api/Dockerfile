FROM golang:1.20.4-alpine3.16 AS build_base

ARG ROCKSDB_VERSION="8.1.1"

RUN apk add --update --no-cache build-base linux-headers git cmake bash perl
RUN apk add --update --no-cache zlib zlib-dev bzip2 bzip2-dev snappy snappy-dev lz4 lz4-dev zstd zstd-dev

# Install latest gflags
RUN cd /tmp && \
  git clone https://github.com/gflags/gflags.git && \
  cd gflags && \
  mkdir build && \
  cd build && \
  cmake -DBUILD_SHARED_LIBS=1 -DGFLAGS_INSTALL_SHARED_LIBS=1 -DBUILD_gflags_LIBS=1 .. && \
  make -j32 install && \
  cd /tmp && \
  rm -R /tmp/gflags/

# Install Rocksdb
RUN cd /tmp && \
  git clone https://github.com/facebook/rocksdb.git && \
  cd rocksdb && \
  git checkout v${ROCKSDB_VERSION} && \
  USE_SSE=1 PORTABLE=1 make -j8 shared_lib && \
  mkdir -p /usr/local/rocksdb/lib && \
  mkdir /usr/local/rocksdb/include && \
  cp librocksdb.so* /usr/local/rocksdb/lib && \
  cp /usr/local/rocksdb/lib/librocksdb.so* /usr/lib/ && \
  cp -r include /usr/local/rocksdb/ && \
  cp -r include/* /usr/include/ && \
  rm -R /tmp/rocksdb/

# Set the Current Working Directory inside the container
WORKDIR /tmp/comet

# Populate the module cache based on the go.{mod,sum} files.
COPY go.mod go.sum ./
RUN go get -u -t -d -v ./... && go mod download && go mod tidy && go mod vendor

COPY . .
# Build the Go app
RUN go build -o ./out/comet -v ./cmd/api

# Start fresh from a smaller image
FROM alpine:3.18
RUN apk add --no-cache ca-certificates iwatch=0.2.2-r0
RUN apk add --update --no-cache zlib zlib-dev bzip2 bzip2-dev snappy snappy-dev lz4 lz4-dev zstd zstd-dev

# Create app folder and user to give permission
RUN mkdir -p /app \
  && chown 1001 /app

# Copy required lib(cmake, gflags) for run time
COPY --from=build_base --chown=1001 /usr/local/lib /usr/local/lib/
# Copy rocksdb lib for run time
COPY --from=build_base --chown=1001 /usr/local/rocksdb /usr/local/rocksdb
# Copy config file
COPY ./config.yml /app
# Copy Go executable File
COPY --from=build_base --chown=1001 /tmp/comet/out/comet /app/comet

# This container exposes port 8080 to the outside world
EXPOSE 8080

WORKDIR /app

USER root

# Export lib gorocksdb path
ENV LD_LIBRARY_PATH="/usr/local/rocksdb/lib"

# Init fake process to grab PID 1 before the Fiber app does, https://github.com/gofiber/fiber/issues/1036
RUN apk add dumb-init
ENTRYPOINT ["/usr/bin/dumb-init", "--"]

# Run the binary program produced by `go install`
CMD ["./comet"]
