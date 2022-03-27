FROM golang:alpine

RUN apk add git

WORKDIR /app

COPY . .

RUN go build -o hello

CMD ["./hello"]
