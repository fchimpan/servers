FROM golang:1.22-alpine

# WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY ./app/go.mod ./app/go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
