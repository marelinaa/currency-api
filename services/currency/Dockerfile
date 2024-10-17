FROM golang:1.22.1-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o /app/bin /app/cmd/main.go

FROM alpine AS runner

COPY --from=builder /app/bin /currency/bin
COPY --from=builder /app/migrations /currency/migrations

CMD ["/currency/bin"]