# Build stage
FROM golang:1.24.2-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod ./

# Download dependencies
RUN go mod tidy

# Copy source code
COPY *.go ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rpg .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/rpg .

# Copy static web files
COPY index.html .
COPY style.css .
COPY script.js .

# Expose port 8081
EXPOSE 8081

# Command to run the application in server mode
CMD ["./rpg", "server"]