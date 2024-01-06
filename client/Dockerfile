FROM golang:latest

WORKDIR /app

COPY ../go.mod ./

RUN go mod download

COPY ./client .

RUN go build -o client

EXPOSE 8080

CMD ["./client"]
