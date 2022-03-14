FROM golang:1.17-alpine

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

COPY go.mod src/AirportApi/go.mod
COPY go.sum src/AirportApi/go.sum
COPY airport-service src/AirportApi/airport-service

WORKDIR /go/src/AirportApi/airport-service

RUN CGO_ENABLED=0 GOOS=linux GARCH=amd64 go build -installsuffix cgo -o build/airport-service .

CMD ["./build/airport-service"]
EXPOSE 9000
EXPOSE 80
EXPOSE 443