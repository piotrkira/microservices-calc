module github.com/piotrkira/microservices-calc

go 1.14

replace github.com/piotrkira/microservices-calc => /home/piotr/dev/mc

require (
	github.com/golang/protobuf v1.3.5
	github.com/gorilla/mux v1.7.4
	google.golang.org/grpc v1.28.1
)
