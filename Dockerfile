FROM golang:1.23.4-alpine

RUN go version 
ENV GOPATCH=/

COPY ./ ./

RUN go mod download
RUN go build -o todo-app ./cmd/app/main.go 

CMD ["./todo-app"]
