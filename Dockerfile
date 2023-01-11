# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /go/bin/virtual-strike-backend-go

COPY go.mod ./

RUN go mod download

COPY . ./

RUN go build -o main ./cmd/api/main.go

EXPOSE 8836
ENTRYPOINT [ "./main" ]

