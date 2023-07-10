FROM golang:1.19

WORKDIR /backend

COPY go.mod .
COPY go.sum .

RUN go mod download
