FROM golang:1.23-alpine3.20

WORKDIR /app

RUN go install github.com/go-task/task/v3/cmd/task@latest

COPY go.mod Taskfile.yml ./

WORKDIR /app/test

ENTRYPOINT [ "task", "test" ]