FROM golang:1.17-alpine

COPY go.mod src/AirportApi/go.mod
COPY go.sum src/AirportApi/go.sum
COPY airport-service src/AirportApi/airport-service
COPY grpcTypes src/AirportApi/grpcTypes

RUN protoc --go_out=. --go_opt=paths=grpcTypes \
        --go-grpc_out=. --go-grpc_opt=paths=grpcTypes \
        airport-service/air.proto

WORKDIR /go/src/AirportApi/airport-service
RUN CGO_ENABLED=0 GOOS=linux GARCH=amd64 go build -installsuffix cgo -o build/airport-service .

CMD ["./build/airport-service"]
EXPOSE 8000
EXPOSE 80
EXPOSE 443