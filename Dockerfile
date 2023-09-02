FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o satelite-service ./cmd/main

FROM scratch

EXPOSE 8080

ENTRYPOINT ["/satelite-service"]