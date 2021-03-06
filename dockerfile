# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app
ENV GO111MODULE=on
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /docker-gs-ping

CMD [ "/docker-gs-ping" ]
