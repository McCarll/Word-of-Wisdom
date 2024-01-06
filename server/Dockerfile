# Start from a Debian-based image with the latest version of Go installed
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the server's source code into the container
COPY . .

# Download Go modules
RUN go mod download

# Build the server application
RUN go build -o server

# Expose the port the server listens on
EXPOSE 8080

# Run the server executable
CMD ["./server"]
