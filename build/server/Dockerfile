FROM golang:1.19-alpine

WORKDIR /src

COPY go.mod .

COPY go.sum .

RUN go mod vendor

COPY . /src

RUN GOOS=linux GOARCH=amd64 go build -o /src/bin/vessel-management /src/cmd/gateway/main.go

CMD ["/src/bin/vessel-management"]
