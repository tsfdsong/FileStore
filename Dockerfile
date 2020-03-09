# Start from golang v1.11 base image
FROM golang:1.12-alpine as base_builder

RUN apk add bash --no-cache ca-certificates gcc g++
#RUN apk add bash --no-cache ca-certificates git gcc g++ libc-dev

WORKDIR /FileStore

COPY go.mod go.sum ./

# Force the go compiler to use modules
ENV GO111MODULE=on

ENV GOPROXY=https://goproxy.io

#This is the ‘magic’ step that will download all the dependencies that are specified in 
# the go.mod and go.sum file.
# Because of how the layer caching system works in Docker, the  go mod download 
# command will _ only_ be re-run when the go.mod or go.sum file change 
# (or when we add another docker instruction this line)
RUN go mod download

# This image builds the weavaite server
FROM base_builder as builder

# Set the Current Working Directory inside the container
WORKDIR /FileStore

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

WORKDIR /FileStore/main

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o /go/bin/app .
#RUN go build -o /go/bin/app .

######## Start a new stage from scratch #######
FROM alpine:latest 

RUN apk add bash --no-cache ca-certificates

WORKDIR /root

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/bin/app .

# Document that the service listens on port 8080.
EXPOSE 23000

CMD ["./app"]  