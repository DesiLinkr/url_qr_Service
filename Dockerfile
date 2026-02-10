# =========================
# 1. Builder stage
# =========================
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download
# Copy all project files

COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .



# =========================
# 2. Runtime stage
# =========================
FROM alpine AS runner

WORKDIR /app


COPY --from=builder /app/server .

# Expose your service port
EXPOSE 9000

# Run binary
CMD ["./server"]