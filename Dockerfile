# Start with a base Go image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main ./cmd/server/main.go

# Expose the port the application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
