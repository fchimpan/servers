FROM golang:1.21.3 as build

WORKDIR /workdir

COPY go.mod go.sum /workdir/
RUN go mod download
COPY . /workdir
RUN go build -ldflags='-s -w' -o main ./main.go

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
    ca-certificates \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
COPY --from=build /workdir/main /main

CMD ["/main"]
