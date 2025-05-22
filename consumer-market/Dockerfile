FROM golang:1.23

WORKDIR /app

COPY . .

WORKDIR /app

RUN go build -o app ./cmd/

CMD ["./app"]
