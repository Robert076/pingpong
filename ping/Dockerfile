FROM golang:1.24.2-bullseye

WORKDIR /src

COPY ping.go .

ENTRYPOINT ["go", "run", "ping.go"]