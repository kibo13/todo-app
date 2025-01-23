FROM golang:1.23.4-alpine

RUN apk add --no-cache bash

RUN go version 
ENV GOPATCH=/

COPY ./ ./

RUN chmod +x wait-for-it.sh

RUN go mod download
RUN go build -o todo-app ./cmd/app/main.go 

CMD ["./todo-app"]
