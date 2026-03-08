# Multi-stage build for optimal image size
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Install git, download dependencies, then remove git — всё в одном слое
RUN apk add --no-cache git && \
    go mod download && \
    apk del git

# Copy source code
COPY . .

# Build the application — ldflags убирают debug info и таблицу символов (меньше размер бинарника)
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s" \
    -o main ./cmd/main.go

# Final stage - minimal runtime image
FROM alpine:3.21

LABEL description="Todo Service Go API"

# Install ca-certificates for HTTPS requests
# wget уже встроен в BusyBox (Alpine) — отдельно устанавливать не нужно
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

WORKDIR /home/appuser

# Copy binary from builder
COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations

# Change ownership to non-root user
RUN chown -R appuser:appuser /home/appuser

# Switch to non-root user
USER appuser

# Healthcheck
HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
    CMD wget -qO- http://localhost:8181/health || exit 1

# Expose port
EXPOSE 8181

# Run the application
CMD ["./main"]
