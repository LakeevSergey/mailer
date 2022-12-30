FROM golang:1.19

WORKDIR /usr/src/app

COPY . .

RUN go build -v -o /usr/bin/ ./cmd/server
