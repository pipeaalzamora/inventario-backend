# STEP 1: Build the Go binary
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Cache deps first (better for rebuilds)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary (statically linked)
RUN go build -o sofia-backend .

# STEP 2: Create a small final image
FROM alpine:3.20

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/sofia-backend .

# Expose port
EXPOSE 3000

# Run binary
CMD ["./sofia-backend"]
