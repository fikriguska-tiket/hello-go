FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o hello

CMD ["./hello"]