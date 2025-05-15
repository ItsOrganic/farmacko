# Use the official Golang image
FROM golang:1.21

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of the app code
COPY . .

# Copy config folder
COPY config ./config

# Build the Go app binary
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Start the app with config path
CMD ["./main", "-config", "config/config.yaml"]
