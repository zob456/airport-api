FROM golang:1.17-alpine

ENV GO_PATH=~/go
ENV PATH=$PATH:/$GO_PATH/bin

RUN apk --update add --no-cache
RUN apk add --no-cache make protobuf-dev=3.18.1-r1

RUN go get google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go

COPY go.mod src/AirportApi/go.mod
COPY go.sum src/AirportApi/go.sum
COPY airport-service src/AirportApi/airport-service
COPY grpcTypes src/AirportApi/airport-service

WORKDIR /go/src/AirportApi/airport-service

RUN protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        air.proto

RUN CGO_ENABLED=0 GOOS=linux GARCH=amd64 go build -installsuffix cgo -o build/airport-service .

CMD ["./build/airport-service"]
EXPOSE 8000
EXPOSE 80
EXPOSE 443