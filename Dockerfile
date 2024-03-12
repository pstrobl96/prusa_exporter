# syntax=docker/dockerfile:1

FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

COPY *.go ./

RUN apk add gcc musl-dev

RUN CGO_ENABLED=1 go build -race -v -o /prusa_exporter

FROM alpine:latest

COPY --from=builder /prusa_exporter .

EXPOSE 10009

ENTRYPOINT ["/prusa_exporter"]