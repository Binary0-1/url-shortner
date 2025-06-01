# Use official Go image as a builder
FROM golang:1.21-alpine AS builder

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum files first for caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app as a binary named "app"
RUN go build -o app .

# Use a minimal image for running the binary
FROM alpine:latest

# Copy the built binary from the builder
COPY --from=builder /app/app /app/app

# Expose the port your app will run on
EXPOSE 8082

# Run the app
CMD ["/app/app"]