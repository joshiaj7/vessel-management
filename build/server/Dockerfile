FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod vendor

COPY . /app

RUN GOOS=linux GOARCH=amd64 go build -o /app/bin/vessel-management /app/cmd/gateway/main.go

CMD ["/app/bin/vessel-management"]
