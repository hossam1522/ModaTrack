FROM golang:1.23-alpine3.20

WORKDIR /app

RUN addgroup -S -g 1001 test && \
    adduser -S -u 1001 -G test test && \
    chown -R test:test /app

USER test

RUN go install github.com/go-task/task/v3/cmd/task@latest

COPY go.mod Taskfile.yml ./

WORKDIR /app/test

ENTRYPOINT [ "task", "test" ]