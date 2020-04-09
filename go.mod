module github.com/piotrkira/microserices-calc

go 1.14

replace github.com/piotrkira/microservices-calc => /home/piotr/dev/mc

require (
	github.com/golang/protobuf v1.3.5
	github.com/gorilla/mux v1.7.4
	github.com/piotrkira/microservices-calc v0.0.0-00010101000000-000000000000
	github.com/piotrkira/wzi-project v0.0.0-20200407115037-4ce40f5fc1f0
	github.com/piotrkira/wzi-project/ads/client v0.0.0-20200407115037-4ce40f5fc1f0
	github.com/piotrkira/wzi-project/gateway v0.0.0-20200407115037-4ce40f5fc1f0
	github.com/piotrkira/wzi-project/gateway/client v0.0.0-20200407115037-4ce40f5fc1f0
	github.com/piotrkira/wzi-project/gateway/endpoints v0.0.0-20200407115037-4ce40f5fc1f0
	google.golang.org/grpc v1.28.1
)
