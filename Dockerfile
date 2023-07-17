# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

MAINTAINER Pavel Strobl "mail@pubel.dev"

#WORKDIR /app

#COPY go.mod ./
#COPY go.sum ./
#
#RUN go mod download
#
#COPY *.go ./
#
#RUN go build -o /buddy-prometheus-exporter

EXPOSE 10009

CMD [ "/buddy-prometheus-exporter" ]