FROM golang:1.23-alpine3.20

WORKDIR /usr/src/app

COPY go.mod go.sum Taskfile.yml ./
RUN task install && task update