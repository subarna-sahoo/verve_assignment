# Build Stage
FROM golang:1.23.6 AS builder
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o verve_assignment

# Final Image using Ubuntu (instead of Alpine)
FROM ubuntu:latest  
WORKDIR /root/

# Install required dependencies (if needed)
RUN apt update && apt install -y ca-certificates

# Copy the built binary from the builder stage
COPY --from=builder /app/verve_assignment .

# Expose the port
EXPOSE 8080

# Run the binary when the container starts
CMD ["./verve_assignment"]