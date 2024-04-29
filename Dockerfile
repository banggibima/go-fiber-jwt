FROM golang:1.22.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /out/app ./cmd/app

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /out/app .
COPY .env .

EXPOSE 8080

CMD ["./app"]