# Start from a Debian-based image with the latest version of Go installed
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files from the client directory
COPY go.mod ./

# Download Go modules
RUN go mod download

# Copy the client's source code into the container
COPY . .

# Build the client application
RUN go build -o client

# Run the client executable
CMD ["./client"]
