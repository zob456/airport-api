FROM golang:1.17-alpine

ENV GO111MODULE=on

COPY go.mod src/AirportApi/go.mod
COPY go.sum src/AirportApi/go.sum
COPY rest-api src/AirportApi/rest-api
COPY airport-service/proto src/AirportApi/airport-service/proto

WORKDIR /go/src/AirportApi/rest-api

RUN CGO_ENABLED=0 GOOS=linux GARCH=amd64 go build -installsuffix cgo -o build/rest-api .

CMD ["./build/rest-api"]
EXPOSE 8000
EXPOSE 80
EXPOSE 443