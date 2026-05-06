# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server ./cmd/server

# Runtime stage — alpine for SQLite file persistence support
FROM alpine:3.19

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app
COPY --from=builder /app/server .

# SQLite database will be stored in /app/data — mount a volume here in production
VOLUME /app/data
ENV DB_PATH=/app/data/portfolio.db

EXPOSE 8080
ENTRYPOINT ["./server"]
