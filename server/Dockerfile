# Start from a Debian-based image with the latest version of Go installed
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files from the server directory
COPY go.mod ./

# Download Go modules
RUN go mod download

# Copy the server's source code into the container
COPY . .

# Build the server application
RUN go build -o server

# Expose the port the server listens on
EXPOSE 8080

# Run the server executable
CMD ["./server"]
