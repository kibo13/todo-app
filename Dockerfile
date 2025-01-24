FROM golang:1.23.4-alpine

# Install bash and dos2unix
RUN apk add --no-cache bash dos2unix  

RUN go version 
ENV GOPATCH=/

# Copy all files into the container
COPY ./ ./  

# Convert line endings to LF format
RUN dos2unix wait-for-it.sh  

# Make the wait-for-it.sh script executable
RUN chmod +x wait-for-it.sh

# Download Go modules
RUN go mod download

# Build the Go application
RUN go build -o todo-app ./cmd/app/main.go 

# Set the command to run the application
CMD ["./todo-app"]