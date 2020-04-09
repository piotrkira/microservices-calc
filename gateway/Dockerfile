FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

COPY netrc /root/.netrc
RUN chmod 600 /root/.netrc

RUN go mod download
RUN go mod verify

RUN go env -w GOPRIVATE=github.com/piotrkira/wzi-project

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/gateway-api

FROM scratch
WORKDIR /app
COPY --from=builder /app/gateway-api gateway-api
ENTRYPOINT [ "/app/gateway-api" ]
