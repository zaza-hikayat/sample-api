FROM golang:1.21.6-alpine AS go-api

WORKDIR /app
ADD . /app

RUN go build -o main main.go
RUN chmod 755 ./main

EXPOSE ${APP_PORT}
