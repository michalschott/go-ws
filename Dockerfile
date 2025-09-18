FROM golang:1.24.2-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o server .

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
