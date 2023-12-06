# syntax=docker/dockerfile:1

FROM golang:1.20.3-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /prusa_exporter

FROM alpine:latest

COPY --from=builder /prusa_exporter .

EXPOSE 10009

ENTRYPOINT ["/prusa_exporter"]