FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod init github.com/piotrkira/microservices-calc/addsub
RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/ads-service

FROM scratch
WORKDIR /app
COPY --from=builder /app/ads-service ads-service
ENTRYPOINT [ "/app/ads-service" ]
