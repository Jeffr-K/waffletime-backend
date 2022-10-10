FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o /waffletime

EXPOSE 6060

CMD ["/waffletime"]



##  docker build --tag waffletime:v1 . -f ./.docker/Dockerfile
#
#FROM golang:alpine AS builder
#
#ENV GO111MODULE=on \
#    CGO_ENABLED=0 \
#    GOOS=linux \
#    GOARCH=amd64
#
#WORKDIR /app
#
##COPY ./go.mod ./
##COPY ./cmd/main.go ./
##COPY ./cmd/bootstrap.go ./
#
#COPY . /
#
#RUN go mod tidy
#
#RUN go mod download
#
#RUN go build -o main .
#
#WORKDIR /dist
#RUN cp /app/main .
#
#FROM scratch
#COPY --from=builder dist/main .
#ENTRYPOINT ["/main"]