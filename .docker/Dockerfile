FROM golang:alpine AS builder

RUN mkdir /app

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod ./

COPY . .

RUN go mod download

RUN go build ./cmd/main.go

EXPOSE 8080

CMD ["go", "run", "cmd/main.go"]

##  docker build --tag waffletime:v1 . -f ./.docker/Dockerfile
