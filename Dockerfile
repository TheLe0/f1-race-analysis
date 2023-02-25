FROM golang:1.19.4-buster

WORKDIR /f1-race-analysis

COPY go.mod go.mod
COPY go.sum go.sum
COPY main.go main.go
COPY commands commands
COPY infrastructure infrastructure
COPY repository repository
COPY types types
COPY utils utils
COPY data data

RUN go mod download
RUN go install
RUN go build

CMD  ["./f1-race-analysis"]