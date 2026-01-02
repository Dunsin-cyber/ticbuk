FROM golang:1.25.3

WORKDIR /src/app

RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod tidy
RUN go mod vendor
