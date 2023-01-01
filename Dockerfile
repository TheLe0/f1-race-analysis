FROM golang:1.19.4-buster

WORKDIR /f1-race-analysis

COPY go.mod go.mod
COPY go.sum go.sum
COPY main.go main.go
COPY data data

RUN go mod download
RUN go install
RUN go build

CMD  ["./f1-race-analysis"]