FROM golang:1.23

WORKDIR /app

COPY . .

WORKDIR /app/producer

RUN go build -o app .

CMD ["./app"]
