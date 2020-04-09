FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod init github.com/piotrkira/microservices-calc/gateway
RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/gateway-api

FROM scratch
WORKDIR /app
COPY --from=builder /app/gateway-api gateway-api
ENTRYPOINT [ "/app/gateway-api" ]
