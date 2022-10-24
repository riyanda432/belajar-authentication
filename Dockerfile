FROM golang:1.17-alpine AS build-env

LABEL maintainer="Riyanda Febri"

ENV APP_NAME=BELAJAR_AUTHENTICATION

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

WORKDIR /src

RUN ls -ls

RUN mkdir -p /src/belajar_authentication
COPY . /src/belajar-authentication
WORKDIR /src/belajar-authentication

RUN go install github.com/gobuffalo/pop/soda@latest

RUN soda migrate up

RUN go mod tidy -compat=1.17

RUN mkdir -p bin
RUN go build

RUN chmod 755 ./belajar-authentication

EXPOSE 8080

CMD "./belajar-authentication"