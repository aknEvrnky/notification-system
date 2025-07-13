FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o notification-system ./cmd/api

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/notification-system .
COPY .env.example .env
EXPOSE 4000
CMD ["./notification-system"]
