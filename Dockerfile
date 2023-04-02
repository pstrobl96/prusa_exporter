# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /buddy-link-prometheus-exporter

CMD [ "/buddy-link-prometheus-exporter" ]