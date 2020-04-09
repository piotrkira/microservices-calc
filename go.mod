module github.com/piotrkira/microservices-calc/gateway

go 1.14

replace github.com/piotrkira/microservices-calc => github.com/piotrkira/microservices-calc v0.0.0-20200409175932-0498954c253c

require (
	github.com/golang/protobuf v1.3.5
	github.com/gorilla/mux v1.7.4
	github.com/piotrkira/microservices-calc v0.0.0-20200409174901-68a92f7aec6a
	google.golang.org/grpc v1.28.1
)
