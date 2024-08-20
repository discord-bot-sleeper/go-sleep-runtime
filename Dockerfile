FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o sleeper .

FROM alpine:latest

ENV GIN_MODE=release

COPY --from=builder /app/sleeper /app/sleeper

EXPOSE 8080

CMD ["/app/sleeper"]
