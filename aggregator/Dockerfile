FROM golang:1.23

WORKDIR /app

COPY . .

WORKDIR /app/aggregator

RUN go build -o app .

CMD ["./app"]
