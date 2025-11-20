FROM golang:1.25 AS builder

WORKDIR /app

COPY . .

RUN mv migrations cmd/

RUN go mod tidy

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/config/config.yaml ./config/config.yaml
COPY --from=builder /app/.env ./.env

RUN chmod +x ./server

EXPOSE 8081

ENTRYPOINT ["./server"]
