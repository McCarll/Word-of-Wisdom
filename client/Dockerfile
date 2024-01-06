FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o client

ENTRYPOINT ["./client"]
