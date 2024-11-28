FROM golang:1.23.3-alpine

WORKDIR /app

RUN adduser -D test

USER test

ENV GOCACHE=/home/test/app/.cache/go-build

RUN go install github.com/go-task/task/v3/cmd/task@latest && \
    chmod -R 777 /home/test/app/.cache

COPY go.mod Taskfile.yml ./

WORKDIR /app/test

ENTRYPOINT [ "task", "test" ]