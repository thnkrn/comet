FROM golang:1.20.4-alpine3.16 AS build_base
RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/comet-puller

# Populate the module cache based on the go.{mod,sum} files.
COPY . .
RUN go get -u -t -d -v ./... && go mod download && go mod tidy && go mod vendor

# Build the Go app
RUN go build -o ./out/comet-puller -v ./cmd/api

# Start fresh from a smaller image
FROM alpine:3.16
RUN apk add --no-cache ca-certificates iwatch=0.2.2-r0
COPY --from=build_base /tmp/comet-puller/out/comet-puller /app/comet-puller
COPY ./config.yml /app
COPY ./key.json /app

USER root
# This container exposes port 8080 to the outside world
EXPOSE 8081

WORKDIR /app

# Run the binary program produced by `go install`
CMD ["./comet-puller"]