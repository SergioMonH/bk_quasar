FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . .
COPY cmd/main/satelites.json .
RUN go get -d -v ./cmd/main

RUN go build -o app -v ./cmd/main

ENTRYPOINT [ "./app" ]




#CMD ["./main"]
# docker build -t myapp . 
# dsudo s