# Stage 1: Build the Go application
FROM --platform=$BUILDPLATFORM golang:1.23.2-alpine AS builder

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Set the working directory
WORKDIR /app

# Copy only the dependency files first
COPY go.mod go.sum ./

# Download dependencies (will be cached if no changes)
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application with platform-specific optimizations
ARG TARGETOS TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} CGO_ENABLED=0 go build \
    -ldflags="-w -s" \
    -o vlink_backend ./cmd/server/main.go

# Stage 2: Create a minimal runtime image
FROM --platform=$TARGETPLATFORM alpine:3.19

# Install CA certificates for HTTPS
RUN apk add --no-cache ca-certificates tzdata

# Set the working directory
WORKDIR /app

# Copy only the compiled binary from builder
COPY --from=builder /app/vlink_backend .

# Expose the application port
EXPOSE 5000

# Set environment variables
ENV ENV=development

# Run the application
CMD ["./vlink_backend"]
