FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/web-server

FROM scratch
WORKDIR /app
COPY --from=builder /app/web-server web-server
EXPOSE 8080
ENTRYPOINT [ "/app/web-server" ]
