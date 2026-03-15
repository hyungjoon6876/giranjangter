# Build stage
FROM golang:1.22-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o lincle-api ./cmd/server/

# Runtime stage
FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app
COPY --from=builder /app/lincle-api .
COPY --from=builder /app/db ./db

RUN mkdir -p /app/data /app/uploads

EXPOSE 8080

ENV ENV=production
ENV PORT=8080
ENV DATABASE_URL=file:/app/data/lincle.db?_journal_mode=WAL&_busy_timeout=5000&_foreign_keys=on

CMD ["./lincle-api"]
