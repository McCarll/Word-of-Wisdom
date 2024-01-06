# Start from a Debian-based image with the latest version of Go installed
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the client's source code into the container
COPY . .

# Download Go modules
RUN go mod download

# Build the client application
RUN go build -o client

# Set the client executable as the entrypoint
ENTRYPOINT ["./client"]
