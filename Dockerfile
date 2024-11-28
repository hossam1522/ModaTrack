FROM debian:stable-slim

RUN apt-get update && \
    apt-get install -y --no-install-recommends wget ca-certificates && \ 
    rm -rf /var/lib/apt/lists/* && \
    useradd -m test

RUN wget -q https://go.dev/dl/go1.23.3.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.23.3.linux-amd64.tar.gz && \
    rm go1.23.3.linux-amd64.tar.gz && \
    apt-get purge -y wget

USER test
    
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/home/test/go
ENV GOCACHE=/home/test/app/.cache/go-build
ENV PATH=$PATH:$GOPATH/bin

WORKDIR /app/test

COPY Taskfile.yml ./

RUN go install github.com/go-task/task/v3/cmd/task@latest && \
    chmod -R 777 /home/test

ENTRYPOINT ["task", "test"]