FROM golang:1.19-alpine

WORKDIR /src

COPY go.mod .
COPY go.sum .

RUN go mod vendor

COPY . /src

ARG TARGETOS TARGETARCH

RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /src/bin/script_main /src/cmd/script/main.go
