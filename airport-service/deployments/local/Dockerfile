FROM golang:1.17-alpine

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

WORKDIR /go/src/AirportApi/airport-service

COPY go.mod .
COPY go.sum .
COPY airport-service src/AirportApi/airport-service

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GARCH=amd64 go build -installsuffix cgo -o build/airport-service .

CMD ["./build/airport-service"]
EXPOSE 8000
EXPOSE 80
EXPOSE 443