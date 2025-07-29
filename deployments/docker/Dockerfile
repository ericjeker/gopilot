# --- Build Stage ---
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY ../../go.mod ./
COPY ../../go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY ../.. .

# Build the Go application statically
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# --- Runtime Stage ---
FROM alpine:latest

WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the application runs on (if applicable)
EXPOSE 8080

# Command to run the executable
CMD ["./main"]