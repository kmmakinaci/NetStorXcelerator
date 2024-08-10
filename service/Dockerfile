# Official Go image
FROM golang:1.20 AS builder
# Set the current working directory inside the container
WORKDIR /NetStorXcelerator

# Copy the Go modules manifests
COPY go.mod go.sum ./
# Download dependencies
RUN go mod download
# Copy source code to container
COPY . .
# Build the Go system services
RUN go build -o service

# Target container
FROM rockylinux:8
WORKDIR /NetStorXcelerator
COPY --from=builder /NetStorXcelerator/service .

# Expose port to outside world
EXPOSE 8080
# Run the executable
CMD ["./service"]


